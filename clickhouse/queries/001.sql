SELECT
    rank_column,
    sumMerge(c0) AS c
FROM
    rank
GROUP BY
    rank_column
ORDER BY
    c DESC
LIMIT
    20 FORMAT PrettyCompactNoEscapes;