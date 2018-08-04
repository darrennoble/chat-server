package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Server is the config data for a chat server
type Server struct {
	Port       int    `json:"port"`
	IP         string `json:"ip"`
	LogFile    string `json:"log_file,omitempty"`
	Historylen int    `json:"history_length"`
}

var defaults = Server{
	Port:       8765,
	IP:         "",
	LogFile:    "~/.chat-server.log",
	Historylen: 10,
}

//Load reads a json config file and returns a Server struct or error
func Load(file string) (*Server, error) {
	data, err := ioutil.ReadFile(file)
	exists := !os.IsNotExist(err)
	if err != nil && exists {
		return nil, fmt.Errorf("Error loading config file (%v): %v", file, err)
	}

	s := &Server{}
	if exists {
		err = json.Unmarshal(data, &s)
		if err != nil {
			return nil, fmt.Errorf("Error unmarshalling config file (%v): %v", file, err)
		}
	}

	err = copyNonzero(defaults, s)
	if err != nil {
		return nil, fmt.Errorf("Error checking for defaults: %v", err)
	}

	return s, nil
}

//Save writes a json config file to disk or returns an error
func (s Server) Save(file string) error {
	data, err := json.MarshalIndent(&s, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling save data: %v", err)
	}

	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing save file: %v", err)
	}
	return nil
}
