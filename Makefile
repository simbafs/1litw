.PHONY: ent
ent:
	go generate ./ent
	go mod tidy


.PHONY: dev
dev: 
	nodemon --ext go --signal SIGTERM --exec "go run ."
