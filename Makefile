migrate-create:  ### create new migration
	migrate create -ext sql -dir db/migrations '$(title)'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path db/migrations -database '$(DB_URL)' -verbose up
.PHONY: migrate-up

migrate-down: ### migration up
	migrate -path db/migrations -database '$(DB_URL)' -verbose down
.PHONY: migrate-up