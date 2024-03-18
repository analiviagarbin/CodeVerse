package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser creates a new user repository
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User // user body
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// The repository is open, it receives the database
	repository := repositories.NewUserRepository(db)
	userID, err := repository.Create((user))
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Last user inserted: %d", userID)))
}

// SearchUser searches for a specific user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando usuário"))
}

// SearchUsers lists all registered users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

// UpdateUser updates a specific user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeleteUser deletes a specific user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
