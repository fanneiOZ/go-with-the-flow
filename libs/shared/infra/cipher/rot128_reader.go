package cipher

import "io"

type rot128Reader struct {
	input io.Reader
}

func NewRot128Reader(r io.Reader) (io.Reader, error) {
	return &rot128Reader{input: r}, nil
}

func (r *rot128Reader) Read(p []byte) (int, error) {
	byteRead, err := r.input.Read(p)
	for i := 0; i < byteRead; i++ {
		p[i] = p[i] + 128
	}

	return byteRead, err
}
