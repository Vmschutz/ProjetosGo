package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Define a struct named "Person" with fields for name, age, and email.
//type Clientes struct {
//	Clientes []Cliente `json:"clientes"`
//}

type Cliente struct {
	Nome  string
	Cpf   int
	Email string
}

var Clientes []Cliente
var ListaClientes []Cliente

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/listacliente", getClientes).Methods("GET")
	r.HandleFunc("/cadastrocliente", postCliente).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil)) //--- log.Fatal() é um tipo de controle de excessão\
	//http.ListenAndServe(":8000", nil)  --> Ativa o servidor
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "*** Bolos da Vovó ***")
	fmt.Fprint(w, "Os metodos disponíveis são:")
	fmt.Fprint(w, "/listacliente")
	fmt.Fprint(w, "/cadastrocliente ")
}

func getClientes(w http.ResponseWriter, req *http.Request) {
	for i := 0; i < len(ListaClientes); i++ {
		fmt.Println("ID é:" + strconv.Itoa(i+1))
		fmt.Println("Usuario Nome: " + ListaClientes[i].Nome)
		fmt.Println("Usuario Email: " + ListaClientes[i].Email)
		fmt.Println("Usuario CPF:" + strconv.Itoa(ListaClientes[i].Cpf))
	}
	json.NewEncoder(w).Encode(ListaClientes)
}

func postCliente(w http.ResponseWriter, r *http.Request) {
	var newCliente Cliente

	err := json.NewDecoder(r.Body).Decode(&newCliente)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	Clientes = append(Clientes, newCliente)

	json.NewEncoder(w).Encode(newCliente)

	fmt.Println("O Range é:" + strconv.Itoa(len(Clientes)))

	ListaClientes = Clientes

	for ii := 0; ii < len(Clientes); ii++ {
		fmt.Println("ID é:" + strconv.Itoa(ii+1))
		fmt.Println("Usuario Nome: " + Clientes[ii].Nome)
		fmt.Println("Usuario Email: " + Clientes[ii].Email)
		fmt.Println("Usuario CPF:" + strconv.Itoa(Clientes[ii].Cpf))
	}

}
