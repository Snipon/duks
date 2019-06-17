package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Employee model.
type Employee struct {
	Name string "json:name"
}

// Employees model.
type Employees struct {
	Employees []Employee "json:employees"
}

// Message model.
type Message struct {
	Message string "json:message"
}

// Middleware globals.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// GetDuks - Return name of employee.
func GetDuks(w http.ResponseWriter) {
	var week = WeekNumber(time.Now())

	jsonFile, err := os.Open("employees.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var employees Employees

	json.Unmarshal(byteValue, &employees)
	var count = 0
	for i := 0; i < week; i++ {
		if count > len(employees.Employees) {
			count = 0
		}
		count++
	}

	var duks = employees.Employees[count-1].Name

	msg := Message{
		Message: fmt.Sprintf("It is week %d, and %s has the duks.", week, duks),
	}

	output, _ := json.Marshal(&msg)

	w.Write([]byte(output))

}

// GetWeekNumber - Get week number.
func WeekNumber(now time.Time) int {
	_, thisWeek := now.ISOWeek()
	return thisWeek
}

func main() {
	var Router = chi.NewRouter()

	// Router.
	Router.Use(middleware.RequestID)
	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)
	Router.Use(middleware.Timeout(60 * time.Second))
	Router.Use(Middleware)

	Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg := Message{
			Message: fmt.Sprintf("Yes, this is API."),
		}

		output, _ := json.Marshal(&msg)

		w.Write([]byte(output))
	})

	Router.Get("/duks", func(w http.ResponseWriter, r *http.Request) {
		GetDuks(w)
	})

	fmt.Println(http.ListenAndServe(":3000", Router))

}
