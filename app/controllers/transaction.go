package controllers

import (
	v "gem-exp/app/utils/view"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	data := v.Message(true, "Hello")
	v.RespondSuccess(w, data)
}
