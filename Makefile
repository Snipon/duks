deps:
	go get -d -v ./...
	go install -v ./...

build:
	@echo 'Building docker container...'
	@docker build --build-arg gpg_secret=$GPG_SECRET . -t snipon/duks_api

run:
	@echo 'Starting API...'
	@docker run -dit --ssh --name duks-api -p 3000:3000 snipon/duks_api

clean:
	docker rf -f duks-api

decrypt:
	@gpg --batch --yes --passphrase ${GPG_SECRET} -o token.json -d token.gpg
	@gpg --batch --yes --passphrase ${GPG_SECRET} -o credentials.json -d credentials.gpg
