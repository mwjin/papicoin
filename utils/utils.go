package utils

import "log"

func HandleErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}
