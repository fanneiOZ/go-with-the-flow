package handler

import (
	"domain/donation"
	"domain/payment"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"net/http"
	"sharedinfra/external/omise"
	"sharedinfra/external/omise/v20190529"
	"sharedinfra/fileio"
	"sharedinfra/httpserver"
	"sync"
	"time"
)

const (
	GroupPathPahpa = "/pahpa"
)

var (
	ErrPayloadInvalidForm = httpserver.ErrorPayload{
		Code:    "INVALID_FORM",
		Message: "Multipart form named `file` missing",
	}
	ErrPayloadInvalidImport = httpserver.ErrorPayload{
		Code:    "INVALID_IMPORT",
		Message: "Import file missing or exceeded",
	}
	ErrPayloadInvalidFile = httpserver.ErrorPayload{
		Code:    "INVALID_FILE",
		Message: "Unable to read the uploaded file",
	}
	ErrPayloadInvalidInput = httpserver.ErrorPayload{
		Code:    "INVALID_INPUT",
		Message: "Input file is not readable in rot128 or csv format",
	}

	expectedHeaders = [6]string{"Name", "AmountSubunits", "CCNumber", "CVV", "ExpMonth", "ExpYear"}
)

func bulkFryPahpaHandler() gin.HandlerFunc {
	cfg := omise.ApiConfig{}

	tokenApi := v20190529.NewTokenAPI(cfg)
	chargeApi := v20190529.NewChargeAPI(cfg)

	paymentService := payment.NewOmisePaymentService(tokenApi, chargeApi)
	chargeUseCase := payment.NewChargeCreditCard(paymentService)
	_ = donation.NewFryPahPaUseCase(chargeUseCase)

	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		if file == nil {
			httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidForm)

			return
		}

		if len(c.Request.MultipartForm.File["file"]) != 1 {
			httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidImport)

			return
		}
		fileReader, err := file.Open()
		if err != nil {
			httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidFile)

			return
		}

		defer func() {
			if err := fileReader.Close(); err != nil {
				log.Printf("Error closing file: %v", err.Error())
			}
		}()

		rot128Reader, _ := fileio.DecodeRot128(fileReader)
		csvReader := csv.NewReader(rot128Reader)

		headers, _ := csvReader.Read()
		if !validateCsvHeader(headers) {
			httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidInput)

			return
		}
		wg := sync.WaitGroup{}
		tasks := 0
		for {
			data, err := csvReader.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidInput)

				return
			}
			tasks++

			go execute(&wg, data)
		}

		wg.Wait()

		httpserver.Success(c, gin.H{"tasks": tasks})
	}
}

//func bulkFryPahpaHandler(c *gin.Context) {
//	file, _ := c.FormFile("file")
//	if file == nil {
//		httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidForm)
//
//		return
//	}
//
//	if len(c.Request.MultipartForm.File["file"]) != 1 {
//		httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidImport)
//
//		return
//	}
//	fileReader, err := file.Open()
//	if err != nil {
//		httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidFile)
//
//		return
//	}
//
//	defer func() {
//		if err := fileReader.Close(); err != nil {
//			log.Printf("Error closing file: %v", err.Error())
//		}
//	}()
//
//	rot128Reader, _ := fileio.DecodeRot128(fileReader)
//	csvReader := csv.NewReader(rot128Reader)
//
//	headers, _ := csvReader.Read()
//	if !validateCsvHeader(headers) {
//		httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidInput)
//
//		return
//	}
//	wg := sync.WaitGroup{}
//	tasks := 0
//	for {
//		data, err := csvReader.Read()
//		if err == io.EOF {
//			break
//		}
//
//		if err != nil {
//			httpserver.Error(c, http.StatusBadRequest, ErrPayloadInvalidInput)
//
//			return
//		}
//		tasks++
//
//		go execute(&wg, data)
//	}
//
//	wg.Wait()
//
//	httpserver.Success(c, gin.H{"tasks": tasks})
//}

func execute(wg *sync.WaitGroup, input []string) {
	wg.Add(1)
	defer wg.Done()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

	log.Println(input)
}

func validateCsvHeader(input []string) bool {
	for pos, header := range input {
		if header != expectedHeaders[pos] {
			return false
		}
	}

	return true
}

func PahpaRouterGroup(engine *gin.Engine) *gin.RouterGroup {
	routerGroup := engine.Group(GroupPathPahpa)
	routerGroup.POST("/bulk", bulkFryPahpaHandler())

	return routerGroup
}
