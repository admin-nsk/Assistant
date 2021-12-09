package app_context

import (
	"note_service/app/internal/config"
	"note_service/app/pkg/logging"
	"sync"
)

type AppContext struct {
	Config *config.Config
}

var instance *AppContext
var once sync.Once

func GetInstance() *AppContext {
	logging.Getlogger().Println("initialization application context")
	once.Do(func() {
		instance = &AppContext{
			Config: config.GetConfig(),
		}
	})
	return instance
}
