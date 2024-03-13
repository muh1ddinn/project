package storage

import "lms_backed_pr/model"

type IStorage interface {
	CloseDB()
	Student() IStudentStorage
}

// ////
type IStudentStorage interface {
	GetallStudent(string) (model.GetAllstudent, error)
	//UpdateStudent(car model.Student) (string, error)
	DeleteStudent(string) error
	CreateStudent(student model.Student) (string, error)
	GetByIDStudent(string) ([]model.Student, error)
}
