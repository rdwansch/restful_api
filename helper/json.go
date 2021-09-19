package helper

import (
	"encoding/json"
	"net/http"
)

func ReadJSONFromRequest(request *http.Request, noteRequest interface{})  {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(noteRequest)
	PanicIfError(err)
}

func WriteJSONToBody(writer http.ResponseWriter, noteResponse interface{})  {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(noteResponse)
	PanicIfError(err)
}

