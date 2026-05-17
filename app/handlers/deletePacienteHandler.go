package handlers

import (
	"net/http"
)

func DeletePacienteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "./static/forms/deletePaciente.html")
		return
	}

	id := r.FormValue("id")

	if id == "" {
		http.Error(w, "ID do paciente é obrigatório", http.StatusBadRequest)
		return
	}

	_, err := DB.Exec(`
		DELETE FROM pacientes
		WHERE id = $1
	`, id)

	if err != nil {
		http.Error(w, "Erro ao excluir paciente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
}