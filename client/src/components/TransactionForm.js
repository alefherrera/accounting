import React from "react";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import TransactionTypeSelect from "./TransactionTypeSelect";
import Api from "../api";

function TransactionForm() {
  const [type, setType] = React.useState("credit");
  const [amount, setAmount] = React.useState("");

  function handleSubmit(event) {
    event.preventDefault();
    Api.commitTransaction({ transaction_type: type, amount });
  }

  function handleTextChange(event) {
    setAmount(parseFloat(event.target.value));
  }

  function handleSelectChange(value) {
    setType(value);
  }

  return (
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
  );
}

export default TransactionForm;
