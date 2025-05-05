## Plan

### Frontend:

- [x] Sales page

  - Components
    - [x] Header - Nav
    - [x] Graph - MUI
    - [x] Table - Sortable - MUI
    - [x] Button

- [x] Purchase order page

  - Components
    - [x] Header - Nav
    - [x] List
    - [x] Form - CRUD
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

- [x] Sales API
- [x] Purchase order API

Database:

- [x] Sales
- [x] Purchase orders

### Notes

#### Rows.Close()

I've never seen rows.Close() before. What it apparently does:

- This line ensures that the database connection's resources are properly released once you're done using the rows.
- Holds a connection from the database connection pool.
- Consumes resources on both the client and the server.
- Should be called right after checking for err.

#### Drawbacks

- No hooks
- Bad error handling
- No tests

### Criterias

- Sales page
- [x] A graph of sales per month
- [x] A filter to choose which product(s) you are looking at

- Purchase Order page
- [x] A simple list of all existing purchase orders
- [x] Ability to add new purchase orders or delete existing

#### Tech stack

- [x] Golang backend with Chi
- [x] Postgres database

### Running

I'm on node v18.12.1

docker-compose up --build -d
docker-compose up
