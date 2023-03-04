package models

import "github.com/jinzhu/gorm"

type persona struct {
	gorm.Model

	ID        int64  `json:"id"`
	Nombre    string `json:"diego"`
	Apellido  string `json:"perez"`
	Direccion string `json:"cra 1000"`
	Telefono  string `json:"312456789"`
}
