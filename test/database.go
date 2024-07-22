package test

import "gorm.io/gorm"

func TruncateUser(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
}

func CreateRootUser(db *gorm.DB) {
	db.Exec("INSERT INTO users (username, password, email, role) VALUES ('root', 'root@email.com', 'root', 'ADMIN')")
}
