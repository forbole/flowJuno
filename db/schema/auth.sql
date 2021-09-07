CREATE TABLE account(
    address TEXT UNIQUE PRIMARY KEY NOT NULL
);

CREATE TABLE locked_account(
    account_address TEXT UNIQUE PRIMARY KEY NOT NULL REFERENCES address(account),
    locked_address TEXT NOT NULL
);