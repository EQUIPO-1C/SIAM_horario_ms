
package models

import (
	
)
type Schedule struct {
    Inicio   string   `json:"inicio" validate:"required"`
    Fin 	string        `json:"fin" validate:"required"`
    Lunes    []Subject           `json:"lunes" validate:"required"`
	Martes    []Subject           `json:"martes" validate:"required"`
	Miercoles    []Subject           `json:"miercoles" validate:"required"`
	Jueves    []Subject           `json:"jueves" validate:"required"`
	Viernes    []Subject           `json:"viernes" validate:"required"`
	Sabado   []Subject           `json:"sabado" validate:"required"`
	Domingo    []Subject           `json:"domingo" validate:"required"`
	IdEstudiante string    `json:"idestudiante" validate:"required"`
}

type Subject struct {
	CodMateria string `json:"codmateria" validate:"required"`
	Materia string `json:"materia" validate:"required"`
	Plan string `json:"plan" validate:"required"`
	Horas string `json:"horas" validate:"required"`
	Grupo int `json:"grupo" validate:"required"`
	Salon string `json:"salon" validate:"required"`
} 