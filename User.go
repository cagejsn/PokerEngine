package main

type User struct {
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	PasswordHash uint32 `json:password`
}

type NewUser struct {
	Email        string `json:"email"`
	FirstName    string `json:"firstName"`
	Password string `json:password`
}

//"CREATE TABLE user (id int NOT NULL AUTO_INCREMENT,
// email varchar(40),
// first_name varchar(40),
// password_hash varchar(40),
// d PRIMARY KEY (id))