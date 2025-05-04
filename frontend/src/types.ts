// --------------- Company ---------------

export type Company = {
  id: number;
  name: string;
  address: string;
};

// --------------- Item ---------------

export type Item = {
  id: number;
  name: string;
  price: number;
  currency: string;
  quantity: number;
  companyId: number;
};

// --------------- Purchase Order ---------------

export type PurchaseOrder = {
  id: number;
  item: Item;
  amount: number;
  currency: string;
  createdAt: string;
  status: string;
  companyId: number;
};

// --------------- Sale ---------------

export type Sale = {
  id: number;
  amount: number;
  currency: string;
  date: string;
};

export type SaleGroupedByMonth = {
  month: string;
  sales: Sale[];
};
