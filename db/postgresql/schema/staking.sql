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
