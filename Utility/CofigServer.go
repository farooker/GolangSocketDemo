package utility

import (
	"encoding/json"
	"fmt"
	"os"
)

var ConfigServer EntConfiguration

type EntConfiguration struct {

	Port   string `json:"port"`
	Ip   string `json:"ip"`
}
func InitializeConfigServer() {
	defer panicHandler();
	file, err := os.Open("./Config/ConfigServer.json")
	if err != nil {
		panic("nofileException")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ConfigServer)
	if err != nil {
		panic(err)
	}
}
func panicHandler(){
	x := recover()
	if x == "nofileException"{
		fmt.Printf("Cannot read file,program . \n")
	}
}
