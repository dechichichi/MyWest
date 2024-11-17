package service

import (
	"net/http"
	jwtMaker "todolist/jwt"
)

func HandleRequest(w http.ResponseWriter, r *http.Request, claims *jwtMaker.Payload) {

}
