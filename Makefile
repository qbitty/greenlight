migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up

psql $GREENLIGHT_DB_DSN