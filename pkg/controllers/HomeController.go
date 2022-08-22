package controllers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	respmsg := map[string]string{
		"message": "Hello Go Developer",
	}
	res, _ := json.Marshal(respmsg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
