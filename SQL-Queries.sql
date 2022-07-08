-- Create table queries
CREATE TABLE IF NOT EXISTS transactions(
  id VARCHAR(255) PRIMARY KEY,
  order_date TIMESTAMP NOT NULL,
  status VARCHAR(255) NOT NULL,
  paid_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transaction_details(
  id VARCHAR(255) PRIMARY KEY,
  transaction_id VARCHAR(255) NOT NULL,
  price NUMERIC NOT NULL,
  amount NUMERIC NOT NULL,
  subtotal NUMERIC NOT NULL
);

-- Seeding Query
INSERT INTO transactions (id, order_date, status, paid_at)
VALUES
  ('trx-0001', NOW(), 'PAID', NOW()),
  ('trx-0002', NOW(), 'PENDING', NULL);

INSERT INTO transaction_details (id, transaction_id, price, amount, subtotal)
VALUES
  ('trx-d-0001', 'trx-0001', 10000, 1, 10000),
  ('trx-d-0002', 'trx-0001', 20000, 2, 40000),
  ('trx-d-0003', 'trx-0002', 30000, 1, 30000),
  ('trx-d-0004', 'trx-0002', 40000, 2, 80000);

-- Select Query
SELECT * FROM transactions;
SELECT * FROM transaction_details;

SELECT trx.id, trx.order_date, trx.status, trx.paid_at, SUM(trxd.subtotal) grand_total, SUM(trxd.amount) item_amount
FROM transactions as trx LEFT JOIN transaction_details as trxd ON trxd.transaction_id = trx.id
GROUP BY trx.id;
