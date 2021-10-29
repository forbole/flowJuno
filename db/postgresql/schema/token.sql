CREATE TABLE supply(
  one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
  height BIGINT  NOT NULL ,
  supply BIGINT  NOT NULL
);

CREATE INDEX supply_height_index ON supply (height);
