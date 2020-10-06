package main

import (
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Persons struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	Prename string `json:"prename"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func main() {

	//Open Example JSON File
	f, err := os.Open("example.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Read Data and Parse to Structs
	bytes, _ := ioutil.ReadAll(f)
	var persons Persons
	json.Unmarshal(bytes, &persons)

	//Hashing Names
	h := sha256.New()
	for i, p := range persons.Persons {
		h.Write([]byte(p.Prename))
		p.Prename = string(h.Sum(nil))

		h.Reset()

		h.Write([]byte(p.Surname))
		p.Surname = string(h.Sum(nil))

		persons.Persons[i] = p
	}

	//Write Results to File
	result, _ := json.Marshal(persons)
	ioutil.WriteFile("result.json", result, 0644)
}
