package main

import (
	"github.com/marcossegovia/apiai-go"
	"fmt"
)

// Dialog Flow required parameters
var (
	Token      = "566af48d6e7a48e88ed978c895ad88f5"
	QueryLang  = "es"    //Default en
	SpeechLang = "es-ES" //Default en-US
	Sessionid  = "b2f1bfae-a5f6-45ff-8041-0c7666d0735e"
)

// Dialog Flow main connection
func initDialogFlow(){

	client, err := apiai.NewClient(
		&apiai.ClientConfig{
			Token:      Token,
			QueryLang:  QueryLang,
			SpeechLang: SpeechLang,
		},
	)
	if err != nil {
		fmt.Printf("%v", err)
	}

	qr, err := client.Query(apiai.Query{Query: []string{"Hola"}, SessionId: Sessionid})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v", qr.Result.Fulfillment.Speech)

}