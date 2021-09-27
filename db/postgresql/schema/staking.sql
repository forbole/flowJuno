CREATE TABLE total_stake_by_type
  (  height BIGINT  NOT NULL ,
	role TEXT NOT NULL ,
	total_stake TEXT NOT NULL
  );

CREATE TABLE stake_requirements
(  height BIGINT  NOT NULL ,
  role TEXT NOT NULL ,
  requirements TEXT NOT NULL 
);

CREATE TABLE weekly_payout
(  height BIGINT  NOT NULL ,
  payout TEXT NOT NULL
);

CREATE TABLE total_stake
(  height BIGINT  NOT NULL ,
  total_stake TEXT NOT NULL
);

CREATE TABLE staking_table
(  height BIGINT  NOT NULL ,
  staking_table TEXT NOT NULL
);

CREATE TABLE proposed_table
(  height BIGINT  NOT NULL ,
  proposed_table TEXT NOT NULL
);
