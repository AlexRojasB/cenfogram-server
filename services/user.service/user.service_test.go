package user_service_test

import (
	"testing"
	"time"

	m "github.com/AlexRojasB/go-mongoAtlas-connection.git/models"
	userService "github.com/AlexRojasB/go-mongoAtlas-connection.git/services/user.service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userId string

func TestCreate(t *testing.T) {
	var old primitive.ObjectID
	userId = old.Hex()
	user := m.User{
		ID:        old,
		Nick:      "xzaokyx",
		Email:     "proxtos@gmail.com",
		Password:  "admin123.",
		UpdatedAt: time.Now(),
	}
	err := userService.Create(user)

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}

func TestRead(t *testing.T) {
	loginUser := m.User{
		Email:    "proxtos@gmail.com",
		Password: "admin13.",
	}

	user, err := userService.Read(loginUser)
	if err != nil {
		t.Error(err.Error())
		t.Fail()
	} else if user.Email == "" {
		t.Error("No hay datos")
		t.Fail()
	} else {
		t.Log("La prueba finalizo correctamente")
	}
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Nick:     "xzaokyx",
		Password: "admin123.3",
	}

	err := userService.Update(user, userId)
	if err != nil {
		t.Error("Se ha presentado un error")
		t.Fail()
	}
	t.Log("La prueba finalizo correctamente")
}

func TestDelete(t *testing.T) {
	err := userService.Delete(userId)
	if err != nil {
		t.Error("Se ha presentado un error")
		t.Fail()
	}
	t.Log("La prueba finalizo correctamente")
}
