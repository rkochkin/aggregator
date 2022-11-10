package main

import (
	"aggregator/agm_telegram"
	"aggregator/agm_telegram/app"
	"aggregator/agm_twitter"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type MainConfig struct {
	Version        string                      `json:"version"`
	TelegramConfig agm_telegram.TelegramConfig `json:"telegram"`
	TwitterConfig  agm_twitter.TwitterConfig   `json:"twitter"`
}

func GenDefaultConfig(path string) {
	arr, _ := json.MarshalIndent(MainConfig{}, "", "  ")

	// Creating an empty file
	// Using Create() function
	simpleConfigFile, e := os.Create(path)
	if e != nil {
		log.Fatal(e)
	}
	log.Println(simpleConfigFile)
	defer simpleConfigFile.Close()
	_, err := simpleConfigFile.Write(arr)
	if err != nil {
		panic(err)
	}
	simpleConfigFile.Sync()
}

func main() {
	resetConfigs := flag.Bool("reset-configs", false, "Restore all default configs")
	flag.Parse()

	if *resetConfigs {
		GenDefaultConfig("config/config.json")
	}

	// Open our jsonFile
	jsonFile, err := os.Open("config/config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Successfully Opened config/config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our Users array
	var config MainConfig
	//c := make(map[string]interface{})

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	if err := json.Unmarshal(byteValue, &config); err != nil {
		panic(err)
	}

	t := agm_telegram.NewTelegram(config.TelegramConfig, app.RpnCalcApp)
	t.Run()
}
