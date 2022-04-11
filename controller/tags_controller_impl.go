package controller

import (
	domain "cors/domain/response"
	"cors/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TagsControllerImpl struct {
}

func NewsTagsController() TagsController {
	return &TagsControllerImpl{}
}

func (t *TagsControllerImpl) Index(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	webResponse := domain.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Ok",
		Data:    "Hello, here golang service 👋",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
