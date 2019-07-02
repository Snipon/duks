package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

// Employee model.
type Employee struct {
	Name string
}

// Employees model.
type Employees struct {
	Employees []Employee
}

// GetDuks - Return name of employee.
func GetDuks() string {
	names := GetNames()
	count := len(names.Employees)
	week := getWeekNumber(time.Now())
	i := week % count
	if i == 0 {
		i = count
	}
	name := names.Employees[i-1].Name
	return name
}

// WeekNumber - Get week number.
func getWeekNumber(now time.Time) int {
	_, thisWeek := now.ISOWeek()
	return thisWeek
}

// Output - Generic output struct.
type Output struct {
	Name string `json:"name"`
}

// Handler - Lambda handler.
func Handler(events.APIGatewayProxyResponse) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       GetDuks(),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
