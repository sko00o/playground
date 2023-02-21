SELECT
    rank_column,
    sum(c0) AS c
FROM
    rank_simple
GROUP BY
    rank_column
ORDER BY
    c DESC
LIMIT
    20;