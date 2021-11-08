CREATE TABLE block
(
    height           BIGINT UNIQUE PRIMARY KEY,
    id               TEXT NOT NULL,
    parent_id        TEXT NOT NULL,
    collection_guarantees JSONB NOT NULL,
    timestamp        TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE block_seal
(
    height BIGINT NOT NULL REFERENCES block (height),
    execution_receipt_id TEXT ,
    execution_receipt_signatures TEXT[][]
);

CREATE TABLE collection
(  height BIGINT  NOT NULL REFERENCES block (height),
  id TEXT  NOT NULL ,
  processed BOOLEAN  NOT NULL ,
  transaction_id TEXT  NOT NULL UNIQUE
);

CREATE TABLE transaction
(
		height BIGINT NOT NULL REFERENCES block (height),
        transaction_id TEXT NOT NULL REFERENCES collection (transaction_id),

		script TEXT ,
		arguments TEXT[],
		reference_block_id TEXT,
		gas_limit BIGINT,
		proposal_key TEXT,
		payer TEXT,
		authorizers TEXT[],
		payload_signature JSONB,
		envelope_signatures JSONB
);

CREATE TABLE transaction_result
(  height BIGINT  NOT NULL REFERENCES block (height),
  transaction_id TEXT  NOT NULL REFERENCES collection (transaction_id),
  status TEXT  NOT NULL ,
  error TEXT 
);


CREATE TABLE event
(
    height BIGINT NOT NULL REFERENCES block (height),
    type TEXT,
    transaction_id TEXT REFERENCES collection (transaction_id),
    transaction_index TEXT,
    event_index BIGINT,
    value TEXT
);

CREATE TABLE pruning
(
    last_pruned_height BIGINT NOT NULL
);

