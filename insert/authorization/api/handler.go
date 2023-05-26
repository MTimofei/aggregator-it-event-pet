package api

import (
	"net/http"
	"new/insert/authorization/auth/myauth"
	"new/insert/authorization/token"
	"new/insert/authorization/token/mytoken"
	"time"
)

func handlerReg(w http.ResponseWriter, r *http.Request) {
	var a = myauth.New(r.FormValue("user"), r.FormValue("password"))
	err := a.Reg()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handlerAuth(w http.ResponseWriter, r *http.Request) {
	var a = myauth.New(r.FormValue("user"), r.FormValue("password"))
	user, err := a.Auth()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	k, err := mytoken.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := k.Create(token.NewUser(user.Id, user.Login, user.Role))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var cookie = http.Cookie{
		Name:     "AccJWT",
		Value:    token,
		Path:     "/",
		Domain:   "localhost:8080",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
}
