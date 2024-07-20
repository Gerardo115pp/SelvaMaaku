package main

import (
	"context"
	"fmt"
	app_config "libery_llm_chat_base_service/Config"
	"libery_llm_chat_base_service/database"
	qdrant_db "libery_llm_chat_base_service/database/qdrant"
	"libery_llm_chat_base_service/handlers"
	"libery_llm_chat_base_service/middleware"
	"libery_llm_chat_base_service/repository"
	"libery_llm_chat_base_service/server"

	"github.com/Gerardo115pp/patriot_router"
	"github.com/Gerardo115pp/patriots_lib/echo"
)

func BinderRoutes(server server.Server, router *patriot_router.Router) {
	router.RegisterRoute(patriot_router.NewRoute("/alive", true), handlers.AliveHandler(server))
	if app_config.CHAT_ENABLED {
		router.RegisterRoute(patriot_router.NewRoute("^/chat-claims.*", false), handlers.ChatClaimsHandler(server))
		router.RegisterRoute(patriot_router.NewRoute("/chat", true), middleware.CheckAuthCookie(handlers.ChatHandler(server)))
		router.RegisterRoute(patriot_router.NewRoute("/knowledge", true), handlers.KnowledgeHandler(server)) // Add domain secret auth(aka request most be sent from user with max level of trust)
	}
}

func main() {
	app_config.VerifyConfig()

	echo.Echo(echo.GreenFG, "Starting llm_chat_base_service")

	var new_server_config *server.ServerConfig = new(server.ServerConfig)
	new_server_config.Port = app_config.SERVICE_PORT

	echo.EchoDebug(fmt.Sprintf("server config: %+v", new_server_config))

	knowledge_repository, err := qdrant_db.NewChatKnowledgeRepository(app_config.KNOWLEDGE_QDRANT_REST_API_URL)
	if err != nil {
		echo.EchoFatal(err)
	}

	chat_repository, err := database.NewChatDatabase()
	if err != nil {
		echo.EchoFatal(err)
	}

	repository.SetKnowledgeRepository(knowledge_repository)
	repository.SetChatRepositoryImplementation(&chat_repository)

	server, err := server.NewBroker(context.Background(), new_server_config)
	if err != nil {
		echo.EchoFatal(err)
	}

	server.Run(BinderRoutes)
}
