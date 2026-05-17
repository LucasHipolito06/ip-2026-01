package handlers

import (
	"net/http"
)

func UpdatePacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "./static/forms/updatePaciente.html")
		return
	}

	id := r.FormValue("id")
	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")
	dataNascimento := r.FormValue("data_nascimento")
	telefone := r.FormValue("telefone")
	endereco := r.FormValue("endereco")
	convenio := r.FormValue("convenio")
	diagnostico := r.FormValue("diagnostico")

	if id == "" || nome == "" || cpf == "" || dataNascimento == "" {
		http.Error(w, "ID, nome, CPF e data de nascimento são obrigatórios", http.StatusBadRequest)
		return
	}

	_, err := DB.Exec(`
		UPDATE pacientes
		SET 
			nome = $1,
			cpf = $2,
			data_nascimento = $3,
			telefone = $4,
			endereco = $5,
			convenio = $6,
			diagnostico = $7
		WHERE id = $8
	`,
		nome,
		cpf,
		dataNascimento,
		telefone,
		endereco,
		convenio,
		diagnostico,
		id,
	)

	if err != nil {
		http.Error(w, "Erro ao atualizar paciente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
}