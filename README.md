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
    - [x] Header - Nav
    - [x] Graph - MUI
    - [x] Table - Sortable - MUI
    - [x] Button

- [ ] Purchase order page

  - Components
    - [x] Header - Nav
    - [x] List
    - [ ] Form - CRUD
    - [x] Button

- [x] 404 page - This won't actually be needed, I'll just redirect to the sales page

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
  itemId int
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
    itemId int
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

### Notes

#### Rows.Close()

I've never seen rows.Close() before. What it apparently does:

- This line ensures that the database connection's resources are properly released once you're done using the rows.
- Holds a connection from the database connection pool.
- Consumes resources on both the client and the server.
- Should be called right after checking for err.

#### Sorting the sales by month

I'm putting the hard work on the SQL query, but not sure if Go should handle the sorting.
