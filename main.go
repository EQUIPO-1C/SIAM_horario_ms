package main

import (
    "schedule_ms/configs"
    "schedule_ms/routes" //add this
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    //run database
    configs.ConnectDB()

    //routes
    routes.ScheduleRoute(router)

    router.Run("localhost:4000")
}
