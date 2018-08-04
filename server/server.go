package main

import (
	"github.com/darrennoble/chat-server/server/config"
	"github.com/darrennoble/chat-server/server/msg"
)

type ChatServer struct {
	conf *config.Server
	mgr  *msg.Manager
}

func NewChatServer(confFile string) (*ChatServer, error) {
	cs := &ChatServer{}
	conf, err := config.Load(confFile)
	cs.conf = conf
	return cs, err
}

func (cs *ChatServer) Start() {
	return
}

func (cs ChatServer) WriteConfig(confFile string) error {
	return cs.conf.Save(confFile)
}
