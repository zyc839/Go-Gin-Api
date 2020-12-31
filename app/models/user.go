package models

import "go-api/global"

type User struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

type UserSwagger struct {
	Lists []*User
	Total int
}

func GetUsers(pageNum int, pageSize int, maps interface{}) (users []User) {
	global.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users)

	return
}

func GetUser(maps interface{}) (user User, err error) {
	if err := global.DB.Where(maps).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByID(id int) (user User, err error) {
	if err := global.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func GetUserTotal(maps interface{}) (count int64) {
	global.DB.Model(&User{}).Where(maps).Count(&count)
	return
}

func ExistUserByMaps(maps interface{}) bool {
	var user User
	global.DB.Select("id").Where(maps).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func AddUser(Users map[string]interface{}) bool {
	user := User{
		Name:      Users["Name"].(string),
		CreatedBy: Users["CreatedBy"].(string),
	}
	if err := global.DB.Create(&user); err != nil {
		return true
	} else {
		return false
	}
}

func ExistTagByID(id int) bool {
	var user User
	global.DB.Select("id").Where("id = ?", id).First(&user)
	if user.ID > 0 {
		return true
	}

	return false
}

func DeleteUser(maps interface{}) (bool, error) {
	if err := global.DB.Where(maps).Delete(&User{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

func EditUser(id int, data interface{}) (bool, error) {
	if err := global.DB.Model(&User{}).Where("id = ?", id).Updates(data).Error; err != nil {
		return false, err
	}
	return true, nil
}
