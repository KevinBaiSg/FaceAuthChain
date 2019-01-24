package face

import (
	"bytes"
	"encoding/gob"
	"github.com/Kagami/go-face"
)

func Descriptor2Bytes(descriptor face.Descriptor) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(descriptor)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Bytes2Descriptor(data []byte) (*face.Descriptor, error) {
	var buf bytes.Buffer
	var descriptor face.Descriptor

	buf.Write(data)
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(&descriptor)
	if err != nil {
		return nil, err
	}
	return &descriptor, nil
}