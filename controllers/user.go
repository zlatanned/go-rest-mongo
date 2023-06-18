package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/zlatanned/go-rest-mongo/models"
	"gopkg.in/mgo.v2"
	"net/http"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

/*
 * TODO
	func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	}


	func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	}


	func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	}
*/
