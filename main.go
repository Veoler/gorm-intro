package main

import (
	"github.com/Veoler/gorm-intro/gormintro"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

func main() {
	dsn := "host=localhost user=postgres password=4545 dbname=gorm-intro port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	} else {
		db.AutoMigrate(&gormintro.Student{}, &gormintro.Group{})
	}

	errCreGr := gormintro.CreateGroup(db, "Группа 5")
	if errCreGr != nil {
		fmt.Printf("Ошибка при создании группы: %v\n", errCreGr)
	} else {
		fmt.Println("Группа успешно создана!")
	}

	errCreSt := gormintro.CreateStudent(db, "Ибра", 19)
	if errCreSt != nil {
		fmt.Printf("Ошибка при создании студента: %v\n", errCreSt)
	} else {
		fmt.Println("Студент успешно создан!")
	}

	studentID := uint(5) // ID студента, которого хочу удалить
	errDel := gormintro.DeleteStudent(db, studentID)
	if errDel != nil {
		fmt.Printf("Не удалось удалить: %v\n", errDel)
	} else {
		fmt.Printf("Студент с ID %d успешно удален\n", studentID)
	}

	id := uint(1)
	student, errGet := gormintro.GetStudentByID(db, id)
	if errGet != nil {
		// Запись просто не найдена или это ошибка сети/базы?
		if errGet == gorm.ErrRecordNotFound {
			fmt.Printf("Студент с ID %d не существует\n", id)
		} else {
			fmt.Printf("Ошибка при поиске: %v\n", errGet)
		}
		return
	}

	// Если ошибок нет, можно обращаться к полям
	fmt.Printf("Найден студент: %s %d\n", student.Name, student.Age)

	errUp := gormintro.UpdateStudentName(db, 1, "Ибрагим")
		if errUp != nil {
			fmt.Printf("Ошибка обновления: %v\n", errUp)
		} else {
			fmt.Println("Имя успешно изменено!")
		}
	
}