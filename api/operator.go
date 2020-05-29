package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type Operator struct {
	line          string
	phoneNumber   string
	addressPickup string
	customerType  string
	partnerCode   string
	longitude     string
	latitude      string
	callState     int
	commandState  string
}

func OperatorInsertFirstCall(op []byte) {
	log.Printf(" [x]Goi cai deo gi day %s", op)

	resp, err := http.Post(apiUrl+"taxioperator/firstcall", "application/json", bytes.NewBuffer(op))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}
