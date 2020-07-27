package app

import (
	"gitlab.com/quangdangfit/gocommon/utils/logger"
	"go.uber.org/dig"

	"goshop/app/api"
	"goshop/app/repositories"
	"goshop/app/services"
)

func Init() {
	BuildContainer()
}

func BuildContainer() *dig.Container {
	container := dig.New()

	// Inject repositories
	err := repositories.Inject(container)
	if err != nil {
		logger.Error("Failed to inject repositories", err)
	}

	// Inject services
	err = services.Inject(container)
	if err != nil {
		logger.Error("Failed to inject services", err)
	}

	// Inject APIs
	err = api.Inject(container)
	if err != nil {
		logger.Error("Failed to inject APIs", err)
	}

	return container
}
