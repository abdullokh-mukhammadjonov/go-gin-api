package models

type Person struct {
	ID      uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

// type Person struct {
// 	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
// 	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
// 	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
// 	Age       int    `json:"age" binding:"gte=1,lte=130"`
// 	Email     string `json:"email" binding:"required,email" gorm:"type:varchar(256);UNIQUE"`
// }
