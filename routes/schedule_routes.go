package routes

import (
    "schedule_ms/controllers" //add this
    "github.com/gin-gonic/gin"
)

func ScheduleRoute(router *gin.Engine)  {
    router.POST("/schedule", controllers.CreateSchedule()) 
    router.GET("/schedule/:idEstudiante", controllers.GetASchedule()) 
    router.PUT("/schedule/:idEstudiante", controllers.EditASchedule()) 
    router.DELETE("/schedule/:idEstudiante", controllers.DeleteASchedule())
    router.GET("/schedules", controllers.GetAllSchedules())
}
