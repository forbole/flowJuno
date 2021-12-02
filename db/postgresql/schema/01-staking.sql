CREATE TABLE total_stake_by_type
  (  height BIGINT  NOT NULL ,
	role TEXT NOT NULL ,
	total_stake TEXT NOT NULL
  );

CREATE INDEX total_stake_by_type_index ON total_stake_by_type (height);


CREATE TABLE stake_requirements
(  height BIGINT  NOT NULL ,
  role TEXT NOT NULL ,
  requirements TEXT NOT NULL 
);

CREATE INDEX stake_requirements_index ON stake_requirements (height);


CREATE TABLE weekly_payout
(  height BIGINT  NOT NULL ,
  payout TEXT NOT NULL
);

CREATE INDEX weekly_payout_index ON weekly_payout (height);


CREATE TABLE total_stake
(  height BIGINT  NOT NULL ,
  total_stake BIGINT NOT NULL
);

CREATE INDEX total_stake_index ON total_stake (height);

CREATE TABLE staking_table
(  
  node_id TEXT NOT NULL UNIQUE PRIMARY KEY
);



CREATE TABLE proposed_table
(  height BIGINT  NOT NULL ,
  proposed_table TEXT NOT NULL
);

CREATE INDEX proposed_table_index ON proposed_table (height);


CREATE TABLE current_table
(  height BIGINT  NOT NULL ,
  node_id TEXT NOT NULL
);

CREATE INDEX current_table_index ON current_table (height);

/* Start from here node id is needed */
CREATE TABLE node_total_commitment
(  node_id TEXT NOT NULL REFERENCES staking_table (node_id),
  total_commitment TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE INDEX node_total_commitment_index ON node_total_commitment (height);


CREATE TABLE node_total_commitment_without_delegators
(  node_id TEXT NOT NULL REFERENCES staking_table (node_id),
  total_commitment_without_delegators TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE INDEX node_total_commitment_without_delegators_index ON node_total_commitment_without_delegators (height);


CREATE TABLE node_infos_from_table
(  id TEXT  NOT NULL REFERENCES staking_table (node_id),
  role BIGINT  NOT NULL ,
  networking_address TEXT  NOT NULL ,
  networking_key TEXT  NOT NULL ,
  staking_key TEXT  NOT NULL ,
  tokens_staked BIGINT  NOT NULL ,
  tokens_committed BIGINT  NOT NULL ,
  tokens_unstaking BIGINT  NOT NULL ,
  tokens_unstaked BIGINT  NOT NULL ,
  tokens_rewarded BIGINT  NOT NULL ,
  delegators BIGINT[]  NOT NULL ,
  delegator_i_d_counter BIGINT  NOT NULL ,
  tokens_requested_to_unstake BIGINT  NOT NULL ,
  initial_weight BIGINT  NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE INDEX node_infos_from_table_index ON node_infos_from_table (height);


CREATE TABLE cut_percentage
(  cut_percentage BIGINT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE INDEX cut_percentage_index ON cut_percentage (height);


CREATE TABLE delegator_info
(  id BIGINT NOT NULL ,
  node_id TEXT NOT NULL REFERENCES staking_table (node_id),
  tokens_committed TEXT NOT NULL ,
  tokens_staked TEXT NOT NULL ,
  tokens_unstaking TEXT NOT NULL ,
  tokens_rewarded TEXT NOT NULL ,
  tokens_unstaked TEXT NOT NULL ,
  tokens_requested_to_unstake TEXT NOT NULL ,
  height TEXT NOT NULL
);

CREATE INDEX delegator_info_index ON delegator_info (height);

