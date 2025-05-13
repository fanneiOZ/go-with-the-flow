package fileio

import (
	"github.com/opn-ooo/challenges/challenge-go/cipher"
	"io"
)

type EncodedWriter struct {
	io.Writer
	//file *os.File
}

func EncodeRot128(input io.Writer) (*EncodedWriter, error) {
	writer, err := cipher.NewRot128Writer(input)
	if err != nil {
		return nil, err
	}

	return &EncodedWriter{Writer: writer}, nil
}
