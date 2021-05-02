package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

type Config struct {
	WorkingDirectory string     `json:"WorkingDirectory"`
	Key []string      			`json:"key"`
	Offset string 				`json:"offset"`
}

type Parameter struct {
	Config string
	Port string
}


func main() {
	fmt.Println("Starting")
	para := getCommandLineArguments()

	config, err := getConfigData(para.Config)
	if err != nil {
		errors.New("Couldn't load config")
	}

	handleRequests(para, config)
}

func handleRequests(para Parameter, config Config) {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", defaultHandler)
	myRouter.HandleFunc("/health", getHealth)
	myRouter.HandleFunc("/tiles/{z}/{x}/{y}", getTile)

	log.Fatal(http.ListenAndServe(":" + para.Port , myRouter))
}

func defaultHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Listening")
}

func getHealth(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Is healthy")
}

func getTile(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", vars["z"])
}

func getCommandLineArguments() Parameter {
	args := os.Args
	var para Parameter
	para.Config = "./config/config.json"
	para.Port = "8080"

	for i := 1; i < len(args); i++ {
		switch os := args[i]; os {
		case "-config":
			i++
			para.Config = args[i]
			break

		case "-port":
			i++
			para.Port = args[i]
			break

		default:
			fmt.Printf("%s.\n", os)
		}
	}

	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXX PARAMETERS XXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("   [-config]      : " + para.Config)
	fmt.Println("   [-port]        : " + para.Port)
	fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXX PARAMETERS XXXXXXXXXXXXXXXXXXXXXXXXXXX")
	fmt.Println("")
	return para

}


func getConfigData(path string) (Config, error) {
	var result Config
	if FileExists(path) {
		jsonFile, err := os.Open(path)
		// if we os.Open returns an error then handle it
		if err != nil {
			return result, err
		}
	
		fmt.Println("Successfully Opened " + path)
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
	
		// read our opened xmlFile as a byte array.
		byteValue, _ := ioutil.ReadAll(jsonFile)
	
		// we initialize our Users array
		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal(byteValue, &result)
		return result, nil
		
	} else {
		return result, errors.New("Couldn't load config")
	}


}
