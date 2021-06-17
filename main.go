package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"log"
	"strings"
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

type AddSource struct {
	Url      string `json:"url"`
}

var config Config
var para Parameter

func main() {
	fmt.Println("Starting")
	para = getCommandLineArguments()

	configRaw, err := getConfigData(para.Config)
	if err != nil {
		errors.New("Couldn't load config")
	} else {
		config = configRaw
	}

	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", defaultHandler)
	myRouter.HandleFunc("/health", getHealth)
	myRouter.HandleFunc("/tiles/{source}/{z}/{x}/{y}", getTile)
	myRouter.HandleFunc("/tiles/{source}", addSource).Methods("POST")

	log.Fatal(http.ListenAndServe(":" + para.Port , myRouter))
}

func defaultHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Listening")
}

func getHealth(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Is healthy")
}

func getTile(w http.ResponseWriter, r *http.Request){
	// curl localhost:8080/tiles/base/180/9/137.hgt?key=23
	key := r.FormValue("key")
	if CheckKey(config.Key, key) {
		vars := mux.Vars(r)
		y := strings.Split(vars["y"], ".")
		url := GetTileUrl(config, vars["source"], vars["z"], vars["x"], y[0], y[1])
		w.Header().Set("Content-Type", GetMimeTypeFromFileType(y[1]))
		http.ServeFile(w, r, url)

	} else {

	}
}

func addSource(w http.ResponseWriter, r *http.Request){
	// curl localhost:8080/tiles/base -X POST -H "key: asdf" -H "Content-Type: application/json" --data "{\"url\": \"https://onedrive.live.com/download?cid=1D70A897B7EE577D&resid=1D70A897B7EE577D%2141908&authkey=AAq3K6TPlnu68KQ\",\"dest\":\"latest/base\",\"name\":\"wt_top\"}"
	key := r.Header.Get("key")
	if CheckKey(config.Key, key) {
		vars := mux.Vars(r)
		reqBody, _ := ioutil.ReadAll(r.Body)
		var source AddSource 
		json.Unmarshal(reqBody, &source)
		// update our global Articles array to include
		// our new Article
	
		dest, err := DownloadTileSource(config, vars["source"], source.Url)
		if err != nil {
			fmt.Fprintf(w, "Couldn't download file")
		}
		target := [2]string{config.Offset, vars["source"]}
		err = UnzipFile(dest, strings.Join((target)[:], "/"))
		if err != nil {
			fmt.Fprintf(w, "Couldn't unzip file")
		}
		fmt.Fprintf(w, "Success")

	} else {
		fmt.Fprintf(w, "Key check failed")

	}

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
