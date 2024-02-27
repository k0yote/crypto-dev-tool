build:
	@go build -o ./bin/backend-wallet

run:build
	@./bin/backend-wallet

start-hardhat:
	docker build . -t local-hardhat
	docker start hardhat-node || docker run --name hardhat-node -d -p 8545:8545 local-hardhat
	sh ./scripts/test/await-hardhat.sh

stop-hardhat:
	docker stop hardhat-node
	docker rm hardhat-node