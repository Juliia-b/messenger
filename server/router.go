package server

import (
	"github.com/gorilla/mux"
	"messenger/db"
	"net/http"
	"time"
)

type Server struct {
	router *mux.Router
	Server *http.Server
}

func Init(dbCli db.DB) *http.Server {
	var r = mux.NewRouter()
	var h = handlerInit(dbCli)

	r.Use(h.headersMiddleware)

	r.HandleFunc("/user/registration", h.userRegistration).Methods("POST")

	return &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

// сервер посылает рандомную строку клиенту
// клиент хуширует ? пароль этой строкой
// посылает пароль серверу
// сервер выполняет нужное

// идентификация - проверка по никнейму (существует ли такой в базе)
// аутентификации - пользователь доказывает,что он является человеком, который регистрировался под никнеймом. (пароль, отпечаток пальца)
// Если все верно, и пара логин-пароль верны, то система предоставит пользователю доступ к его ресурсам и совершение банковских операций, то есть произойдет авторизация.

//// run the update handler in database "top_3_hour"
//trackTopMsgsIn3Hours(h)
//
//r.HandleFunc("/ws", h.UpgradeToWs).Methods("GET")
//r.HandleFunc("/best", h.getBestInPeriod).Methods("GET").Queries("period", "{period}")
//r.HandleFunc("/best/3hour", h.getTopMsgsIn3Hours).Methods("GET") // DEPRECATE
//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui/")))
//r.Use(h.sessionMiddleware)
