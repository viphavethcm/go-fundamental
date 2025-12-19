package models

import "fmt"

type Sex string

const (
	Male   Sex = "Male"
	Female Sex = "Female"
)

type Student struct {
	name          string
	sex           Sex
	mathMark      float64
	physicsMark   float64
	chemistryMark float64
	englishMark   float64
	averageMark   float64
}

func AddNewStudent(name string, sex int) *Student {
	defer fmt.Printf("Thêm sinh viên: %s thành công \n", name)
	if sex == 1 {
		return &Student{name: name, sex: Male}
	}
	return &Student{name: name, sex: Female}
}

func (student *Student) InputMark(mathMark, chemistryMark, physicsMark, englishMark float64) {
	student.mathMark = mathMark
	student.chemistryMark = chemistryMark
	student.physicsMark = physicsMark
	student.englishMark = englishMark
	student.averageMark = (mathMark + chemistryMark + physicsMark + englishMark) / 4
}

func (student *Student) GetName() string {
	return student.name
}

func (student *Student) GetSex() Sex {
	return student.sex
}

func (student *Student) GetAverageMark() float64 {
	return student.averageMark
}
