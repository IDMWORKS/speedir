package processor

import (
	"log"

	"github.com/idmworks/speedir/errors"
	"github.com/idmworks/speedir/models"
	"github.com/mmitton/asn1-ber"
)

func handleBindRequest(messageID uint64, request *ber.Packet) {
	version := request.Children[0].Value.(uint64)
	username := request.Children[1].Value.(string)
	auth := request.Children[2]
	password := auth.Data.String()

	log.Println("\nBindRequest:",
		"\n\tmessageID:", messageID,
		"\n\tLDAP version:", version,
		"\n\tusername:", username,
		"\n\tpassword:", "********")

	var users []models.User
	_, err := DbMap.Select(&users, "select * from users where username=$1", username)
	errors.CheckErr(err, "Select failed")

	if len(users) == 1 {
		log.Println("User found:", username)

		if users[0].ComparePassword(password) {
			log.Println("Password for user valid:", username)
		} else {
			log.Println("Password for user invalid:", username)
		}

	} else {
		log.Println("User not found:", username)
	}
}
