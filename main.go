package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/kabukky/httpscerts"
	"encoding/json"
	"hash/fnv"
	"errors"
)


func createUserTable(db *sql.DB){
	stmt, err := db.Prepare("CREATE TABLE user ( email varchar(40), first_name varchar(40), password_hash int(32) UNSIGNED , PRIMARY KEY (email));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Person Table successfully created....")
	}
}

func createNewUser(db *sql.DB, newUser *NewUser, passwordHash uint32){
	stmt, err := db.Prepare("INSERT INTO user ( email, password_hash, first_name) VALUES(?,?,?);")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec(newUser.Email, passwordHash ,newUser.FirstName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Person, %s successfully created....",newUser.FirstName)
	}
}

func openDatabaseConnection() (db *sql.DB, err error) {

	db, err = sql.Open("mysql",
		"cagejsn:password@tcp(127.0.0.1:3306)/pokerdb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}  else {
		fmt.Print("Database Connection Opened")
	}

	return
}

func findExistingUser(db *sql.DB, email string, passwordHash uint32) (User, error) {

	var existingUser User
	row := db.QueryRow("select email, first_name, password_hash from user where email = ?;", email)
	err := row.Scan(&existingUser.Email, &existingUser.FirstName, &existingUser.PasswordHash)

	if err != nil {
		// If no results send null
		panic(err)
	}

	if existingUser.PasswordHash != passwordHash {
		return User{}, errors.New("No User Found")
	}

	return existingUser, nil
}

var (
	// DBCon is the connection handle
	// for the database
	DBCon *sql.DB
)


func main() {

	var db *sql.DB
	db, _ = openDatabaseConnection()
	DBCon = db
	createUserTable(db)
	createUserSessionTable(db)

	hub := newHub()
	createdGameState := newGameState()
	dealer := new(Dealer)

	controller := &GameController{createdGameState,hub, dealer, false,make(map[string]*Player)}

	hub.stateController = controller
	go hub.run()

	hashPassword := func (s string) uint32 {
		h := fnv.New32a()
		h.Write([]byte(s))
		return h.Sum32()
	}

	http.HandleFunc( "/register", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)
		if(r.Method == http.MethodOptions){
			w.Write([]byte("good"))
			return
		}
		decoder := json.NewDecoder(r.Body)

		var newUser NewUser
		err := decoder.Decode(&newUser)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()
		createNewUser(db, &newUser, hashPassword(newUser.Password))
	})



	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		enableCors(&w)
		if(r.Method == http.MethodOptions){
			w.Write([]byte("good"))
			return
		}

		decoder := json.NewDecoder(r.Body)

		var loginAttempt struct {
			EmailToLookup string `json:"username"`
			Password string `json:"password"`
		}

		err := decoder.Decode(&loginAttempt)
		if err != nil {
			fmt.Print(err.Error())
			panic(err.Error())
		}

		var returnMessage []byte;
		defer r.Body.Close()

		user, err := findExistingUser(db, loginAttempt.EmailToLookup, hashPassword(loginAttempt.Password))
		if err != nil {
			//something went wrong or couldn't find user
			returnMessage = []byte("login unsuccessful")
			w.Write(returnMessage);
		} else {

			userSession := makeUserSession(user)
			storeUserSession(db, userSession)
			setUserSessionCookie(userSession, w)

			userToReturn := struct {
				Email string
				SessionKey string `json:"token"`
			}{
				user.Email,
				userSession.sessionKey,
			}

			userAsBytes, err := json.Marshal(userToReturn)
			if err != nil {
				fmt.Print(err.Error())
			}
			returnMessage = userAsBytes
			w.Write(returnMessage);
		}
	})





	//userPlayHandler := UserHandlerFunc(func(user User, w http.ResponseWriter,r *http.Request){
	//	//enableCors(&w)
	//	serveWs(hub, w, r)
	//})

	//unauthorizedHandler := http.HandlerFunc( func(w http.ResponseWriter,r *http.Request){
	//	var returnMessage []byte;
	//	returnMessage = []byte("unauthorized")
	//	w.Write(returnMessage);
	//})

	http.HandleFunc("/play", func(w http.ResponseWriter,r *http.Request){
		serveWs(hub, w, r)
	})
	startServerTLS()
	//http.ListenAndServe(":8080",nil)
}





func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:4201")
	//(*w).Header().Set("Access-Control-Allow-Origin", "*")

	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization")

}



func startServerTLS(){
	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8080")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


