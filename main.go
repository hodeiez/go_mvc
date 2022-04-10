package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	c "hodeinaiz/go_mvc/controllers"
	"hodeinaiz/go_mvc/models"
)

func main() {

	// Start the MVC App
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	app.Model.Register(&models.Message{})

	// CReate Models tables if they dont exist yet
	app.Model.AutoMigrateAll()

	// Register Controller
	app.AddController(c.NewMessage)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	log.Fatal(http.ListenAndServe(port, app))
}
