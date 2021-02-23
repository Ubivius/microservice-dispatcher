package data

import (
	"encoding/json"
	"io"
)

func FromJSON(i interface{}, reader io.Reader) error {
	return json.NewDecoder(reader).Decode(i)
}

func ToJSON(i interface{}, writer io.Writer) error {
	return json.NewEncoder(writer).Encode(i)
}
