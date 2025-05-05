CREATE TABLE IF NOT EXISTS companies (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  address TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  price NUMERIC(10,2) NOT NULL,
  currency TEXT NOT NULL,
  quantity INT NOT NULL,
  company_id INT NOT NULL REFERENCES companies(id)
);

CREATE TABLE IF NOT EXISTS sales (
  id SERIAL PRIMARY KEY,
  amount   NUMERIC(10,2) NOT NULL,
  item_id INT NOT NULL REFERENCES items(id),
  currency TEXT NOT NULL,
  date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS purchase_orders (
  id SERIAL PRIMARY KEY,
  item_id INT NOT NULL REFERENCES items(id),
  amount   NUMERIC(10,2) NOT NULL,
  currency TEXT NOT NULL,
  created_at DATE NOT NULL,
  status TEXT NOT NULL,
  company_id INT NOT NULL REFERENCES companies(id)
);

