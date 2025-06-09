package cipher_test

import (
	"bytes"
	"sharedinfra/cipher"
	"testing"
)

func TestRot128Reader(t *testing.T) {
	t.Run("should read the buffer in ROT128 approach", func(t *testing.T) {
		given := bytes.NewBuffer([]byte{128, 255, 0})
		reader, err := cipher.NewRot128Reader(given)
		if err != nil {
			t.Fatal(err)
		}

		result := make([]byte, 3)
		expected := []byte{0, 127, 128}

		read, err := reader.Read(result)

		if err != nil {
			t.Fatalf("should not return error, received %v", err)
		}
		if read != 3 {
			t.Fatalf("read should be 3, received %d", read)
		}

		if !bytes.Equal(expected, result) {
			t.Fatalf("should return expected result, received %v, expected %v", result, expected)
		}
	})
}
