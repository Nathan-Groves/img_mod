package Getpic

import (
    "fmt"
    "net/http"
    "bufio"
    "os"
    "io"
)

func GetImage() {
    
    // URL of the image you want to download

  	 reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give me a number")

	fmt.Print("-> ")

	text, _ := reader.ReadString('\n')
    
    // Create an HTTP GET request
    response, err := http.Get(text)
    if err != nil {
        fmt.Println("Error making the request:", err)
        return
    }
    defer response.Body.Close()

    // Check if the response status code is OK (200)
    if response.StatusCode != http.StatusOK {
        fmt.Println("Error: Status code", response.StatusCode)
        return
    }

    // Create a new file to save the image
    outputFile, err := os.Create("downloaded_image.jpg")
    if err != nil {
        fmt.Println("Error creating the file:", err)
        return
    }
    defer outputFile.Close()

    // Copy the HTTP response body to the file
    _, err = io.Copy(outputFile, response.Body)
    if err != nil {
        fmt.Println("Error saving the image:", err)
        return
    }

    fmt.Println("Image downloaded and saved as " + text)
}
