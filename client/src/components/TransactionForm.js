import React, { useState, useEffect } from "react";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import TransactionTypeSelect from "./TransactionTypeSelect";
import Api from "../api";

function TransactionForm() {
  const [type, setType] = useState("credit");
  const [amount, setAmount] = useState("");
  const [balance, setBalance] = useState(0);

  function handleSubmit(event) {
    event.preventDefault();
    Api.commitTransaction({ transaction_type: type, amount }).then((result) =>
      setBalance(result.balance)
    );
  }

  function handleTextChange(event) {
    setAmount(parseFloat(event.target.value));
  }

  function handleSelectChange(value) {
    setType(value);
  }

  useEffect(() => {
    Api.getBalance().then((result) => setBalance(result.balance));
  });

  return (
    <div>
      <div>
        <h1>Balance</h1>
        <h2>$ {balance}</h2>
      </div>
      <form onSubmit={handleSubmit}>
        <TransactionTypeSelect value={type} onChange={handleSelectChange} />
        <TextField
          id="outlined-basic"
          label="Amount"
          variant="outlined"
          value={amount}
          onChange={handleTextChange}
        />
        <Button type="submit" variant="contained" color="primary">
          Submit
        </Button>
      </form>
    </div>
  );
}

export default TransactionForm;
