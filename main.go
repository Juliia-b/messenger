package main

import (
	"github.com/sirupsen/logrus"
	"messenger/helpers"
	"messenger/server"
)

func main() {
	helpers.ConfigureLogrus()

	s := server.Init()

	logrus.Info("Server is running on ", s.Addr)
	logrus.Panic(s.ListenAndServe())
}
