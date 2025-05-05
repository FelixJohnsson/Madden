import * as t from "../types";
import { Button, TextField } from "@mui/material";

const NewPurchaseOrderForm = ({
  newPurchaseOrder,
  handleChange,
  handleSubmit,
}: {
  newPurchaseOrder: t.PurchaseOrder;
  handleChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  handleSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
}) => {
  return (
    <div>
      <form onSubmit={handleSubmit}>
        <TextField
          id="item-name"
          label="Item Name"
          variant="outlined"
          name="itemName"
          value={newPurchaseOrder.itemName}
          onChange={handleChange}
        />
        <TextField
          id="amount"
          label="Amount"
          variant="outlined"
          name="amount"
          type="number"
          value={newPurchaseOrder.amount}
          onChange={handleChange}
        />
        <Button type="submit">Submit</Button>
      </form>
    </div>
  );
};

export default NewPurchaseOrderForm;
