#!/usr/bin/env sh

# sql reference: https://gist.github.com/den-crane/9b5f871e7949fec54e86837eb0949747
# cmd reference: https://github.com/golang-migrate/migrate/tree/master/database/clickhouse
migrate -path "${DIR}" -database "${DSN}" up
