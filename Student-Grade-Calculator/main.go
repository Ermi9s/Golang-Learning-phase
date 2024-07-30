package main

import (
	io "fmt"
	model "github.com/ermi9s/go/models"
	check "github.com/ermi9s/go/util"
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
		io.Println("\nCourse number: " , i+1);
		io.Print("\n ---------------------\n");

		var course_name string;
		var score float32;

		n := 1;
		for n > 0 {
			io.Println("Enter course Name: ");
			io.Scan(&course_name);
			if !check.Names(course_name) {
				io.Println("!!Enter a valid Name");
				continue
			}

			io.Println("Enter your score: ");
			io.Scan(&score);
			if !check.Course_Score(score) {
				io.Println("!!Enter a valid score");
				continue
			}
			n--;
		}
		io.Println("-------------------------------------------")
		var course model.Subject;
		course.Init_subject(course_name , score);
		courses = append(courses, course);
	}

	student.Init_student(student_name , course_count , courses);

	io.Println("\n\n**************************************************************************")
	io.Println("Student name: ", student_name);
	io.Println("Taken courses: " , course_count);
	io.Print("\n")
	var AvGrade float32;
	for i := 0; i < student.Number_of_subjects; i++ {
		io.Println("\tCourse Name: " , student.Subjets[i].Name)
		io.Println("\tCourse grade: " , student.Subjets[i].Grade);
		io.Println("\tGrade value: ", student.Subjets[i].Grade_value);
		AvGrade += student.Subjets[i].Grade_value;
		io.Println("\t-----------------------------------------------")
	}
	AvGrade /= float32(student.Number_of_subjects);
	io.Println("\t\tAvarage Grade: ", AvGrade)




}