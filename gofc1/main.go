package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Route for uploading an image and performing face recognition
	e.POST("/upload", func(c echo.Context) error {
		// URL endpoint Flask Anda
		flaskURL := "http://localhost:5000/upload" // Ganti dengan URL Flask Anda

		// Membaca file gambar dari form-data
		file, err := c.FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "No file part"})
		}

		// Membuka file gambar
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Membuat buffer untuk menampung form-data
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// Membuat file form-data
		formFile, err := writer.CreateFormFile("image", file.Filename)
		if err != nil {
			return err
		}

		// Menyalin file gambar ke form-data
		_, err = io.Copy(formFile, src)
		if err != nil {
			return err
		}

		// Menambahkan form-data ke body
		writer.Close()

		// Membuat permintaan POST ke Flask
		resp, err := http.Post(flaskURL, writer.FormDataContentType(), body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Membaca respons JSON dari Flask
		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// Mengirim respons JSON kembali ke pengguna dengan status code yang sesuai
		return c.JSONBlob(resp.StatusCode, responseData)
	})

	e.Start(":8080") // Ganti dengan port yang sesuai
}
