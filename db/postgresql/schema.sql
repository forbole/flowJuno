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
CREATE INDEX block_hash_index ON block (hash);
CREATE INDEX block_proposer_address_index ON block (proposer_address);

CREATE TABLE pre_commit
(
    validator_address TEXT                        NOT NULL REFERENCES validator (consensus_address),
    height            BIGINT                      NOT NULL,
    timestamp         TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    voting_power      BIGINT                      NOT NULL,
    proposer_priority INTEGER                     NOT NULL,
    UNIQUE (validator_address, timestamp)
);
CREATE INDEX pre_commit_validator_address_index ON pre_commit (validator_address);
CREATE INDEX pre_commit_height_index ON pre_commit (height);

CREATE TABLE transaction
(
    hash         TEXT    NOT NULL UNIQUE PRIMARY KEY,
    height       BIGINT  NOT NULL REFERENCES block (height),
    success      BOOLEAN NOT NULL,

    /* Body */
    messages     JSONB   NOT NULL DEFAULT '[]'JSONB,
    memo         TEXT,
    signatures   TEXT[]  NOT NULL,

    /* AuthInfo */
    signer_infos JSONB   NOT NULL DEFAULT '[]'JSONB,
    fee          JSONB   NOT NULL DEFAULT '{}'JSONB,

    /* Tx response */
    gas_wanted   BIGINT           DEFAULT 0,
    gas_used     BIGINT           DEFAULT 0,
    raw_log      TEXT,
    logs         JSONB
);
CREATE INDEX transaction_hash_index ON transaction (hash);
CREATE INDEX transaction_height_index ON transaction (height);

CREATE TABLE message
(
    transaction_hash            TEXT   NOT NULL REFERENCES transaction (hash),
    index                       BIGINT NOT NULL,
    type                        TEXT   NOT NULL,
    value                       JSONB  NOT NULL,
    involved_accounts_addresses TEXT[] NULL
);
CREATE INDEX message_transaction_hash_index ON message (transaction_hash);

CREATE TABLE pruning
(
    last_pruned_height BIGINT NOT NULL
)