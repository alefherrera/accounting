import React, { useState, useEffect } from "react";
import Api from "../api";

function TransactionItem({ transaction }) {
  return (
    <div>
      <h2>
        Type: {transaction.transaction_type} - Amount: $ {transaction.amount}
      </h2>
    </div>
  );
}

function TransactionList() {
  const [transactions, setTransactions] = useState([]);

  useEffect(() => {
    Api.getTransactions().then((result) => setTransactions(result || []));
  });

  return (
    <div>
      <h1>Transactions</h1>
      {transactions.map((transaction) => (
        <TransactionItem transaction={transaction} />
      ))}
    </div>
  );
}

export default TransactionList;
