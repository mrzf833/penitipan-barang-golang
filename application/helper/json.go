package helper

import (
	"encoding/json"
	"net/http"
	"penitipan-barang/application/model/web"

	"github.com/mitchellh/mapstructure"
)

func ReadFromRequestBody(r *http.Request, result any) {
	r.ParseForm()
	r.ParseMultipartForm(0)
	if len(r.Form) > 0 {
		dataBody := make(map[string]any)
		for key, value := range r.Form {
			if len(value) == 1 {
				dataBody[key] = value[0]
			} else {
				dataBody[key] = value
			}
		}
		config := &mapstructure.DecoderConfig{TagName: "json", Result: result, WeaklyTypedInput: true}
		decoder, err := mapstructure.NewDecoder(config)
		PanicIfError(err)
		err = decoder.Decode(dataBody)
		PanicIfError(err)
	} else {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(result)
		PanicIfError(err)
	}
}

func WriteToResponseBody(w http.ResponseWriter, response any) {
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}

func WriteToResponseBodyAuthorizationError(w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	webResponse := web.WebResponse{
		Message: "UNAUTHORIZED",
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}

func WriteToResponseBodyForbiddenError(w http.ResponseWriter) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	webResponse := web.WebResponse{
		Message: "FORBIDDEN",
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}
