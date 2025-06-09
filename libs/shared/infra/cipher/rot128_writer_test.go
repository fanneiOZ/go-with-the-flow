package cipher_test

import (
	"bytes"
	"sharedinfra/cipher"
	"testing"
)

func TestRot128Writer(t *testing.T) {
	t.Run("should write the buffer in ROT128 approach", func(t *testing.T) {
		result := bytes.NewBuffer([]byte{})
		writer, err := cipher.NewRot128Writer(result)
		if err != nil {
			t.Errorf("should not have an error")
		}

		expected := []byte{128, 255, 0}

		byteWritten, _ := writer.Write([]byte{0, 127, 128})

		if byteWritten != 3 {
			t.Errorf("should have written 3 bytes")
		}

		if !bytes.Equal(expected, result.Bytes()) {
			t.Errorf("should have written bytes")
		}
	})
}
