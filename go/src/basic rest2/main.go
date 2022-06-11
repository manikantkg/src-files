// basic rest2 project main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Organisation string `json: "ORGANISATION"`
	Desc         string `json:"PURPOSE"`
	Content      string `json:"USAGE"`
	Number       string `json: "Number"`
}

type User struct {
	Firstname string `json: "Firstname"`
	Lastname  string `json: "Lastname"`
	Enterage  int    `json: 28`
	Address   string `json: "Adress"`
}

type Otp struct {
	Enterotp string
}
type Articles []Article
type Users []User
type Otps []Otp

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{

		Article{Organisation: "TTD ORGANISATION",
			Desc:    "ONLINE BOOKING",
			Content: "FOR DEVOTEES",
			Number:  "9346070012"},
	}
	fmt.Println("End point Hit : All Articles End point")

	json.NewEncoder(w).Encode(articles)

	fmt.Println("Rest API all are executed")
}
func allUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Firstname: "MANI",
			Lastname: "KANTA",
			Enterage: 28,
			Address:  "ANANTAPUR"},
		User{Firstname: "sharat",
			Lastname: "Gajula",
			Enterage: 25,
			Address:  "hyderabad"},
		User{Firstname: "suma",
			Lastname: "Latha",
			Enterage: 25,
			Address:  "Bangalore"},
	}

	fmt.Println("allow users executed")
	json.NewEncoder(w).Encode(users)
}
func allOtps(w http.ResponseWriter, r *http.Request) {
	otps := Otps{
		Otp{Enterotp: "664422"},
	}
	fmt.Println("OTP ENTERED")
	json.NewEncoder(w).Encode(otps)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page created by MANIKANTA")
	//fmt.Println(w, "Endpoint hit: homepage")
}
func handleRequest() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/users", allUsers)
	http.HandleFunc("/otps", allOtps)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}
