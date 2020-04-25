package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20200425_013519 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20200425_013519{}
	m.Created = "20200425_013519"

	migration.Register("User_20200425_013519", m)
}

// Run the migrations
func (m *User_20200425_013519) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`encrypted_email` varchar(128) NOT NULL,`encrypted_password` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *User_20200425_013519) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
