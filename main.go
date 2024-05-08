package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Date struct {
	year  int
	month time.Month
	day   int
}

type Task struct {
	Name       string
	IsFinished bool
}

type StudyDay struct {
	hours int
	tasks []Task
	date  Date
}

func getTodayDate() Date {
	year, month, day := time.Now().Date()
	return Date{year: year, month: month, day: day}
}

func (d *Date) toString() string {
	return fmt.Sprintf("%v %v %v", d.day, d.month, d.year)
}

func (s *StudyDay) writeable() []string {
	return []string{fmt.Sprintf("%v", s.hours), fmt.Sprintf("%v", s.tasks), s.date.toString()}
}

func main() {
	file, err := os.OpenFile("data.csv", os.O_CREATE|os.O_RDWR, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	day := StudyDay{hours: 5, tasks: nil, date: getTodayDate()}
	writer.Write([]string{"hours", "tasks", "date"})
	err = writer.Write(day.writeable())
	if err != nil {
		log.Fatal(err)
	}
}
