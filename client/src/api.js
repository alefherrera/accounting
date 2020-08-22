const API_URL = "http://localhost:8080/";

function getBalance() {
  return fetch(API_URL).then((res) => res.json());
}

function commitTransaction(transaction) {
  return fetch(API_URL + "transactions", {
    method: "POST",
    body: JSON.stringify(transaction),
  }).then((res) => res.json());
}

function getTransactionById(id) {
  return fetch(API_URL + "transactions/" + id).then((res) => res.json());
}

function getTransactions() {
  return fetch(API_URL + "transactions").then((res) => res.json());
}

const Api = {
  getBalance,
  commitTransaction,
  getTransactions,
  getTransactions,
};

export default Api;
