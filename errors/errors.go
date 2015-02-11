package errors

import "log"

// CheckErr calls log.Fatalln if there was an error
func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
