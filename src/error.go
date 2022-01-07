package src

import "log"

var err error

func checkErr() {
	if err != nil {
		log.Panicln(err)
	}
}
