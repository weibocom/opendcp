#!/bin/sh

# make sure all database & tables are created
echo "Host is $DB_HOST"
SQLS=`ls sql/`
for SQL in $SQLS; do
    echo "Loading $SQL ..."
    mysql -h $DB_HOST -u root --default-character-set=utf8 -p${MYSQL_ROOT_PASSWORD} < sql/$SQL
done

# start listening on INIT_PORT, so other contaner could wait for it
nc -l $INIT_PORT > file
