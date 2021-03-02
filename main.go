package main

import (
	"github.com/sirupsen/logrus"
	"messenger/db"
	"messenger/helpers"
	"messenger/server"
)

func main() {
	helpers.ConfigureLogrus()

	dbClient, err := db.ConnectToPostgres()
	if err != nil {
		logrus.Panic(err)
	}

	s := server.Init(dbClient)

	logrus.Info("Server is running on ", s.Addr)
	logrus.Panic(s.ListenAndServe())
}
