package main

import (
	"github.com/marcossegovia/apiai-go"
)

// Dialog Flow required parameters
var (
	Token      = "3c25ae9299644ea596f8e89104d5d58e"
	QueryLang  = "es"    //Default en
	SpeechLang = "es-ES" //Default en-US
	Sessionid  = "84641b96-775b-452c-b87e-fd6d8a588385"
)

// Dialog Flow main connection
func initDialogFlow() (*apiai.ApiClient, error) {

	client, err := apiai.NewClient(
		&apiai.ClientConfig{
			Token:      Token,
			QueryLang:  QueryLang,
			SpeechLang: SpeechLang,
		},
	)

	if err != nil {
		return client, err
	}

	return client, nil
}

func dialogFlowQuery(cl *apiai.ApiClient, qr string) (string, error){

	resp, err := cl.Query(apiai.Query{Query: []string{"Enviar 1 euro a Adrià Ros con concepto Cena"}, SessionId: Sessionid})

	if err != nil {
		return "Something is wrong. Please check the required dialog flow parameters.", err
	}

	r := resp.Result.Fulfillment.Speech

	return r, nil

}