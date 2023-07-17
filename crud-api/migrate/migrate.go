package main

import (
	"github.com/claytondixon/crud-api/initializers"
	"github.com/claytondixon/crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
