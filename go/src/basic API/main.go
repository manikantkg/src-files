// basic API project main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json: "Titile"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

type User struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Phnum     string `json:"Number"`
}
type Otp struct {
	EnterOtp string `json:"Otp"`
}
type Token struct {
	Token string `json:"Token"`
}
type Articles = []Article // =
type Users []User
type Otps []Otp
type Tokens []Token

func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{

		Article{Title: "Test Title", Desc: "Test Description", Content: "Manikanta"},

		Article{Title: "Exam Timetable", Desc: "Schedule", Content: "Final Round"},
	}
	fmt.Println("End point Hit : All Articles End point")
	json.NewEncoder(w).Encode(articles)
}
func allUsers(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Firstname: "Mani", Lastname: "Kanta", Phnum: "9346070012"},
	}
	fmt.Println("Users details accepted")
	json.NewEncoder(w).Encode(users)
}
func allOtps(w http.ResponseWriter, r *http.Request) {
	otps := Otps{
		Otp{EnterOtp: "664422"},
	}
	fmt.Println("OTP ENTERD")
	json.NewEncoder(w).Encode(otps)
}

func allTokens(w http.ResponseWriter, r *http.Request) {
	tokens := Tokens{
		Token{Token: "0x6df3456"},
	}
	fmt.Println("Token generated")
	json.NewEncoder(w).Encode(tokens)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page created by MANIKANTA")
	//fmt.Println(w, "Endpoint hit: homepage")

}
func handleRequest() {
	http.HandleFunc("/", homepage)

	http.HandleFunc("/articles", allArticles)

	http.HandleFunc("/users", allUsers)

	http.HandleFunc("/otp", allOtps)

	http.HandleFunc("/token", allTokens)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequest()
}

//{"page":null,"per_page":10,"total":0,"total_pages":0,"data":[]}
//https://jsonmock.hackerrank.com/api/articles?author=%3CauthorName%3E&page=%3Cnum%3E
