package cipher

import "io"

type rot128Writer struct {
	input io.Writer
}

func NewRot128Writer(w io.Writer) (io.Writer, error) {
	return &rot128Writer{input: w}, nil
}

func (w *rot128Writer) Write(p []byte) (int, error) {
	encoded := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		encoded[i] = p[i] + 128 // rot128 transformation
	}

	return w.input.Write(encoded)
}
