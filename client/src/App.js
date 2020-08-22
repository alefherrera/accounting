import React from "react";
import logo from "./logo.svg";
import "./App.css";
import TransactionForm from "./components/TransactionForm";

function App() {
  return (
    <div className="App">
      <h1>Accounting notebook</h1>
      <TransactionForm />
    </div>
  );
}

export default App;
