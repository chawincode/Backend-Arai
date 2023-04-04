package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Classroom struct {
	ID        primitive.ObjectID `bson:"_id"`
	ClassRoom string             `json:"classroom"`
}

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Nick_name     *string            `json:"nick_name"`
	Password      *string            `json:"password"`
	Email         *string            `json:"email"`
	Phone         *string            `json:"phone"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type"` // validate:"required,eq=ADMIN|eq=USER"
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
	Student_id    string             `json:"student_id"`
	Classrooms    []Classroom        `json:"classrooms"`
}
