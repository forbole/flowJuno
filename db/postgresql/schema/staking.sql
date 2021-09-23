CREATE TABLE total_stake
  (  height BIGINT  NOT NULL ,
	role TEXT NOT NULL ,
	total_stake TEXT NOT NULL ,
	timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL ,
  );

CREATE TABLE stake_requirements
(  height BIGINT  NOT NULL ,
  role TEXT NOT NULL ,
  requirements TEXT NOT NULL ,
  timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL ,
);

