package main

import (
	"flag"
	"fmt"
	"log"
	"viper"
)

// install viper:
// go install  github.com/spf13/viper@latest

var username, password string

func userCredentials(username, password string) {
	fmt.Printf("your username is", viper.GetString("credentials.username"))
}

func main() {

	flag.StringVar(&username, "user", "", "username to use")
	flag.StringVar(&password, "password", "", "password")

	flag.Parse()

	if username == "" || password == "" {
		log.Fatalln("crednetials not provided")
	}


//    viper.setDefault()
	viper.Set("credentials.username", username)
	viper.Set("credentials.password", password)

	//read from cfg file
	viper.AddConfigPath(".")
	viper.SetConfigName("creds")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("unable toread the config file")
	}
	userCredentials(username, password)
}
