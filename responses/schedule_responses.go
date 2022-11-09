
package responses

import(
    "schedule_ms/models"
)

type ScheduleResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    []interface{}           `json:"data"`
}

type ScheduleResponseAll struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    []models.Schedule          `json:"data"`
}

