package main

import (
	"encoding/json"
	"github.com/alefherrera/accounting/api/domain/account"
	"github.com/alefherrera/accounting/api/domain/usecases"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	accountRepository := account.NewInmemoryRepository()
	commitTransaction := usecases.NewCommitTransactionImpl(accountRepository)
	getTransactions := usecases.NewGetTransactionsImpl(accountRepository)
	getTransactionById := usecases.NewGetTransactionByIdImpl(accountRepository)

	r.HandleFunc("/transactions", func(writer http.ResponseWriter, request *http.Request) {

		var input usecases.CommitTransactionInput
		_ = json.NewDecoder(request.Body).Decode(&input)

		result, err := commitTransaction.Execute(request.Context(), input)

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

	}).Methods(http.MethodPost)

	r.HandleFunc("/transactions/{id}", func(writer http.ResponseWriter, request *http.Request) {

		vars := mux.Vars(request)
		id := vars["id"]
		result, err := getTransactionById.Execute(request.Context(), id)

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

	}).Methods(http.MethodGet)

	r.HandleFunc("/transactions", func(writer http.ResponseWriter, request *http.Request) {

		result, err := getTransactions.Execute(request.Context())

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

	}).Methods(http.MethodGet)

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello!"))
	})

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("server listening on 8080")
}
