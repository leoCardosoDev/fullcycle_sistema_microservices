CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date);
CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date);
CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date);

UPDATE accounts SET balance = 1000 where id = "7372c79a-e1c5-467f-9178-b9b4072bf901";
