CREATE TABLE account
(
    address TEXT UNIQUE PRIMARY KEY NOT NULL
);

CREATE TABLE account_balance(
    address TEXT UNIQUE PRIMARY KEY NOT NULL REFERENCES account(address),
    balance BIGINT NOT NULL,
    code TEXT NOT NULL,
    contract_map JSONB,
    height BIGINT NOT NULL
);

CREATE TABLE locked_account
(
    address TEXT  NOT NULL NOT NULL UNIQUE REFERENCES account(address),
    locked_address TEXT  NOT NULL UNIQUE
);

CREATE TABLE locked_account_balance(
    locked_address TEXT NOT NULL REFERENCES locked_account(locked_address),
    balance BIGINT NOT NULL,
    unlock_limit BIGINT NOT NULL,
    height BIGINT NOT NULL,
    PRIMARY KEY (locked_address,height)
);

CREATE INDEX locked_account_balance_index ON locked_account_balance (height);


CREATE TABLE delegator_account(
    account_address TEXT NOT NULL REFERENCES account(address),
	delegator_id    BIGINT NOT NULL ,
	delegator_node_id   TEXT NOT NULL,
    PRIMARY KEY (delegator_id,delegator_node_id)

);

CREATE TABLE account_key_list( 
  address TEXT  NOT NULL REFERENCES account(address),
  index BIGINT NOT NULL UNIQUE,
  weight TEXT  NOT NULL ,
  revoked BOOLEAN  NOT NULL ,
  sig_algo TEXT  NOT NULL ,
  hash_algo TEXT  NOT NULL ,
  public_key TEXT  NOT NULL ,
  sequence_number BIGINT  NOT NULL,
  PRIMARY KEY (address,index)
);

CREATE TABLE staker_node_id(
    address TEXT  NOT NULL REFERENCES account(address),
    node_id TEXT NOT NULL UNIQUE REFERENCES staking_table (node_id)
);