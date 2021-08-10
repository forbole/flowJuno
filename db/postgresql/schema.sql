CREATE TABLE node_info
(
    id TEXT UNIQUE NOT NULL PRIMARY KEY,
    role BIGINT NOT NULL,
    networkingAddress TEXT NOT NULL,
    networkingKey TEXT NOT NULL,
    stakingKey TEXT NOT NULL,
    tokensStaked NUMBER NOT NULL,
    tokensCommitted NUMBER NOT NULL,
    tokensUnstaking NUMBER NOT NULL,
    tokensUnstaked NUMBER NOT NULL,
    tokensRewarded NUMBER NOT NULL,

    delegators NUMBER[] ,
    delegatorIDCounter BIGINT NOT NULL,
    tokensRequestedToUnstake NUMBER NOT NULL,
    initialWeight BIGINT NOT NULL
);

CREATE TABLE block
(
    height           BIGINT UNIQUE PRIMARY KEY,
    id TEXT NOT NULL,
    parent_id TEXT NOT NULL,
    collection_guarantees []TEXT NOT NULL,
    timestamp        TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE block_seal
(
    height BIGINT NOT NULL,
    execution_receipt_id TEXT NOT NULL,
    execution_receipt_signatures [][]TEXT NOT NULL
)
CREATE INDEX block_proposer_address_index ON block (proposer_address);

CREATE TABLE transaction
(
   		status TEXT NOT NULL,
		height BIGINT NOT NULL,
        id TEXT,

		script TEXT ,
		Arguments TEXT,
		ReferenceBlockID TEXT,
		GasLimit BIGINT,
		ProposalKey TEXT,
		Payer TEXT,
		Authorizers TEXT,
		PayloadSignature TEXT,
		EnvelopeSignatures TEXT
);
CREATE INDEX transaction_hash_index ON transaction (hash);
CREATE INDEX transaction_height_index ON transaction (height);

CREATE TABLE event
(
    height BIGINT NOT NULL,
    type TEXT,
    transaction_id TEXT,
    transaction_index TEXT,
    event_index BIGINT,
    value TEXT,
)

CREATE TABLE pruning
(
    last_pruned_height BIGINT NOT NULL
)