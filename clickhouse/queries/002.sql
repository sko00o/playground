SELECT
    rank_column,
    c
FROM
    (
        SELECT
            rank_column,
            sumMerge(c0) AS c
        FROM
            rank_a
        GROUP BY
            rank_column
        ORDER BY
            c DESC
        LIMIT
            20
        UNION
        ALL
        SELECT
            rank_column,
            sumMerge(c0) AS c
        FROM
            rank_b
        GROUP BY
            rank_column
        ORDER BY
            c DESC
        LIMIT
            20
    )
ORDER BY
    c DESC
LIMIT
    20;