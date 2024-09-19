# Include the main .env file
LANG=en_US.UTF-8
SHELL=/bin/bash
.SHELLFLAGS=--norc --noprofile -e -u -o pipefail -c

include config/state.env
# Construct the variable name based on STATE
CURRENT_STATE_FILE = config/$(STATE).env
# Include the appropriate .env file (e.g., dev.env or prod.env)
include $(CURRENT_STATE_FILE)

# Include the additional .env file
include config/shared.env


migu: 
	migrate -path ../mln_rms_core/common/db/migration -database ${DB_SOURCE} -verbose up
cdb: 
	docker exec -it postgres  createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}
ddb:
	docker exec -it postgres  dropdb --username=${DB_USER}   ${DB_NAME}  --force
rdb:
	make  ddb cdb migu
run:
	 go run main.go ./config ./data/
sqlc :
	rm -rf common/db/gen/*.sql.go && sqlc generate


test:
	go test -v -cover -race -short ./...

	