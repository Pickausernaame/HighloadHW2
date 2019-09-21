package main

import (
	"github.com/Pickausernaame/HighloadHW2/application"
	log "gopkg.in/inconshreveable/log15.v2"
	"os"
)

func main() {

	// Создание нового приложения и его запуск
	app, err := application.New()
	if err != nil {
		log.Crit("We have some troubles with creating application " + err.Error())
		os.Exit(1)
	}
	err = app.Run(9090)
	if err != nil {
		log.Crit("We have some troubles with running application " + err.Error())
		os.Exit(1)
	}

}
