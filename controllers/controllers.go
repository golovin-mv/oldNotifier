package controllers

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type ControllerManager struct {
	gin *gin.Engine
}

func (m *ControllerManager) register(f func(r *gin.Engine)) {
	f(m.gin)
}

func (m *ControllerManager) run() {
	if err := m.gin.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}

func NewControllerManager() *ControllerManager {
	g := gin.Default()
	m := ControllerManager{g}
	m.register(MessageController)

	return &m
}
