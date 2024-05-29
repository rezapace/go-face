package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestData struct {
	Name     string `json:"name"`
	Encoding string `json:"encoding"`
}

func main() {
	e := echo.New()

	// Route to receive user input and send a POST request to Flask
	e.POST("/send_request", func(c echo.Context) error {
		requestData := new(RequestData)

		if err := c.Bind(requestData); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// Convert user input to JSON
		jsonData, err := json.Marshal(requestData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// Define the Flask API endpoint
		flaskURL := "http://localhost:5000/add_known_face" // Replace with your Flask API URL

		// Send a POST request to Flask with user input
		resp, err := http.Post(flaskURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer resp.Body.Close()

		// Read and display the response from Flask
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, string(responseData))
	})

	e.Start(":8080") // Change the port if needed
}
