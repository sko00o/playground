-- rank_a
create table rank_a (
    id UInt64,
    rank_column LowCardinality(String),
    c0 AggregateFunction(sum, Int32)
) engine = AggregatingMergeTree
order by
    id;

insert into
    rank_a
select
    *
from
    rank
where
    cityHash64(rank_column) % 2 = 1;

optimize table rank_a final;

-- rank_b
create table rank_b (
    id UInt64,
    rank_column LowCardinality(String),
    c0 AggregateFunction(sum, Int32)
) engine = AggregatingMergeTree
order by
    id;

insert into
    rank_b
select
    *
from
    rank
where
    cityHash64(rank_column) % 2 = 0;

optimize table rank_b final;