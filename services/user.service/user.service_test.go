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
		Name:      "Alexander",
		Email:     "alexrrojas.b@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := userService.Create(user)

	if err != nil {
		t.Error("Error en la prueba de persistencia de datos")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}

func TestRead(t *testing.T) {
	users, err := userService.Read()
	if err != nil {
		t.Error("Se ha presentado un error")
		t.Fail()
	}
	if len(users) == 0 {
		t.Error("No hay datos")
		t.Fail()
	}
	t.Log("La prueba finalizo correctamente")
}

func TestUpdate(t *testing.T) {
	user := m.User{
		Name:  "Alexander Rojas Benavides",
		Email: "proxtos@gmail.com",
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
