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

    router.Run("0.0.0.0:4000")
}
