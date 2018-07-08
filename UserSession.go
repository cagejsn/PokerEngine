package main

import (
	"time"
	"database/sql"
	"fmt"
	"errors"
	"github.com/satori/go.uuid"
)

type UserSession struct {
	user User
	sessionKey string
	expires time.Time
}


func makeUserSession(user User) UserSession {

	duration := time.Hour
	expiration := time.Now().Add(duration)
	sessionKey,_  := uuid.NewV4()
	userSession := UserSession{user: user , sessionKey: sessionKey.String() , expires: expiration}

	return userSession
}


func createUserSessionTable(db *sql.DB){
	stmt, err := db.Prepare("CREATE TABLE user_session ( email varchar(40), session_key varchar(40), expiration DATE , PRIMARY KEY (session_key));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("User Session Table successfully created....")
	}
}

func storeUserSession(db *sql.DB, session UserSession){

	stmt, err := db.Prepare("INSERT INTO user_session (email, session_key, expiration) VALUES(?,?,?);")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec(session.user.Email , session.sessionKey ,session.expires)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User Session, %s successfully created....",session)
	}

}


func retrieveUserSession(db *sql.DB, sessionKey string) (UserSession, error) {

	var existingUserSession UserSession
	row := db.QueryRow("select email, session_key, expiration from user_session where session_key = ?;", sessionKey)
	err := row.Scan(&existingUserSession.user.Email, &existingUserSession.sessionKey, &existingUserSession.expires)

	if err != nil {
		// If no results send null
		return existingUserSession, errors.New("no UserSession found")
	}

	return existingUserSession, nil
}