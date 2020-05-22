package main

import (
	"errors"
	"fmt"
	"log"
)

// Date struct.
type Date struct {
	Year  int
	Month int
	Day   int
}

// SetYear method.
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid year")
	}
	d.Year = year
	return nil
}

// SetMonth method.
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.Month = month
	return nil
}

// SetDay method.
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(4)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(17)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}
