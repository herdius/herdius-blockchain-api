package handler

import (
	"encoding/json"
	"io"
)

func newJSONEncoder(w io.Writer) *json.Encoder {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder
}
