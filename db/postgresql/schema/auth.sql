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
    address TEXT  NOT NULL NOT NULL REFERENCES account(address),
    locked_address TEXT  NOT NULL UNIQUE
);

CREATE TABLE locked_account_delegator
(  
  locked_address TEXT  NOT NULL REFERENCES locked_account(locked_address),
  node_id TEXT  NOT NULL  ,
  delegator_id BIGINT  NOT NULL 
);

CREATE TABLE locked_account_staker
(
    locked_address TEXT  NOT NULL NOT NULL REFERENCES locked_account(locked_address),
    node_id TEXT  NOT NULL 
);

CREATE TABLE locked_account_balance(
    locked_address TEXT NOT NULL REFERENCES locked_account(locked_address),
    balance BIGINT NOT NULL,
    unlock_limit BIGINT NOT NULL,
    height BIGINT NOT NULL
);

CREATE INDEX locked_account_balance_index ON locked_account_balance (height);


CREATE TABLE delegator_account(
    account_address TEXT NOT NULL REFERENCES account(address),
	delegator_id    BIGINT NOT NULL,
	delegator_node_id   TEXT NOT NULL
);

CREATE TABLE account_key_list( 
  address TEXT  NOT NULL REFERENCES account(address),
  index TEXT NOT NULL ,
  weight TEXT  NOT NULL ,
  revoked BOOLEAN  NOT NULL ,
  sig_algo TEXT  NOT NULL ,
  hash_algo TEXT  NOT NULL ,
  public_key TEXT  NOT NULL ,
  sequence_number TEXT  NOT NULL
);
