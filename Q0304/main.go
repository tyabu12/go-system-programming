package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	filename := "test.txt"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=%v.zip", filename))
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	writer, err := zipWriter.Create(filename)
	if err != nil {
		panic(err)
	}
	io.Copy(writer, file)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
