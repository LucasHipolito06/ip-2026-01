package main

import (
	"fmt"
	"log"
	"net/http"

	"projeto-saude-crud/app/handlers"
	"projeto-saude-crud/app/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado. Usando variáveis do sistema.")
	}

	db := utils.ConnectToDB()
	defer db.Close()

	utils.CreatePacientesTable(db)

	handlers.SetDB(db)

	http.HandleFunc("/", handlers.HomeHandler)

	http.HandleFunc("/pacientes", handlers.ListPacienteHandler)
	http.HandleFunc("/pacientes/create", handlers.CreatePacienteHandler)
	http.HandleFunc("/pacientes/update", handlers.UpdatePacienteHandler)
	http.HandleFunc("/pacientes/delete", handlers.DeletePacienteHandler)

	http.Handle("/forms/", http.StripPrefix("/forms/", http.FileServer(http.Dir("./static/forms"))))
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("./static/styles"))))

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}