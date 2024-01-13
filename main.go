package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func handlePostUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON data from the request body into a User struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// Process the user data as needed (e.g., store in a database)
	fmt.Printf("Received user data: %+v\n", user)

	// Send a response back to the client
	response := map[string]string{"message": "User data received successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	sendMail(user)
}

func handleRequests() {
	http.HandleFunc("/postUser", handlePostUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sendMail(user User) {
	godotenv.Load()
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	stmpAuthor := os.Getenv("SMTP_AUTHOR")

	auth := smtp.PlainAuth(
		"",
		stmpAuthor,
		smtpPassword,
		smtpHost,
	)

	msg := "From: " + stmpAuthor + "\n" + "To: " + user.Email + "\n" + "Subject: Welcome to the club!\n\n" + "Hello " + user.FirstName + " " + user.LastName + ",\n\n" + "Welcome to the club!\n\n" + "Best regards,\n" + "Matthew Araujo"

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		stmpAuthor,
		[]string{user.Email},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	handleRequests()
}
