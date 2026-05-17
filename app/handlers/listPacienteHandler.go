package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

var DB *sql.DB

type Paciente struct {
	ID             int
	Nome           string
	CPF            string
	DataNascimento string
	Telefone       string
	Endereco       string
	Convenio       string
	Diagnostico    string
}

func SetDB(db *sql.DB) {
	DB = db
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func ListPacienteHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query(`
		SELECT 
			id, 
			nome, 
			cpf, 
			data_nascimento, 
			telefone, 
			endereco, 
			convenio, 
			diagnostico
		FROM pacientes
		ORDER BY id ASC
	`)
	if err != nil {
		http.Error(w, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var pacientes []Paciente

	for rows.Next() {
		var p Paciente

		err := rows.Scan(
			&p.ID,
			&p.Nome,
			&p.CPF,
			&p.DataNascimento,
			&p.Telefone,
			&p.Endereco,
			&p.Convenio,
			&p.Diagnostico,
		)

		if err != nil {
			http.Error(w, "Erro ao ler paciente", http.StatusInternalServerError)
			return
		}

		pacientes = append(pacientes, p)
	}

	tmpl := `
	<!DOCTYPE html>
	<html lang="pt-br">
	<head>
		<meta charset="UTF-8">
		<title>Lista de Pacientes</title>
		<link rel="stylesheet" href="/styles/style.css">
	</head>
	<body>
		<div class="container">
			<h1>Pacientes cadastrados</h1>

			<a href="/" class="btn">Voltar ao início</a>
			<a href="/forms/createPaciente.html" class="btn">Cadastrar paciente</a>

			<table>
				<thead>
					<tr>
						<th>ID</th>
						<th>Nome</th>
						<th>CPF</th>
						<th>Data de nascimento</th>
						<th>Telefone</th>
						<th>Endereço</th>
						<th>Convênio</th>
						<th>Diagnóstico</th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{.ID}}</td>
						<td>{{.Nome}}</td>
						<td>{{.CPF}}</td>
						<td>{{.DataNascimento}}</td>
						<td>{{.Telefone}}</td>
						<td>{{.Endereco}}</td>
						<td>{{.Convenio}}</td>
						<td>{{.Diagnostico}}</td>
					</tr>
					{{else}}
					<tr>
						<td colspan="8">Nenhum paciente cadastrado.</td>
					</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</body>
	</html>
	`

	t, err := template.New("pacientes").Parse(tmpl)
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, pacientes)
	if err != nil {
		fmt.Println("Erro ao executar template:", err)
	}
}