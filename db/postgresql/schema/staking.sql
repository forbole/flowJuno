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

CREATE TABLE node_unstaking_tokens
(  node_id TEXT NOT NULL ,
  token_unstaking TEXT NOT NULL ,
  height BIGINT  NOT NULL
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

CREATE TABLE node_staking_key
(  node_id TEXT NOT NULL ,
  node_staking_key TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_staked_tokens
(  node_id TEXT NOT NULL ,
  node_staked_tokens TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_role
(  node_id TEXT NOT NULL ,
  role TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_rewarded_tokens
(  node_id TEXT NOT NULL ,
  node_rewarded_tokens TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_networking_key
(  node_id TEXT NOT NULL ,
  networking_key TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_networking_address
(  node_id TEXT NOT NULL ,
  networking_address TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_initial_weight
(  node_id TEXT NOT NULL ,
  initial_weight TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_info_from_address
(  address TEXT NOT NULL ,
  node_info TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_info_from_node_id
(  node_id TEXT NOT NULL ,
  node_info TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE node_committed_tokens
(  node_id TEXT NOT NULL ,
  committed_tokens TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE cut_percentage
(  cut_percentage TEXT NOT NULL ,
  height BIGINT  NOT NULL
);

CREATE TABLE delegator_committed
(  committed TEXT NOT NULL ,
  height TEXT NOT NULL ,
  node_id TEXT NOT NULL ,
  delegator_id TEXT NOT NULL
);

CREATE TABLE delegator_info
(  delegator_info TEXT NOT NULL ,
  height TEXT NOT NULL ,
  node_id TEXT NOT NULL ,
  delegator_id TEXT NOT NULL
);

CREATE TABLE delegator_info_from_address
(  delegator_info TEXT NOT NULL ,
  height BIGINT  NOT NULL ,
  address TEXT NOT NULL
);

CREATE TABLE delegator_request
(  request_to_unstake TEXT NOT NULL ,
  height BIGINT  NOT NULL ,
  node_id TEXT NOT NULL ,
  delegator_id TEXT NOT NULL
);

CREATE TABLE delegator_rewarded
(  rewarded TEXT NOT NULL ,
  height BIGINT  NOT NULL ,
  node_id TEXT NOT NULL ,
  delegator_id TEXT NOT NULL
);