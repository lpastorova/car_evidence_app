package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Car struct {
	Spz         string `json:"spz"`
	Description string `json:"description"`
}

func getCarHandler(w http.ResponseWriter, _ *http.Request) {
	//Convert the "cars" variable to json

	cars, err := store.GetCar()

	carListBytes, err := json.Marshal(cars)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of cars to the response
	_, err = w.Write(carListBytes)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		return
	}
}

func createCarHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Car
	car := Car{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the car from the form info
	car.Spz = r.Form.Get("spz")
	car.Description = r.Form.Get("description")

	err = store.CreateCar(&car)
	if err != nil {
		fmt.Println("Err creating car, ", err)
	}

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
