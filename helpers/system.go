package helpers

import (
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
	"os"
)

// ConfigureLogrus minimally configures logrus.
func ConfigureLogrus() {
	formatter := runtime.Formatter{
		ChildFormatter: &logrus.TextFormatter{
			TimestampFormat: "02-01-2006 15:04:05", // "Mon Jan 2 15:04:05 MST 2006"
			FullTimestamp:   true,
		},
		File:         true,
		BaseNameOnly: true}

	formatter.Line = true

	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

type env struct {
	key     string
	comment string
}

var envs = []env{
	{"PGHOST", "Host on which postgreSQL is listened."},
	{"PGPORT", "Port on which postgreSQL is listened."},
	{"PGUSER", "PostgreSQL user name."},
	{"PGPASSWORD", "Password to access postgreSQL."},
}

// CheckEnv checks for all required global variables.
func CheckEnv() {
	for _, env := range envs {
		if os.Getenv(env.key) == "" {
			logrus.Fatalf("Missing global variable %v. Usage : %v\n ", env.key, env.comment)
		}
	}
}
