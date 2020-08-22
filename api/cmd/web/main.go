package main

import (
	"encoding/json"
	"github.com/alefherrera/accounting/api/domain/account"
	"github.com/alefherrera/accounting/api/domain/usecases"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	r := mux.NewRouter()

	accountRepository := account.NewInmemoryRepository()
	commitTransaction := usecases.NewCommitTransactionImpl(accountRepository)
	getTransactions := usecases.NewGetTransactionsImpl(accountRepository)
	getTransactionById := usecases.NewGetTransactionByIdImpl(accountRepository)
	getBalanceImpl := usecases.NewGetBalanceImpl(accountRepository)

	r.HandleFunc("/transactions", func(writer http.ResponseWriter, request *http.Request) {
		var input usecases.CommitTransactionInput
		err := json.NewDecoder(request.Body).Decode(&input)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
			return
		}
		result, err := commitTransaction.Execute(request.Context(), input)
		sendResponse(writer, result, err)
	}).Methods(http.MethodPost)

	r.HandleFunc("/transactions/{id}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id := vars["id"]
		result, err := getTransactionById.Execute(request.Context(), id)
		sendResponse(writer, result, err)
	}).Methods(http.MethodGet)

	r.HandleFunc("/transactions", func(writer http.ResponseWriter, request *http.Request) {
		result, err := getTransactions.Execute(request.Context())
		sendResponse(writer, result, err)
	}).Methods(http.MethodGet)

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		result, err := getBalanceImpl.Execute(request.Context())
		sendResponse(writer, result, err)
	}).Methods(http.MethodGet)

	http.Handle("/", handlers.CORS()(r))

	go func() {
		err := http.ListenAndServe(":8080", nil)

		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Println("server listening on 8080")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func sendResponse(writer http.ResponseWriter, result interface{}, err error) {
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	if result == nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found"))
		return
	}

	json.NewEncoder(writer).Encode(result)
}
