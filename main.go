package main

import (
	"fmt"

	"github.com/DoonaaldArmando/go-library/init/message"
	"github.com/DoonaaldArmando/go-library/init/register"
)

func main() {

	gateway, _ := register.
		InitRegister().
		InitDataConsumers().
		InitDataServiceName("service").
		InitDataProject("project").
		HandleQuery("donald.torres", mostrar).
		ListenEvent("service", "project", "donald.armando", show).
		Build()

	gateway.SendQuery(
		"project",
		"service",
		"userAgent",
		"user",
		"donald.torres",
		"0011223344",
		show,
		persona{
			Name: "Armando",
			Age:  28,
		},
	)

	/*
		gateway.SendEvent("project",
			"service",
			"userAgent",
			"user",
			"donald.armando",
			"0011223344",
			persona{
				Name: "Armando",
				Age:  28,
			})*/

	for {
	}
}

func mostrar(m message.Message) (interface{}, string) {
	fmt.Println("Message para responder-> ", m)
	return persona{
			Name: "Arteaga",
			Age:  28,
		},
		"OtroUser"
}

func show(m message.Message) {
	fmt.Println("Message de respuesta -> ", m)
}

type persona struct {
	Name string
	Age  int
}
