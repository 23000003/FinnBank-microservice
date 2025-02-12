package main

import (
    "github.com/gin-gonic/gin"
    "graphql-service/utils"
    "graphql-service/routers"
)

func main() {
    router := gin.New()
    logger, err := utils.NewLogger()
    if err != nil {
        panic(err)
    }
    logger.Info("Starting the application...")

    // router group with base path for each services (to identify the service)
    serviceAPI := router.Group("/api/graphql-service")
    {
        // Use the group for your routes
        routers.GraphqlRouter(serviceAPI)
    }
    
    routers.InitializeSwagger(router)

    logger.Info("Server running on http://localhost:8083")

    if err := router.Run("localhost:8083"); err != nil {
        logger.Error("Failed to start server: %v", err)
    }
}