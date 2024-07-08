#!/bin/bash

# Variables
DB_NAME="postgres"
DB_USER="khushi"
DB_PASSWORD="go"
TABLE_NAME="parameters"
CSV_FILE="/Users/khushi.nijhawan/Downloads/phd.csv"
DELIMITER=","
\copy parameters FROM '/Users/khushi.nijhawan/Downloads/ppp.csv' DELIMITER ',' CSV HEADER;
# Export password to avoid password prompt
export PGPASSWORD=$DB_PASSWORD
\copy $TABLE_NAME FROM '/Users/khushi.nijhawan/Downloads/phd.csv' DELIMITER ',' CSV HEADER;
# Import CSV into PostgreSQL table
psql -U $DB_USER -d $DB_NAME -c "\copy $TABLE_NAME FROM '$CSV_FILE' DELIMITER '$DELIMITER' CSV HEADER;"
psql -U $DB_USER -d $DB_NAME -c "\copy $TABLE_NAME FROM '$CSV_FILE' DELIMITER '$DELIMITER' CSV HEADER;"


# Unset the password
unset PGPASSWORD

echo "CSV file $CSV_FILE has been imported into $TABLE_NAME in $DB_NAME database."
