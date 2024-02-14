package json 

import (
	"github.com/vmihailenco/msgpack/v4"
	"github.com/pkg/errors"
	"github.com/marcelluseasley/hex-ms-poc-1/shortener"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	bytes, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return bytes, nil
}