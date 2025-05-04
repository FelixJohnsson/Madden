import * as t from "../types";

const PORT = 3000;
const URL = `http://localhost:${PORT}/api`;

/* 
| Page            | Read | Create | Update | Delete |
| --------------- | ---- | ------ | ------ | ------ |
| Sales           | ✅    | ❌      | ❌      | ❌      |
| Purchase Orders | ✅    | ✅      | ❌      | ✅      |
*/

// --------------- Sales ---------------

export const getSales = async (
  page: number = 1,
  pageSize: number = 100
): Promise<t.Sale[]> => {
  try {
    const response = await fetch(
      `${URL}/sales?page=${page}&pageSize=${pageSize}`
    );

    if (!response.ok) {
      throw new Error(`API error: ${response.status} ${response.statusText}`);
    }

    const data = await response.json();

    if (!Array.isArray(data)) {
      throw new Error("Invalid response format: expected an array of sales");
    }

    console.log(data);
    return data;
  } catch (error) {
    console.error("Failed to fetch sales:", error);
    throw error;
  }
};

export const getSaleById = async (id: number): Promise<t.Sale> => {
  const response = await fetch(`${URL}/sales/${id}`);
  return response.json();
};

// --------------- Purchase Orders ---------------

// GET
export const getPurchaseOrders = async (): Promise<t.PurchaseOrder[]> => {
  const response = await fetch(`${URL}/purchase-orders`);
  return response.json();
};

// POST
export const createPurchaseOrder = async (
  purchaseOrder: t.PurchaseOrder
): Promise<t.PurchaseOrder> => {
  const response = await fetch(`${URL}/purchase-orders`, {
    method: "POST",
    body: JSON.stringify(purchaseOrder),
  });
  return response.json();
};

// DELETE
export const deletePurchaseOrder = async (id: number): Promise<void> => {
  await fetch(`${URL}/purchase-orders/${id}`, {
    method: "DELETE",
  });
};

/* NOT NEEDED, BUT I'LL KEEP IT FOR NOW */
// --------------- Companies ---------------

export const getCompanies = async (): Promise<t.Company[]> => {
  const response = await fetch(`${URL}/companies`);
  return response.json();
};

export const getCompanyById = async (id: number): Promise<t.Company> => {
  const response = await fetch(`${URL}/companies/${id}`);
  return response.json();
};

// --------------- Items ---------------

export const getItems = async (): Promise<t.Item[]> => {
  const response = await fetch(`${URL}/items`);
  return response.json();
};

export const getItemById = async (id: number): Promise<t.Item> => {
  const response = await fetch(`${URL}/items/${id}`);
  return response.json();
};
