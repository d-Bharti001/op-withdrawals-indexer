# OP Withdrawals Indexer

Withdrawals indexer for an **OP Stack Rollup**

## Services

1. Indexer
2. API

### Indexer

Fetches L2 and L1 chain events, and saves withdrawals in PostgreSQL.

* Build the executable with `just build-indexer`
* Run `./bin/indexer --help` to see the required arguments to run the program

**L2 Scanner**

* Fetches `L2ToL1MessagePasser` contract's `MessagePassed` events, and saves the withdrawals, along with any decodable information into the database tables `withdrawals`, `withdrawal_info`, and `tokens`.

**L1 Scanner**

* Fetches `OptimismPortal2` contract's `WithdrawalProvenExtension1` and `WithdrawalFinalized` events, and saves their information into the database tables `withdrawal_proven_txs` and `withdrawal_finalized_txs`.
* For withdrawal proven transactions, `WithdrawalProven` event is not captured, and the table `withdrawal_proven_txs` requires a `proof_submitter` to be recorded. Only `WithdrawalProvenExtension1` event is captured for this purpose.

### API

For frontend usage.

* Build the executable with `just build-api`
* Run `./bin/api --help` to see the required arguments to run the program

List of APIs:
* `/withdrawal-history` [GET]  
    Get withdrawals history related to an address

Swagger documentation:
* Swagger docs are available at the endpoint `docs/` [GET]
* To test the API service locally, run the program and visit `http://localhost:<HTTP_PORT>/docs/`.

## Running the services with env

**Tools used**

* `go` (version 24+ should work)
* `docker` (or OrbStack for MacOS): used just for the database. A separate PostgreSQL installation would also work.
* `just` https://github.com/casey/just?tab=readme-ov-file#installation
* `golang-migrate` https://github.com/golang-migrate/migrate
* `solc` https://docs.soliditylang.org/en/latest/installing-solidity.html (not required to run / set up the app)
* `abigen v2` https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings-v2 (not required to run / set up the app)
* `swagger` https://github.com/swaggo/swag?tab=readme-ov-file#getting-started

**Database setup**

The project contains a docker compose file for setting up a PostgreSQL database.

* Build the container (using [docker-compose.yml](./docker-compose.yml)):
    ```sh
    docker compose up -d
    ```
* Set up the relations:
    ```sh
    just migrate-up
    ```

**Project setup**

* Install the dependencies:
    ```sh
    go mod tidy
    ```

**Running with env file**

The project provides `.env.sample` which contains the environment variables needed to run the services.

* Copy into `.env`, and provide values as needed
    ```sh
    cp .env.sample .env
    ```
* Build the services. They are available as `./bin/indexer` and `./bin/api`.
    ```sh
    just build-all
    ```
* Run a service, for example:
    ```sh
    set -a && source .env && set +a
    ./bin/api
    ```

**Running as systemd services**

After building the executables (available in `./bin/`), they can be run as systemd services.  
See the service files located in `./systemd-services/` for details.

## Bridge Frontend

The `web/` subdirectory is a git submodule which points to another git repo `https://github.com/d-Bharti001/op-bridge-frontend.git`.  
The frontend application, based on **ReactJS**, is compatible with this **Golang** backend, and provides an interface for:
* deposits
* withdrawal initiation
* withdrawal management (i.e. proving and finalization)

Initialize the module with:
```sh
git submodule update --init --recursive
```

Then go through its README (`web/README.md`) for steps for installation and deployment.

## References

* [Optimism Legacy Indexer](https://github.com/ethereum-optimism/optimism-legacy/tree/develop/indexer)
