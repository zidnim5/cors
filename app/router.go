package app

import (
	"cors/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(tagsController controller.TagsController) *httprouter.Router {
	router := httprouter.New()

	// tags
	router.GET("/api/tags", tagsController.Index)

	return router

}
