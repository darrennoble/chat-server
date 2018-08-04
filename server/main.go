package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
)

var configFile = flag.String("config", getHomeFile(".chat-server.json"), "config file path")
var writeConfig = flag.Bool("writeConfig", false, "write the default values for missing fields to the config file")

func main() {
	flag.Parse()

	chatSever, err := NewChatServer(*configFile)
	if err != nil {
		log.Fatalf("Error creating ChatServer: %v", err)
	}

	if *writeConfig {
		err = chatSever.WriteConfig(*configFile)
		if err != nil {
			log.Fatalf("Error writing config file: %v", err)
		}
	}

	chatSever.Start()
}

func getHomeFile(file string) string {
	u, err := user.Current()
	if err != nil {
		log.Fatalf("Error lookingup home folder: %v", err)
	}

	return fmt.Sprintf("%s%s%s", u.HomeDir, string(os.PathSeparator), file)
}
