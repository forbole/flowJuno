CREATE TABLE node_info
(
    id TEXT UNIQUE NOT NULL PRIMARY KEY,
    role BIGINT NOT NULL,
    networkingAddress TEXT NOT NULL,
    networkingKey TEXT NOT NULL,
    stakingKey TEXT NOT NULL,
    tokensStaked NUMERIC NOT NULL,
    tokensCommitted NUMERIC NOT NULL,
    tokensUnstaking NUMERIC NOT NULL,
    tokensUnstaked NUMERIC NOT NULL,
    tokensRewarded NUMERIC NOT NULL,

    delegators NUMERIC[] ,
    delegatorIDCounter BIGINT NOT NULL,
    tokensRequestedToUnstake NUMERIC NOT NULL,
    initialWeight BIGINT NOT NULL
);

CREATE TABLE block
(
    height           BIGINT UNIQUE PRIMARY KEY,
    id               TEXT NOT NULL,
    parent_id        TEXT NOT NULL,
    collection_guarantees TEXT NOT NULL,
    timestamp        TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE block_seal
(
    height BIGINT NOT NULL,
    execution_receipt_id TEXT ,
    execution_receipt_signatures TEXT[][]
);

CREATE TABLE transaction
(
   		status TEXT NOT NULL,
		height BIGINT NOT NULL,
        transaction_id TEXT,

		script TEXT ,
		arguments TEXT,
		reference_block_id TEXT,
		gas_limit BIGINT,
		proposal_key TEXT,
		payer TEXT,
		authorizers TEXT,
		payload_signature TEXT,
		envelope_signatures TEXT
);

CREATE TABLE event
(
    height BIGINT NOT NULL,
    type TEXT,
    transaction_id TEXT,
    transaction_index TEXT,
    event_index BIGINT,
    value TEXT
);

CREATE TABLE pruning
(
    last_pruned_height BIGINT NOT NULL
);