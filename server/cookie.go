package server

import (
	"encoding/base64"
	"encoding/json"
	"messenger/db"
	"net/http"
)

var cookieName = "uu"

type cookiePayload struct {
	id        int64  `json="user_id"`
	firstName string `json="first_name"`
	lastName  string `json="last_name"`
	nickname  string `json="nickname"`
}

func generateCookie(user *db.User) (*http.Cookie, error) {
	cookiePayload := getCookiePayload(user)

	cookiePayloadAsBytes, err := json.Marshal(cookiePayload)
	if err != nil {
		return nil, err
	}

	payload := encodeCookiePayload(cookiePayloadAsBytes)

	var ONE_DAY_IN_SECONDS = 24 * 60 * 60
	var DAYS_COUNT = 7

	return &http.Cookie{
		Name:     cookieName,
		Value:    payload,
		Path:     "/",
		MaxAge:   ONE_DAY_IN_SECONDS * DAYS_COUNT,
		Secure:   true,
		HttpOnly: false,
	}, nil
}

func getCookiePayload(user *db.User) *cookiePayload {
	return &cookiePayload{
		id:        user.ID,
		firstName: user.FirstName,
		lastName:  user.LastName,
		nickname:  user.Nickname,
	}
}

func encodeCookiePayload(payload []byte) string {
	return base64.StdEncoding.EncodeToString(payload)
}

func decodeCookiePayload(encodedCookie string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encodedCookie)
}
