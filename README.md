# OP Withdrawals Indexer

## Services
1. Indexer
2. API

### Indexer

Fetches L2 and L1 chain events, and saves, updates withdrawals in PostgreSQL.

* Build the executable with `just build-indexer`
* Run `./bin/indexer --help` to see the required arguments to run the program

### API

For frontend usage.

* Build the executable with `just build-api`
* Run `./bin/api --help` to see the required arguments to run the program

List of APIs:
* `/withdrawal-history` [GET]  
    Get withdrawals history related to an address

Swagger documentation:
* Run the program locally, and visit `http://localhost:<HTTP_PORT>/swagger/` to view swagger docs.

## Running the services with env

The project provides `.env.sample` which contains the environment variables needed to run the services.

* Copy into `.env`, and provide values as needed
    ```sh
    cp .env.sample .env
    ```
* Build the services
    ```sh
    just build-all
    ```
* Run a service, for example:
    ```sh
    set -a && source .env && set +a
    ./bin/api
    ```

## Tools used

These additional tools are used in the project:
* `just` (it can also be used for execution, e.g. for building the service executables)
* `golang-migrate`
* `solc`
* `abigen v2`
* `swagger`

Required external components:
* Postgresql database  
    Build with docker: `docker compose up -d`  
    See: `docker-compose.yml`

## References

* [Optimism Legacy Indexer](https://github.com/ethereum-optimism/optimism-legacy/tree/develop/indexer)
