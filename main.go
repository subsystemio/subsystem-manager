package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//HashData describes expected hash results of source and build files.
type HashData struct {
	Source string `json:"source"`
	Build  string `json:"build"`
}

//RequirementsData outlines required running subsystems.
type RequirementsData struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

//APIData describes inputs and outputs of subsystem.
type APIData struct {
	Read  []string `json:"read"`
	Write []string `json:"write"`
}

//SubSystem main structure.
type SubSystem struct {
	Name         string             `json:"name"`
	Version      string             `json:"version"`
	Hash         HashData           `json:"hash"`
	API          APIData            `json:"api"`
	Requirements []RequirementsData `json:"requirements"`
	Limit        int                `json:"limit"`
}

func main() {
	file, _ := ioutil.ReadFile("./subsystem.json")

	s := SubSystem{}
	json.Unmarshal(file, &s)
	fmt.Printf("Found: %v\n", s)
}
