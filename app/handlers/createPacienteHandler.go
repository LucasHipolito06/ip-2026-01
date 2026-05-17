package handlers

import (
	"net/http"
)

func CreatePacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "./static/forms/createPaciente.html")
		return
	}

	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")
	dataNascimento := r.FormValue("data_nascimento")
	telefone := r.FormValue("telefone")
	endereco := r.FormValue("endereco")
	convenio := r.FormValue("convenio")
	diagnostico := r.FormValue("diagnostico")

	if nome == "" || cpf == "" || dataNascimento == "" {
		http.Error(w, "Nome, CPF e data de nascimento são obrigatórios", http.StatusBadRequest)
		return
	}

	_, err := DB.Exec(`
		INSERT INTO pacientes 
		(nome, cpf, data_nascimento, telefone, endereco, convenio, diagnostico)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`,
		nome,
		cpf,
		dataNascimento,
		telefone,
		endereco,
		convenio,
		diagnostico,
	)

	if err != nil {
		http.Error(w, "Erro ao cadastrar paciente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
}