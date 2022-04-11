package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TagsController interface {
	Index(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
