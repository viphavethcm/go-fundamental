package main

import (
	"fmt"
	"fundamental/common"
	"fundamental/models"
)

var storages map[int]*models.Student

func main() {
	storages = make(map[int]*models.Student)
	cnt := 1
	for {
		showMenu()
		choice := common.ReadIntInput("Chọn chức năng: ")
		switch choice {
		case 1:
			showStudent()
		case 2:
			student := addStudent()
			storages[cnt] = student
			cnt++
		case 3:
			removeStudent()
		case 4:
			inputMark()
		case 5:
			fmt.Println("Xin tạm biệt")
			return
		}
	}
}

func showMenu() {
	fmt.Println("---- Phần mềm quản lý sinh viên ----")
	fmt.Println("1. Hiển thị sinh viên")
	fmt.Println("2. Thêm mới sinh viên")
	fmt.Println("3. Xoá sinh viên")
	fmt.Println("4. Nhập điểm học sinh")
	fmt.Println("5. Thoát")
}

func showStudent() {
	if len(storages) == 0 {
		fmt.Println("Danh sách hiện tại đang trống")
	}
	for id, student := range storages {
		fmt.Printf("MSSV: %d  |  Tên: %s   |  Giới tính:%s  |  Điểm trung bình: %f\n", id, student.GetName(), student.GetSex(), student.GetAverageMark())
	}
}

func addStudent() *models.Student {
	var name = common.ReadStringInput("Nhập tên: ")
	var sex = common.ReadIntInput("Nhập giới tính(1. Nam - 2. Nữ): ")
	student := models.AddNewStudent(name, sex)
	return student
}

func removeStudent() {
	var id = common.ReadIntInput("Nhập id sinh viên cần xoá: ")
	_, isContain := storages[id]
	if isContain {
		delete(storages, id)
		fmt.Printf("Xoá sinh viên id: %d thành công\n", id)
	} else {
		fmt.Printf("Không tìm thấy thông tin sinh viên cần xoá: %d\n", id)
		return
	}
}

func inputMark() {
	var id = common.ReadIntInput("Nhập id sinh viên cần thao tác: ")
	student, isContains := storages[id]
	if isContains {
		var mathMark = common.ReadFloatInput("Nhập điểm toán: ")
		var chemistryMark = common.ReadFloatInput("Nhập điểm hoá: ")
		var physicsMark = common.ReadFloatInput("Nhập điểm vật lý: ")
		var englishMark = common.ReadFloatInput("Nhập điểm anh: ")
		student.InputMark(mathMark, chemistryMark, physicsMark, englishMark)
	}
}
