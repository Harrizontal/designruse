package loader

import (
	"log"

	"rsc.io/quote"
)

func Hello() string {
	log.Println("Hello")
	return quote.Hello()
}
