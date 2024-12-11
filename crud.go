package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

type Usuario struct {
	ID		int		`json:"id"`		//	As tags json:"id" e json:"nome" especificam como os campos serão 
	Nome	string	`json:"nome"`	//	serializados/deserializados no formato JSON.
}

var usuarios []Usuario

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/usuarios", ListarUsuarios).Methods("GET")
    r.HandleFunc("/usuarios/{id}", ObterUsuario).Methods("GET")
    r.HandleFunc("/usuarios", CriarUsuario).Methods("POST")
    r.HandleFunc("/usuarios/{id}", AtualizarUsuario).Methods("PUT")
    r.HandleFunc("/usuarios/{id}", ExcluirUsuario).Methods("DELETE")

    fmt.Println("Servidor rodando na porta 3000...")
    http.ListenAndServe(":3000", r)


}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuarios)		//Converte a lista usuarios para JSON e a escreve na resposta HTTP
}

func ObterUsuario (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id,_ := strconv.Atoi(params["id"])

	for _, u := range usuarios {
		if u.ID == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
    var u Usuario
    json.NewDecoder(r.Body).Decode(&u)
    u.ID = len(usuarios) + 1
    usuarios = append(usuarios, u)
    json.NewEncoder(w).Encode(u)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for i, u := range usuarios {
        if u.ID == id {
            var atualizado Usuario
            json.NewDecoder(r.Body).Decode(&atualizado)
            atualizado.ID = id
            usuarios[i] = atualizado
            json.NewEncoder(w).Encode(atualizado)
            return
        }
    }
    http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for i, u := range usuarios {
        if u.ID == id {
            usuarios = append(usuarios[:i], usuarios[i+1:]...)
            fmt.Fprintf(w, "Usuário %d excluído com sucesso", id)
            return
        }
    }
    http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}
