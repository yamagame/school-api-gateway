#!/bin/bash
source ./scripts/.env.local
echo "DROP DATABASE IF EXISTS \`$DB_NAME\`;" | mysql -u root -h $DB_HOST -P $DB_PORT -D $DB_NAME
echo "CREATE DATABASE IF NOT EXISTS \`$DB_NAME\`;" | mysql -u root -h $DB_HOST -P $DB_PORT
./scripts/migrate.sh
mysql -u root -h $DB_HOST -P $DB_PORT -D $DB_NAME < ./datasets/sample-school-labos.sql
