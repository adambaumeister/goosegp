package errors

import "log"

/*
Implements error handling for the BGP connector
*/

func RaiseError(e interface{}) {
	//log.Print(e)
	// For now, this kills the program but in the future this should signal the broker
	// that this particular goroutine is screwed
	log.Fatal(e)
}
