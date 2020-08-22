import React from "react";
import MenuItem from "@material-ui/core/MenuItem";
import Select from "@material-ui/core/Select";
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";

function TransactionTypeSelect({ value, onChange }) {
  const handleChange = (event) => {
    onChange(event.target.value);
  };

  return (
    <FormControl>
      <InputLabel id="demo-simple-select-label">Transaction Type</InputLabel>
      <Select
        labelId="demo-simple-select-label"
        value={value}
        onChange={handleChange}
      >
        <MenuItem value={"credit"}>Credit</MenuItem>
        <MenuItem value={"debit"}>Debit</MenuItem>
      </Select>
    </FormControl>
  );
}

export default TransactionTypeSelect;
