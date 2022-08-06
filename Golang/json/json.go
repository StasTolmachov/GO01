package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contacts struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsAlive   bool   `json:"isAlive"`
	Age       int    `json:"age"`
	Address   struct {
		StreetAddress string `json:"streetAddress"`
		City          string `json:"city"`
		State         string `json:"state"`
		PostalCode    string `json:"postalCode"`
	} `json:"address"`
	PhoneNumbers []struct {
		Type   string `json:"type"`
		Number string `json:"number"`
	} `json:"phoneNumbers"`
	Children []string    `json:"children"`
	Spouse   interface{} `json:"spouse"`
}

func main() {

	data, err := os.Open("/Users/st/go/src/GO01/Golang/json/contact.json")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	var cntct contacts

	if err = json.NewDecoder(data).Decode(&cntct); err != nil {
		panic(err)
	}
	fmt.Println(cntct)

	cntct.FirstName = "Stasvv"

	dataNew, err := os.OpenFile("contactNew.json", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	if err = json.NewEncoder(dataNew).Encode(cntct); err != nil {
		panic(err)
	}

	fmt.Println(cntct)

}
