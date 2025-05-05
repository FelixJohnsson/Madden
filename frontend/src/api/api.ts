import * as t from "../types";

const PORT = 8080;
const URL = `http://localhost:${PORT}/api`;

/* 
| Page            | Read | Create | Update | Delete |
| --------------- | ---- | ------ | ------ | ------ |
| Sales           | ✅    | ❌      | ❌      | ❌      |
| Purchase Orders | ✅    | ✅      | ❌      | ✅      |
*/

// --------------- Sales ---------------
export const getSales = async (): Promise<t.Sale[]> => {
  try {
    const response = await fetch(`${URL}/sales`);

    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`);
    }

    const data: t.Sale[] = await response.json();

    if (!Array.isArray(data)) {
      throw new Error("Invalid response format: expected an array of sales");
    }

    return data;
  } catch (error) {
    console.error("Failed to fetch sales:", error);
    throw error;
  }
};

export const getSaleById = async (id: number): Promise<t.Sale> => {
  try {
    const response = await fetch(`${URL}/sales/${id}`);

    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`);
    }

    const data: t.Sale = await response.json();

    if (!Array.isArray(data)) {
      throw new Error("Invalid response format: expected an array of sales");
    }

    return data;
  } catch (error) {
    console.error("Failed to fetch sale by id:", error);
    throw error;
  }
};

// --------------- Purchase Orders ---------------

// GET
export const getPurchaseOrders = async (): Promise<t.PurchaseOrder[]> => {
  try {
    const response = await fetch(`${URL}/purchase-orders`);

    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`);
    }

    const data: t.PurchaseOrder[] = await response.json();

    if (!Array.isArray(data)) {
      throw new Error(
        "Invalid response format: expected an array of purchase orders"
      );
    }

    return data;
  } catch (error) {
    console.error("Failed to fetch purchase orders:", error);
    throw error;
  }
};

// POST
export const createPurchaseOrder = async (
  purchaseOrder: t.PurchaseOrder
): Promise<t.PurchaseOrder> => {
  try {
    const formattedPurchaseOrder = {
      ...purchaseOrder,
      amount: Number(purchaseOrder.amount),
    };
    const response = await fetch(`${URL}/purchase-orders`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formattedPurchaseOrder),
    });

    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Failed to create purchase order:", error);
    throw error;
  }
};

// DELETE
export const deletePurchaseOrder = async (id: number): Promise<boolean> => {
  try {
    const response = await fetch(`${URL}/purchase-orders/${id}`, {
      method: "DELETE",
    });

    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`);
    }

    return true;
  } catch (error) {
    console.error("Failed to delete purchase order:", error);
    throw error;
  }
};
