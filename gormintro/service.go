package gormintro

import (
	"gorm.io/gorm"
	"fmt"
)

func CreateGroup(db *gorm.DB, name string) error {
	// Создаем экземпляр структуры с переданным именем
	group := Group{Name: name}

	// Метод .Create выполняет INSERT запрос в базу данных
	result := db.Create(&group)

	// Возвращаем ошибку, если она возникла при записи
	return result.Error
}

func CreateStudent(db *gorm.DB, name string, age int) error {
	student := Student{Name: name, Age: age}

	result := db.Create(&student)

	return result.Error
}

func DeleteStudent(db *gorm.DB, id uint) error {
	result := db.Delete(&Student{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("студент с ID %d не найден", id)
	}

	return nil
}

func GetStudentByID(db *gorm.DB, id uint) (*Student, error) {
	st := Student{}
	
	result := db.First(&st, id)

	return &st, result.Error
}

func UpdateStudentName(db *gorm.DB, id uint, newName string) error {
	
	result := db.Model(&Student{}).Where("id = ?", id).Update("Name", newName)
	// Проверяем на ошибки выполнения запроса
	if result.Error != nil {
		return result.Error
	}

	// Проверяем, была ли обновлена хоть одна строка
	if result.RowsAffected == 0 {
		return fmt.Errorf("студент с ID %d не найден", id)
	}

	return nil
}



