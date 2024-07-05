package main

import (
	"context"
	"fmt"
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/handlers"
	"libery_llm_chat_base_service/server"

	"github.com/Gerardo115pp/patriot_router"
	"github.com/Gerardo115pp/patriots_lib/echo"
)

func BinderRoutes(server server.Server, router *patriot_router.Router) {
	router.RegisterRoute(patriot_router.NewRoute("/alive", true), handlers.AliveHandler(server))
}

func main() {
	app_config.VerifyConfig()

	echo.Echo(echo.GreenFG, "Starting llm_chat_base_service")

	var new_server_config *server.ServerConfig = new(server.ServerConfig)
	new_server_config.Port = app_config.SERVICE_PORT

	echo.EchoDebug(fmt.Sprintf("server config: %+v", new_server_config))

	server, err := server.NewBroker(context.Background(), new_server_config)
	if err != nil {
		echo.EchoFatal(err)
	}

	server.Run(BinderRoutes)
}
