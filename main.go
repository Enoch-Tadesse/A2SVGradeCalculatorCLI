package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Subject struct {
	Name  string
	Grade float64
}

type User struct {
	Name     string
	Subjects []Subject
}

func main() {
	user, err := takeUserInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user.display()
}

func getNonEmptyInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read input")
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Input can not be empty")
		}
		return input
	}
}

func takeUserInput() (*User, error) {
	var user User
	var subjectLen int
	name := getNonEmptyInput("Your Name: ")
	user.Name = name

	for {
		countStr := getNonEmptyInput("Number of subjects: ")
		count, err := strconv.Atoi(countStr)
		if err != nil {
			fmt.Println("Input must be an integer")
			continue
		}
		subjectLen = count
		break
	}

	var subject Subject
	for subjectLen > 0 {
		name := getNonEmptyInput("Subject Name: ")
		subject.Name = name
		for {
			gradeStr := getNonEmptyInput("Grade: ")
			grade, err := strconv.ParseFloat(gradeStr, 64)
			if err != nil {
				fmt.Println("grade must in range of 0 and 100")
				continue
			}
			if grade <= 0 || grade > 100 {
				fmt.Println("⚠️  Invalid Grade range, Grade must be between 0 and 100")
				continue
			}
			subject.Grade = grade
			break
		}
		user.Subjects = append(user.Subjects, subject)
		subjectLen -= 1
	}
	return &user, nil
}

func (u *User) display() {
	fmt.Printf("\nStudent: %s\n", u.Name)
	fmt.Printf("%-15s | %-6s\n", "Subject", "Grade")
	fmt.Println(strings.Repeat("-", 25))

	for _, sub := range u.Subjects {
		fmt.Printf("%-15s | %6.2f\n", sub.Name, sub.Grade)
	}

	average := u.calculateAverage()
	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("%-15s | %6.2f\n", "Average", average)
}

func (u *User) calculateAverage() float64 {
	var total float64
	var count int

	for _, sub := range u.Subjects {
		total += sub.Grade
		count += 1
	}

	// incase there is subject
	if count == 0 {
		return 0
	}
	return total / float64(count)
}
