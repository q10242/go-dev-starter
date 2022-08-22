package controllers

import (
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	respmsg := map[string]string{
		"message": "Hello Go Developer",
	}
	res, _ := json.Marshal(respmsg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Login() {

}
func Logout() {

}
func ListAllUsers() {

}
func IssueToken() {

}
func RenewToken() {

}
func EditUser() {

}

func GetUser() {

}
