package main

import "fmt"

type student struct {
	id    int
	name  string
	class string
}

//添加学生
func (s *studentMar) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

//编辑学生
func (s *studentMar) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id { //学号相同时即找到了要修改的学生
			s.allStudents[i] = newStu //覆盖赋值
			return
		}
	}
	fmt.Println("暂未该学生信息")
}

//展示学生
func (s *studentMar) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号： %d 姓名：%s 班级：%s\n", v.id, v.name, v.class)
	}
}

//学员管理类型
type studentMar struct {
	allStudents []*student
}

//newStudentMar 是studentMar的构造函数
func newStudentMar() *studentMar {
	return &studentMar{
		allStudents: make([]*student, 0, 100),
	}
}

//newStudent 是Student的构造函数
func newStudent(id int, name string, class string) *student {

	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}
