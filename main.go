package main

import "fmt"

type Student struct {
	StudentId int
	FirstName string
	LastName  string
	Class     int
}

func (s Student) AboutStudent() {
	fmt.Printf("Personal Student info: \n First name: %s \n Last name: %s \n Class: %d ", s.FirstName, s.LastName, s.Class)
}

type Teacher struct {
	TeacherId int
	FirstName string
	LastName  string
	Salary    int
}

func (t Teacher) Grading(s Student, point int, course Course) Grades {
	return Grades{
		TeacherId: t.TeacherId,
		StudentId: s.StudentId,
		course:    course,
		Class:     s.Class,
		Grade:     point,
	}
}

type Grades struct {
	TeacherId int
	StudentId int
	course    Course
	Class     int
	Grade     int
}

func (g Grades) AboutGrade() {
	fmt.Printf("Personal Grade info: \n Course name: %s \n class: %d \n Grade: %d ", g.course.CourseName, g.Class, g.Grade)
}

type Course struct {
	CourseId   int
	CourseName string
}

func main() {
	student1 := Student{
		StudentId: 200103034,
		FirstName: "Meiirzhan",
		LastName:  "Uristemov",
		Class:     3,
	}
	//student2 := Student{
	//	StudentId: 200103035,
	//	FirstName: "Askar",
	//	LastName:  "Maratov",
	//	Class:     2,
	//}
	//student3 := Student{
	//	StudentId: 200103036,
	//	FirstName: "Berik",
	//	LastName:  "Nurov",
	//	Class:     1,
	//}

	teacher := Teacher{
		TeacherId: 100102021,
		FirstName: "Test",
		LastName:  "Test",
		Salary:    150000,
	}
	course := Course{
		CourseId:   101,
		CourseName: "Go programming",
	}
	student1.AboutStudent()
	var point int
	fmt.Print("\nEnter point for test student: ")
	fmt.Scan(&point)
	grade := teacher.Grading(student1, point, course)
	grade.AboutGrade()
}
