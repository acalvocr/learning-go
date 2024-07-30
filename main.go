package main

import (
    "fmt"
    "io/ioutil"
    "log"
    sqlserver "main/database"
    models "main/models"
    utils "main/utils"
    "net/http"

    "github.com/gorilla/mux"
)

func getClientesFromDb()[] models.Cliente {
    var Clientes[] models.Cliente
    sqlserver.Database.Find( & Clientes)
    return Clientes
}

func allClientes(w http.ResponseWriter, r * http.Request) {
    utils.JsonResponse(w, getClientesFromDb())
}

func getCliente(w http.ResponseWriter, r * http.Request) {
    id: = mux.Vars(r)["idCliente"]

    var cliente models.Cliente

    result: = sqlserver.Database.Where("IdCliente = ?", id).First( & cliente)
    if result.Error == nil {
        utils.JsonResponse(w, cliente)
    }
}

func createCliente(w http.ResponseWriter, r * http.Request) {
    reqBody, _: = ioutil.ReadAll(r.Body)
    var newCliente models.Cliente

    utils.JsonDeserialize(reqBody, & newCliente)

    result: = sqlserver.Database.Create( & newCliente)
    fmt.Println(result.Error)

    utils.JsonResponse(w, models.BaseResult {
        Result: true,
        Message: "Cliente fue creado"
    })
}

func deleteCliente(w http.ResponseWriter, r * http.Request) {
    id: = mux.Vars(r)["idCliente"]
    var deletedCliente models.Cliente
    result: = sqlserver.Database.Where("IdCliente = ?", id).Delete(deletedCliente)
    fmt.Println(result.Error)

    utils.JsonResponse(w, models.BaseResult {
        Result: true,
        Message: "Cliente fue borrado"
    })
}

func handleRequests() {
    myrouter: = mux.NewRouter().StrictSlash(false)
    myrouter.HandleFunc("/clientes", allClientes).Methods("GET")
    myrouter.HandleFunc("/cliente/{id}", getCliente).Methods("GET")
    myrouter.HandleFunc("/cliente/{id}", deleteCliente).Methods("DELETE")
    myrouter.HandleFunc("/cliente", createCliente).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", myrouter))
}

func main() {
    sqlserver.Init()
    handleRequests()
}