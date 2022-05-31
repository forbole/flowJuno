CREATE TABLE block
(
    height           BIGINT UNIQUE PRIMARY KEY,
    id               TEXT NOT NULL UNIQUE,
    parent_id        TEXT NOT NULL,
    collection_guarantees JSONB NOT NULL,
    timestamp        TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX block_index ON block (height);
CREATE INDEX block_id_index ON block (id);


CREATE TABLE block_seal
(
    height BIGINT NOT NULL REFERENCES block (height),
    execution_receipt_id TEXT UNIQUE,
    execution_receipt_signatures TEXT[][],
    result_approval_signatures TEXT[][]
);

CREATE INDEX block_seal_index ON block_seal (height);
CREATE INDEX block_seal_execution_receipt_id_index ON block_seal (execution_receipt_id);


CREATE TABLE collection
(  height BIGINT  NOT NULL REFERENCES block (height),
  id TEXT  NOT NULL,
  processed BOOLEAN  NOT NULL ,
  transaction_id TEXT  NOT NULL DEFAULT ' '
);

CREATE INDEX collection_index ON collection (height);
CREATE INDEX collection_transaction_id_index ON collection (transaction_id);


CREATE TABLE transaction
(
		height BIGINT NOT NULL REFERENCES block (height),
        transaction_id TEXT NOT NULL,

		script TEXT ,
		arguments TEXT[],
		reference_block_id TEXT,
		gas_limit BIGINT,
		proposal_key TEXT,
		payer TEXT,
		authorizers TEXT[],
		payload_signature JSONB,
		envelope_signatures JSONB
) PARTITION BY RANGE(height);
CREATE INDEX transaction_index ON transaction (height);
CREATE INDEX transaction_id_index ON transaction (transaction_id);

CREATE TABLE transaction_default PARTITION OF transaction DEFAULT; 



CREATE TABLE transaction_result
(  height BIGINT  NOT NULL REFERENCES block (height),
  transaction_id TEXT  NOT NULL,
  status TEXT  NOT NULL ,
  error TEXT
) PARTITION BY RANGE(height);

CREATE INDEX transaction_result_index ON transaction_result (height);
CREATE TABLE transaction_result_default PARTITION OF transaction_result DEFAULT;

CREATE TABLE event
(
    height BIGINT NOT NULL REFERENCES block (height),
    type TEXT,
    transaction_id TEXT  ,
    transaction_index TEXT,
    event_index BIGINT,
    value TEXT
)  PARTITION BY RANGE(height);

CREATE INDEX event_index ON event (height);
CREATE INDEX event_transaction_index ON event (transaction_id);

CREATE TABLE event_default PARTITION OF event DEFAULT;
