package handler

import (
	"encoding/json"
	"github.com/tapojit047/learning-go/api-server/db"
	"github.com/tapojit047/learning-go/api-server/model"
	"log"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var info model.Credential

	json.NewDecoder(r.Body).Decode(&info)

	pass, ok := db.CredentialList[info.Username]
	log.Println(info.Password, info.Username, pass, ok)
	if !ok || pass != info.Password {
		json.NewEncoder(w).Encode("Invalid Authorization")
		return
	}
	expiretime := time.Now().Add(15 * time.Minute)
	_, tokenString, _ := db.TokenAuth.Encode(map[string]interface{}{"user_id": 123})

	http.SetCookie(w, &http.Cookie{
		Name:    "cookie",
		Value:   tokenString,
		Expires: expiretime,
	})
}
