package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	// Create
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	// Read All
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.SearchUsers,
		AuthRequired: false,
	},
	// Read One
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Function:     controllers.SearchUser,
		AuthRequired: false,
	},
	// Update
	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	},
	// Delete
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	},
}
