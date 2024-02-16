package controllers

import "net/http"

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
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
