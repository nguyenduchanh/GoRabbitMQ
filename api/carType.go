package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CarType struct {
	Id         int
	StaxiType  int
	Name       string
	Seat       int
	IsActive   bool
	Type       int
	OrderBy    int
	UpdateTime time.Time
	CaroType   string
}

func GetAllCarType() {
	resp, err := http.Get(apiUrl + "cartype")
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
func InsertCarType(carType *CarType) {
	requestBody, err := json.Marshal(carType)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(apiUrl+"cartype", "application/json", bytes.NewBuffer(requestBody))
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
func UpdateCaeType(id int, carType *CarType) {

	client := &http.Client{}
	requestBody, err := json.Marshal(carType)
	log.Printf(" [x] %s", requestBody)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.NewRequest(http.MethodPut, apiUrl+"cartype/"+strconv.Itoa(id), bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}
	resp.Header.Set("Content-Type", "application/json; charset=utf-8")
	req, err := client.Do(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(req.StatusCode)
}
