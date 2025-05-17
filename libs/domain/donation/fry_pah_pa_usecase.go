package donation

import (
	"domain/payment"
	"encoding/csv"
	"errors"
	"fmt"
	"shareddomain/money"
	"sync"
	"io"
	"strconv"
	"strings"
)

var (
	ErrInvalidHeader = errors.New("input file has the invalid headers")

	expectedHeaders = [6]string{
		"Name",
		"AmountSubunits",
		"CCNumber",
		"CVV",
		"ExpMonth",
		"ExpYear",
	}
)

type FryPahPaUseCase struct {
	tonPahPa      *TonPahPa
	chargeUseCase *payment.ChargeCreditCard
}

func NewFryPahPaUseCase(chargeCreditCardUseCase *payment.ChargeCreditCard) *FryPahPaUseCase {
	return &FryPahPaUseCase{
		tonPahPa:      CreateTonPahPa(),
		chargeUseCase: chargeCreditCardUseCase,
	}
}

func (uc *FryPahPaUseCase) Execute(inputReader io.Reader) (TonPahPaSummary, error) {
	csvReader := csv.NewReader(inputReader)
	headerRow, err := csvReader.Read()
	if err != nil {
		return TonPahPaSummary{}, err
	}

	if !validateHeaderRow(headerRow) {
		return TonPahPaSummary{}, fmt.Errorf("unexpected csv headers: %w", err)
	}

	rowNumber := 0
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		rowNumber++

		if err != nil {
			// log error and skip
			continue
		}

		pahPaDto, err := parseRecord(record)
		if err != nil {
			// parse data error: log and skip
			continue
		}

		if err := uc.donate(pahPaDto, rowNumber); err != nil {
			// donation error: log and skip
			continue
		}
	}

	return uc.tonPahPa.Summary(), nil
}

func (uc *FryPahPaUseCase) ExecuteBulk(reader *csv.Reader) error {
	if !validateCsvHeader(reader) {
		return ErrInvalidHeader
	}

	wg := sync.WaitGroup{}
	rowNumber := 0
	for {
		data, err := reader.Read()
		if err == io.EOF {
			break
		}
		rowNumber++

		if err != nil {
			continue
		}

		go func(wg *sync.WaitGroup, record []string) {
			wg.Add(1)
			pahPaDto, _ := parseRecord(record)
			uc.donate(pahPaDto, rowNumber)

			defer wg.Done()
		}(&wg, data)

	}

	wg.Wait()

	return nil
}

func (uc *FryPahPaUseCase) donate(input PahPaDto, rowNumber int) error {
	amount, err := strconv.ParseFloat(input.AmountSubunits, 64)
	if err != nil {
		return fmt.Errorf("parse donation amount error: %w", err)
	}

	inputExpMonth, err := strconv.ParseUint(input.ExpMonth, 10, 8)
	if err != nil {
		return fmt.Errorf("parse credit card expMonth error: %w", err)
	}

	inputExpYear, err := strconv.ParseInt(input.ExpYear, 10, 16)
	if err != nil {
		return fmt.Errorf("parse credit card expYear error: %w", err)
	}

	donateByCard, err := payment.CreateCard(
		input.CCNumber,
		input.Name,
		input.CVV,
		uint8(inputExpMonth),
		int(inputExpYear),
	)
	if err != nil {
		return fmt.Errorf("donating card is invalid: %w", err)
	}

	songPahPa, err := CreateSongPahPa(
		strconv.FormatInt(int64(rowNumber), 10),
		input.Name,
		money.New(PahPaCurrency, amount/100),
		donateByCard,
	)
	if err != nil {
		return fmt.Errorf("unable to create song-pah-pa: %w", err)
	}

	transaction := uc.chargeUseCase.Execute(
		songPahPa.DonateByCard(),
		songPahPa.DonateAmount().Amount(),
		songPahPa.DonateAmount().Currency(),
	)
	songPahPa.AttachTransaction(transaction)
	err = uc.tonPahPa.AddSongPahPa(songPahPa)
	if err != nil {
		return fmt.Errorf("unable to attach song-pah-pa: %w", err)
	}

	return nil
}

func validateHeaderRow(input []string) bool {
	if len(input) != len(expectedHeaders) {
		return false
	}

	for pos, header := range input {
		if header != expectedHeaders[pos] {
			return false
		}
	}

	return true
}

func validateCsvHeader(reader *csv.Reader) bool {
	headers, err := reader.Read()
	if err != nil {
		return false
	}

	if len(headers) != len(expectedHeaders) {
		return false
	}

	for pos, header := range headers {
		if header != expectedHeaders[pos] {
			return false
		}
	}

	return true
}

func parseRecord(record []string) (PahPaDto, error) {
	if len(record) < 6 {
		return PahPaDto{}, errors.New("record is too short")
	}

	return PahPaDto{
		Name:           strings.TrimSpace(record[0]),
		AmountSubunits: strings.TrimSpace(record[1]),
		CCNumber:       strings.TrimSpace(record[2]),
		CVV:            strings.TrimSpace(record[3]),
		ExpMonth:       strings.TrimSpace(record[4]),
		ExpYear:        strings.TrimSpace(record[5]),
	}, nil
}
