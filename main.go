package main

import "fmt"
import "strings"

type Subject struct {
	Name  string
	Grade float64
}

type User struct {
	Name     string
	Subjects []Subject
}

func main() {
	user, err := takeInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user.display()
}

func takeInput() (*User, error) {
	var user User
	var subjectLen int
	fmt.Print("Your Name: ")
	scanned, err := fmt.Scanln(&user.Name)
	if err != nil {
		return nil, fmt.Errorf("invalid User Name")
	}
	if scanned == 0 {
		return nil, fmt.Errorf("name is required")
	}

	fmt.Print("Number of subjects: ")
	scanned, err = fmt.Scanf("%d", &subjectLen)
	if err != nil {
		return nil, fmt.Errorf("number of subjects is expected to be Integer")
	}
	if scanned == 0 {
		return nil, fmt.Errorf("number of subjects is required")
	}
	if subjectLen <= 0 {
		return nil, fmt.Errorf("invalid number of subjects")
	}

	var subject Subject
	fmt.Printf("Enter subject name and grade point %d times, \nEg. Philosophy 3.45\n", subjectLen)
	for subjectLen > 0 {
		scanned, err := fmt.Scanf("%s %f", &subject.Name, &subject.Grade)
		if err != nil {
			fmt.Println("⚠️ Invalid format or argument: Eg, Philosophy 3.45")
			continue
		}
		if scanned != 2 {
			fmt.Println("⚠️ Both subject and grade are required")
			continue
		}
		if subject.Grade <= 0 || subject.Grade > 100 {
			fmt.Println("⚠️  Invalid Grade range, Grade must be between 0 and 4")
			continue
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
