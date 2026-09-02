package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func main() {
	err := errors.New("ERROR")
	log.Println(err, reflect.TypeOf(err))

	log.Println(" = = = = = = =")
	err = fmt.Errorf("ERROR-2 :%w", err)
	err = fmt.Errorf("ERROR-3 :%w", err)

	log.Println(err, reflect.TypeOf(err))

	log.Println(" = = = = = = =")
	for e := err; e != nil; e = errors.Unwrap(e) {
		log.Printf(" -> %v\n", e)
	}
}
