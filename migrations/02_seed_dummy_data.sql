BEGIN;

INSERT INTO companies (name, address) VALUES
  ('Nudie Jeans.', '123 Tech Avenue, Silicon Valley, CA 94043'),
  ('H&M', '456 Factory Road, Detroit, MI 48127'),
  ('SNS', '789 Harvest Street, Portland, OR 97201'),
  ('Zara', '101 Design Boulevard, Chicago, IL 60611'),
  ('Bubbleroom', '202 Stadium Way, Boston, MA 02210');

INSERT INTO items (name, price, currency, quantity, company_id) VALUES
  ('Designer Denim Jacket', 1299.99, 'SEK', 50, 1),
  ('Cotton T-Shirt', 249.99, 'SEK', 120, 1),
  ('Slim Fit Chinos', 489.95, 'SEK', 75, 1),
  ('Wool Winter Coat', 1829.50, 'SEK', 30, 2),
  ('Knitted Sweater', 645.75, 'SEK', 50, 2),
  ('Fleece Hoodie', 599.99, 'SEK', 60, 2),
  ('Organic Cotton Dress', 898.99, 'SEK', 40, 3),
  ('Linen Summer Shorts', 454.50, 'SEK', 65, 3),
  ('Hemp Casual Shirt', 576.75, 'SEK', 45, 3),
  ('Formal Business Suit', 3899.00, 'SEK', 15, 4),
  ('Silk Necktie', 349.95, 'SEK', 25, 4),
  ('Leather Belt', 425.50, 'SEK', 40, 4),
  ('Sport Performance Leggings', 589.99, 'SEK', 50, 5),
  ('Quick-dry Running Shorts', 379.99, 'SEK', 60, 5),
  ('Breathable Training Top', 429.95, 'SEK', 80, 5);

INSERT INTO sales (amount, item_id, item_name, currency, date) VALUES
  (3899.97, 1, 'Designer Denim Jacket', 'SEK', '2024-04-15'),
  (999.80, 2, 'Cotton T-Shirt', 'SEK', '2024-04-20'),
  (1799.00, 10, 'Organic Cotton Dress', 'SEK', '2024-04-25'),
  (449.75, 5, 'Fleece Hoodie', 'SEK', '2024-05-01'),
  (899.95, 1, 'Designer Denim Jacket', 'SEK', '2024-05-03'),
  (179.80, 3, 'Slim Fit Chinos', 'SEK', '2024-04-10'),
  (2599.87, 13, 'Formal Business Suit', 'SEK', '2024-03-28'),
  (67.50, 8, 'Silk Necktie', 'SEK', '2024-03-15'),
  (349.95, 11, 'Leather Belt', 'SEK', '2024-02-20'),
  (199.99, 6, 'Breathable Training Top', 'SEK', '2024-02-10'),
  (1299.99, 1, 'Quick-dry Running Shorts', 'SEK', '2024-01-15'),
  (899.00, 10, 'Organic Cotton Dress', 'SEK', '2024-01-05'),
  (379.96, 14, 'Breathable Training Top', 'SEK', '2023-12-20'),
  (45.00, 8, 'Silk Necktie', 'SEK', '2023-12-10'),
  (179.90, 2, 'Cotton T-Shirt', 'SEK', '2023-11-28'),
  (89.95, 3, 'Slim Fit Chinos', 'SEK', '2023-11-15'),
  (259.90, 13, 'Formal Business Suit', 'SEK', '2023-10-25'),
  (17.98, 7, 'Hemp Casual Shirt', 'SEK', '2023-10-10'),
  (251.00, 12, 'Linen Summer Shorts', 'SEK', '2023-09-20'),
  (388.50, 4, 'Organic Cotton Dress', 'SEK', '2023-09-05');

INSERT INTO purchase_orders (item_id, item_name, amount, currency, created_at, status, company_id) VALUES
  (1, 'Designer Denim Jacket', 6499, 'SEK', '2024-04-01', 'COMPLETED', 1),
  (2, 'Cotton T-Shirt', 999, 'SEK', '2024-04-05', 'COMPLETED', 1),
  (4, 'Organic Cotton Dress', 1295, 'SEK', '2024-04-10', 'COMPLETED', 2),
  (7, 'Hemp Casual Shirt', 899, 'SEK', '2024-04-15', 'PROCESSING', 3),
  (10, 'Formal Business Suit', 2697, 'SEK', '2024-04-20', 'PROCESSING', 4),
  (13, 'Quick-dry Running Shorts', 1899, 'SEK', '2024-04-25', 'PENDING', 5),
  (6, 'Breathable Training Top', 599, 'SEK', '2024-05-01', 'PENDING', 2),
  (9, 'Silk Necktie', 337, 'SEK', '2024-05-02', 'NEW', 3),
  (11, 'Leather Belt', 1049, 'SEK', '2024-05-03', 'NEW', 4),
  (15, 'Breathable Training Top', 649, 'SEK', '2024-05-04', 'NEW', 5);

COMMIT;
