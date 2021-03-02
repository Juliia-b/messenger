package server

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"messenger/db"
	"net/http"
)

func sendErr(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

// userRegistration carries out the process of user userRegistration in the system.
func (h *handler) userRegistration(w http.ResponseWriter, r *http.Request) {
	var userFirstName = r.FormValue("firstname")
	var userLastName = r.FormValue("lastname")
	var userNickname = r.FormValue("nickname")
	var password = r.FormValue("password")

	err := validatePostParameters(userFirstName, userLastName, userNickname, password)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err)
		return
	}

	user := &db.User{FirstName: userFirstName, LastName: userLastName, Nickname: userNickname, Password: hashedPassword}

	cookie, err := generateCookie(user)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err)
		return
	}

	_, err = h.DbCli.InsertUser(user)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err)
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

/* ----- helpers ----- */

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
