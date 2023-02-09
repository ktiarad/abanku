CREATE DATABASE abanku;

\c abanku;

CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    balance NUMERIC(14, 2)
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES accounts (id),
    transaction_type VARCHAR(20),
    description VARCHAR(50),
    amount NUMERIC(11, 2),
    ending_balance NUMERIC (14,2),
    transaction_date TIMESTAMP DEFAULT NOW()
);