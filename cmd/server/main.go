package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/arturfil/comments-api/internal/transport/http"
)

// App - the struct with the application dependencies like poitners
// to db connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up application")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up the server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error estartin up our Reat APi")
		fmt.Println(err)
	}
}
