package main

import (
	io "fmt"
	model "github.com/ermi9s/go/models"
)

func main() {
	var student_name string;
	var course_count int;
	var courses []model.Subject;
	var student model.Student;

	io.Println("\t\tHello");
	io.Print("Please Enter your name: ");
	io.Scan(&student_name);
	io.Println("\nNice to meet you",student_name);
	io.Print("Enter the number of courses you have taken: ");
	io.Scan(&course_count);

	println("\nGreat now enter each course's information");

	for i := range course_count {
		io.Println("Course number: " , i+1);
		io.Print("\n");

		var course_name string;
		var score float32;
		
		n := 2
		for n > 0{
			io.Print("Enter course Name: ");
			io.Scan(&course_name);
			io.Print("\nEnter your score: ");
			io.Scan(&score);
			n--;
		}
		var course model.Subject;
		course.Init_subject(course_name , score);
		courses = append(courses, course);
	}

	student.Init_student(student_name , course_count , courses);

	//view
		
}