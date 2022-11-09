package controllers

import (
    "context"
    "schedule_ms/configs"
    "schedule_ms/models"
    "schedule_ms/responses"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson" 
)

var scheduleCollection *mongo.Collection = configs.GetCollection(configs.DB, "schedules")
var valid = validator.New()

func CreateSchedule() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var schedule models.Schedule
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&schedule); err != nil {
            c.JSON(http.StatusBadRequest, responses.ScheduleResponse{Status: http.StatusBadRequest, Message: "error", Data: []interface{}{err.Error()}})
            return
        }
		
        //use the validator library to validate required fields
        if validationErr := valid.Struct(&schedule); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.ScheduleResponse{Status: http.StatusBadRequest, Message: "error", Data: []interface{}{validationErr.Error()}})
            return
        }
		
		newSchedule := models.Schedule{
			//Id:       primitive.NewObjectID(),
			Inicio:   schedule.Inicio ,
			Fin:      schedule.Fin,
			Lunes:    schedule.Lunes,
			Martes:   schedule.Martes,
			Miercoles:schedule.Miercoles,
			Jueves:   schedule.Jueves,
			Viernes:  schedule.Viernes,
			Sabado:   schedule.Sabado,
			Domingo:  schedule.Domingo,
			IdEstudiante:schedule.IdEstudiante, 
        }

        result, err := scheduleCollection.InsertOne(ctx, newSchedule)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ScheduleResponse{Status: http.StatusInternalServerError, Message: "error", Data:  []interface{}{err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.ScheduleResponse{Status: http.StatusCreated, Message: "success", Data:  []interface{}{result}})
    }
}

func GetASchedule() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        idEstudiante := c.Param("idEstudiante")
        var schedule models.Schedule
        defer cancel()
        err := scheduleCollection.FindOne(ctx, bson.M{"idestudiante": idEstudiante}).Decode(&schedule)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ScheduleResponse{Status: http.StatusInternalServerError, Message: "error", Data:  []interface{}{err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.ScheduleResponse{Status: http.StatusOK, Message: "success", Data: []interface{}{schedule}})
    }
}

func EditASchedule() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        idEstudiante := c.Param("idEstudiante")
        var schedule models.Schedule
        defer cancel()

        //validate the request body
        if err := c.BindJSON(&schedule); err != nil {
            c.JSON(http.StatusBadRequest, responses.ScheduleResponse{Status: http.StatusBadRequest, Message: "error", Data:  []interface{}{err.Error()}})
            return
        }

        //use the validator library to validate required fields
        if validationErr := valid.Struct(&schedule); validationErr != nil {
            c.JSON(http.StatusBadRequest, responses.ScheduleResponse{Status: http.StatusBadRequest, Message: "error", Data:  []interface{}{validationErr.Error()}})
            return
        }

        update := bson.M{
            "inicio": schedule.Inicio, 
            "fin": schedule.Fin, 
            "lunes": schedule.Lunes,
            "martes": schedule.Martes, 
            "miercoles": schedule.Miercoles, 
            "jueves": schedule.Jueves,
            "viernes": schedule.Viernes, 
            "sabado": schedule.Sabado, 
            "domingo": schedule.Domingo,
            "idestudiante": idEstudiante,
        }

        result, err := scheduleCollection.UpdateOne(ctx, bson.M{"idestudiante": idEstudiante}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ScheduleResponse{Status: http.StatusInternalServerError, Message: "error", Data:  []interface{}{err.Error()}})
            return
        }

        //get updated user details
        var updatedSchedule models.Schedule
        if result.MatchedCount == 1 {
            err := scheduleCollection.FindOne(ctx, bson.M{"idestudiante": idEstudiante}).Decode(&updatedSchedule)
            if err != nil {
                c.JSON(http.StatusInternalServerError, responses.ScheduleResponse{Status: http.StatusInternalServerError, Message: "error", Data:  []interface{}{err.Error()}})
                return
            }
        }

        c.JSON(http.StatusOK, responses.ScheduleResponse{Status: http.StatusOK, Message: "success", Data:  []interface{}{updatedSchedule}})
    }
}

func DeleteASchedule() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        idEstudiante := c.Param("idEstudiante")
        defer cancel()

        result, err := scheduleCollection.DeleteOne(ctx, bson.M{"idestudiante": idEstudiante})
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ScheduleResponse{Status: http.StatusInternalServerError, Message: "error", Data:  []interface{}{err.Error()}})
            return
        }

        if result.DeletedCount < 1 {
            c.JSON(http.StatusNotFound,
                responses.ScheduleResponse{Status: http.StatusNotFound, Message: "error", Data:  []interface{}{"Schedule with specified UserID not found!"}},
            )
            return
        }

        c.JSON(http.StatusOK,
            responses.ScheduleResponse{Status: http.StatusOK, Message: "success", Data:  []interface{}{ "Schedule successfully deleted!"}},
        )
    }
}


func GetAllSchedules() gin.HandlerFunc {
    return func(c *gin.Context) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var schedules []models.Schedule
        defer cancel()

        results, err := scheduleCollection.Find(ctx, bson.M{})

        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.ScheduleResponse{Status: http.StatusInternalServerError, Message: "error", Data:  []interface{}{ err.Error()} })
            return
        }

        //reading from the db in an optimal way
        defer results.Close(ctx)
        for results.Next(ctx) {
            var singleSchedule models.Schedule
            if err = results.Decode(&singleSchedule); err != nil {
                c.JSON(http.StatusInternalServerError, responses.ScheduleResponse{Status: http.StatusInternalServerError, Message: "error", Data:  []interface{}{ err.Error()}})
            }
            schedules = append(schedules, singleSchedule)
        }

        c.JSON(http.StatusOK,
            responses.ScheduleResponseAll{Status: http.StatusOK, Message: "success", Data: schedules },
        )
    }
}