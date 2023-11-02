CREATE TABLE IF NOT EXISTS subscriptions(
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL UNIQUE,
    description varchar(255) NOT NULL,
    admin_id int NOT NULL,
    max_members int NOT NULL CHECK(max_members > 0),
    price int NOT NULL CHECK(price > -1),
    currency varchar(3) NOT NULL,
    commission int NOT NULL,
    charge_period int NOT NULL,
    creation timestamp NOT NULL,
    start timestamp NOT NULL,
    ending timestamp
);

CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    Username varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    role int NOT NULL
);

CREATE TABLE IF NOT EXISTS subscribers(
     id serial PRIMARY KEY,
     userID int NOT NULL,
     subscriptionID int NOT NULL,
     subscription_start timestamp NOT NULL,
     subscription_ending timestamp,

    FOREIGN KEY (userID) REFERENCES users (id),
    FOREIGN KEY (subscriptionID) REFERENCES subscriptions (id)
);

CREATE TABLE IF NOT EXISTS invoices(
    id serial PRIMARY KEY,
    userID int NOT NULL,
    subscriptionID int NOT NULL,
    amount int NOT NULL,
    issued timestamp NOT NULL,
    payed timestamp,

    FOREIGN KEY (userID) REFERENCES users (id),
    FOREIGN KEY (subscriptionID) REFERENCES subscriptions (id)
);

CREATE TABLE IF NOT EXISTS payments(
     invoiceID serial PRIMARY KEY,
     amount int NOT NULL,
     accepted boolean,

     FOREIGN KEY (invoiceID) REFERENCES invoices (id)
);

CREATE TABLE IF NOT EXISTS auth(
    Username varchar(255) PRIMARY KEY,
    refresh_token varchar(10) NOT NULL,
    expires timestamp NOT NULL
);

DROP TABLE subscriptions CASCADE;
DROP TABLE users CASCADE;
DROP TABLE subscribers CASCADE;
DROP TABLE invoices CASCADE;
DROP TABLE payments CASCADE;
DROP TABLE auth CASCADE;
