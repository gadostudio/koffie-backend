package main

import (
	"github.com/shaderboi/koffie-backend/api"
	"github.com/shaderboi/koffie-backend/api/settings"
)

func main() {

	settings.SetupGlobalMidtransConfig()
	settings.InitializeSnapClient()

	api.Routes()
}
