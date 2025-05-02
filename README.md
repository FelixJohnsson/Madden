# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules. To run the application:

1. Install dependencies

```bash
npm install
```

1. Start the dev-server

```bash
npm run start
```

## Plan

### Frontend:

- [ ] Sales page

  - Components
    - [ ] Header - Nav
    - [ ] Graph
    - [ ] Table - Sortable
    - [ ] Button

- [ ] Purchase order page

  - Components
    - [ ] Header - Nav
    - [ ] List
    - [ ] Form - CRUD
    - [ ] Button

- [ ] 404 page

  - Components
    - [ ] Header - Nav

https://mui.com/

Hooks

- [ ] useSales
- [ ] usePurchaseOrders

API - CRUD

- Get all sales - Needs to be paginated
- Get sales by id

- Get all purchase orders
- Get purchase order by id
- Create purchase order
- Delete purchase order

### Backend:

- [ ] Sales API
- [ ] Purchase order API

Database:

- [ ] Sales
- [ ] Purchase orders

```
struct sale {
id int
amount float64
currency string
date time.Time
}

struct item {
    id int
    name string
    price float64
    currency string
    quantity int
    companyId int
}

struct purchaseOrder {
    id int
    item Item
    amount float64
    currency string
    createdAt time.Time
    status string
    companyId int
}

struct company {
    id int
    name string
    address string
}
```
