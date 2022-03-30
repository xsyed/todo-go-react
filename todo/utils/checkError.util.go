package utils

import (
	"log"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}
}
