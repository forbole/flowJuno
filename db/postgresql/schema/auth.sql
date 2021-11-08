CREATE TABLE account(
    address TEXT UNIQUE PRIMARY KEY NOT NULL,
    balance BIGINT NOT NULL,
    code TEXT NOT NULL,
    keys_list JSONB,
    contract_map JSONB
);

CREATE TABLE locked_account(
    account_address TEXT UNIQUE PRIMARY KEY NOT NULL REFERENCES account(address),
    locked_address TEXT NOT NULL,
    balance BIGINT NOT NULL,
    unlock_limit BIGINT NOT NULL
);

CREATE TABLE delegator_account(
    account_address TEXT UNIQUE PRIMARY KEY NOT NULL REFERENCES account(address),
	delegator_id    BIGINT NOT NULL,
	delegator_node_id   TEXT NOT NULL,
	delegator_node_info JSONB NOT NULL
);

CREATE TABLE staker_account(
    account_address TEXT UNIQUE PRIMARY KEY NOT NULL REFERENCES account(address),
	staker_node_id    TEXT NOT NULL,
	staker_node_info   JSONB NOT NULL
)