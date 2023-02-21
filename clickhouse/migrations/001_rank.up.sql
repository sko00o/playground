-- try AggregateFunction
create table rank (
    id UInt64,
    rank_column LowCardinality(String),
    c0 AggregateFunction(sum, Int32)
) engine = AggregatingMergeTree
order by
    id;

-- intermediate table
create table nop (
    id UInt64,
    rank_column LowCardinality(String),
    c Int32
) Engine = Null;

-- materialized view, just like an insert trigger
create materialized view nop_mv to rank as
select
    id,
    rank_column,
    sumState(c) as c0
from
    nop
group by
    id,
    rank_column;

-- do some insert
insert into
    nop
select
    number id,
    toString(id % 15000),
    toInt32(rand() % 12345)
from
    numbers(100000);

optimize table rank final;