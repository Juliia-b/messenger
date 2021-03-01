package server

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func (h *handler) headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "Origin")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// registration carries out the process of user registration in the system.
func (h *handler) registration(w http.ResponseWriter, r *http.Request) {
	var userFirstName = r.FormValue("firstname")
	var userSecondName = r.FormValue("lastname")
	var userNickname = r.FormValue("nickname")
	var password = r.FormValue("pass")

	err := validatePostParameters(userFirstName, userSecondName, userNickname, password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(hashedPassword)
	//	push to db
}

// validatePostParameters validates post request parameters. All post parameters must be not empty.
func validatePostParameters(userFirstName, userSecondName, userNickname, password string) error {
	var stringErr string

	if isValueEmpty(userFirstName) {
		stringErr += "The 'userFirstName' field must be not empty; "
	}

	if isValueEmpty(userSecondName) {
		stringErr += "The 'userSecondName' field must be not empty; "
	}

	if isValueEmpty(userNickname) {
		stringErr += "The 'userNickname' field must be not empty; "
	}

	if isValueEmpty(password) {
		stringErr += "The 'password' field must be not empty; "
	}

	if isValueEmpty(stringErr) {
		return nil
	}

	return errors.New(stringErr)
}

func isValueEmpty(value string) (isEmpty bool) {
	if value == "" {
		return true
	}

	return false
}
