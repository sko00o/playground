-- try SimpleAggregateFunction
create table rank_simple (
    id UInt64,
    rank_column LowCardinality(String),
    c0 SimpleAggregateFunction(sum, Int64)
) engine = AggregatingMergeTree
order by
    id;

-- intermediate table
create table nop_simple (
    id UInt64,
    rank_column LowCardinality(String),
    c Int32
) Engine = Null;

-- materialized view, just like an insert trigger
create materialized view nop_simple_mv to rank_simple as
select
    id,
    rank_column,
    sum(c) as c0
from
    nop_simple
group by
    id,
    rank_column;

-- do some insert
insert into
    nop_simple
select
    number id,
    toString(id % 15000),
    toInt32(rand() % 12345)
from
    numbers(1000000);