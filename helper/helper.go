package helper

import (
	domain "cors/domain/response"
	"encoding/json"
	"net/http"
)

/**
 * todo: toggle above func
 */
func WriteToResponseBody(writer http.ResponseWriter, webResponse domain.WebResponse) {
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}
