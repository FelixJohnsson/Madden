import { PurchaseOrderList } from "../components/List";
import { useState, useEffect } from "react";
import * as t from "../types";
import Navigation from "../components/Navigation";

import {
  getPurchaseOrders,
  deletePurchaseOrder,
  createPurchaseOrder,
} from "../api/api";
import NewPurchaseOrderForm from "../components/NewPurchaseOrderForm";

const PurchaseOrderPage = () => {
  const [purchaseOrders, setPurchaseOrders] = useState<t.PurchaseOrder[]>([]);
  const [newPurchaseOrder, setNewPurchaseOrder] = useState<t.PurchaseOrder>({
    id: 0,
    itemId: 1,
    itemName: "",
    amount: 0,
    currency: "SEK",
    createdAt: new Date().toISOString(),
    status: "New",
    companyId: 1,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;

    setNewPurchaseOrder({
      ...newPurchaseOrder,
      [name]: value,
    });
  };

  const handleSubmit = () => {
    createPurchaseOrder(newPurchaseOrder);
  };

  useEffect(() => {
    getPurchaseOrders().then((purchaseOrders) => {
      setPurchaseOrders(purchaseOrders);
    });
  }, []);

  const handleDelete = (id: number) => {
    deletePurchaseOrder(id).then(() => {
      setPurchaseOrders(
        purchaseOrders.filter((purchaseOrder) => purchaseOrder.id !== id)
      );
    });
  };

  return (
    <>
      <Navigation />
      <div className="flex flex-col items-center justify-center mb-10">
        <h1 className="text-3xl font-bold">Purchase Order Page</h1>
      </div>

      <PurchaseOrderList data={purchaseOrders} handleDelete={handleDelete} />
      <NewPurchaseOrderForm
        newPurchaseOrder={newPurchaseOrder}
        handleChange={handleChange}
        handleSubmit={handleSubmit}
      />
    </>
  );
};

export default PurchaseOrderPage;
