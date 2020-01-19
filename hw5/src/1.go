package main

import (
	"fmt"
	"time"
)

func retrieveNormalizedTimeFromDataSource() time.Time {
	// Normally you'd fetch the time from a specific datasource here.
	return time.Date(2016, time.April, 26, 6, 10, 0, 0, time.UTC)
}

func main() {
	fmt.Println("В методичку нужно добавить показ работы с часовыми поясами и парсинг даты из строки")

	// Make sure to handle the error ;), omitted for brevity
	userLocation, _ := time.LoadLocation("America/Panama")
	t := retrieveNormalizedTimeFromDataSource()
	t = t.In(userLocation)
	fmt.Println(t.Format(time.RFC822Z))

	td, _ := time.Parse("02-01-2006 15:04:05", "26-04-2016 08:10:23")
	fmt.Println(td.Format(time.RFC822Z))
}
