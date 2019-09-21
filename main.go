package main

import (
	"bytes"
	"encoding/json"
	"github.com/Pickausernaame/HighloadHW2/application"
	"github.com/Pickausernaame/HighloadHW2/models"
	log "gopkg.in/inconshreveable/log15.v2"
	"io/ioutil"
	"os"
)

func main() {
	// Чтение аргументов командной строки
	if len(os.Args) < 2 {
		log.Crit("you must pass the path to the configuration file as the first argument")
		os.Exit(1)
	}

	// Чтение файла из пути, который был подан в качестве аргумента
	configBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Crit("unable to open configuration file '" + os.Args[1] + "': " + err.Error())
		os.Exit(1)
	}

	// Декодирование информации из JSON файла в струкуру данных
	config := &models.Config{}
	err = json.NewDecoder(bytes.NewReader(configBytes)).Decode(&config)
	if err != nil {
		log.Crit("unable to parse configuration file ./config.json : " +
			"it should be json with all fields : " + err.Error())
		os.Exit(1)
	}

	// Создание нового приложения и его запуск
	app, err := application.New(config)
	if err != nil {
		log.Crit("We have some troubles with creating application " + err.Error())
		os.Exit(1)
	}
	err = app.Run(config.ServerPort)
	if err != nil {
		log.Crit("We have some troubles with running application " + err.Error())
		os.Exit(1)
	}

}
