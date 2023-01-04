package models

// untuk struktur table users dengan gorm

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaLengkap string `gorm:"varchar(255)" json:"nama_lengkap"`
	Username    string `gorm:"varchar(255)" json:"username"`
	Password    string `gorm:"varchar(255)" json:"password"`
}
