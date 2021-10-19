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

CREATE TABLE current_table
(  height BIGINT  NOT NULL ,
  current_table TEXT NOT NULL
);

CREATE TABLE node_total_commitment
(  node_id TEXT NOT NULL ,
  total_commitment TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_total_commitment_without_delegators
(  node_id TEXT NOT NULL ,
  total_commitment_without_delegators TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_infos_from_table
(  id TEXT  NOT NULL ,
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

CREATE TABLE cut_percentage
(  cut_percentage BIGINT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE delegator_info
(  id TEXT NOT NULL ,
  node_id TEXT NOT NULL ,
  tokens_committed TEXT NOT NULL ,
  tokens_staked TEXT NOT NULL ,
  tokens_unstaking TEXT NOT NULL ,
  tokens_rewarded TEXT NOT NULL ,
  tokens_unstaked TEXT NOT NULL ,
  tokens_requested_to_unstake TEXT NOT NULL ,
  height TEXT NOT NULL
);

CREATE TABLE delegator_info_from_address
(  delegator_info TEXT NOT NULL ,
  height BIGINT  NOT NULL ,
  address TEXT NOT NULL
);
