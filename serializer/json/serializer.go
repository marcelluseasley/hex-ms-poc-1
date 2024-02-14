package json 

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/marcelluseasley/hex-ms-poc-1/shortener"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return bytes, nil
}