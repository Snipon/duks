deps:
	go get -d -v ./...
	go install -v ./...

build:
	@make decrypt
	@docker build -t snipon/duks_api .

run:
	@echo 'Starting API...'
	@docker run -dit --ssh --name duks-api -p 3000:3000 snipon/duks_api

clean:
	docker rf -f duks-api

decrypt:
	@gpg -d token.gpg > token.json
	@gpg -d credentials.gpg > credentials.json
