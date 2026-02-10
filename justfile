set shell := ["bash", "-ceu", "-o", "pipefail"]
set dotenv-load := true

DB_MIGRATIONS_PATH := "./internal/database/migrations"
CONTRACTS_SOL_INTERFACES_PATH := "./internal/contracts/solidity/interfaces"
CONTRACTS_ABI_PATH := "./internal/contracts/abi"
CONTRACTS_BINDINGS_PATH := "./internal/contracts/bindings"
SWAGGER_DOCS_PATH := "./internal/api/docs"
BUILD_PATH := "./bin"

# Build the indexer service executable
build-indexer:
  go build -o {{BUILD_PATH}}/indexer ./cmd/indexer/*

# Build the api service executable
build-api:
  go build -o {{BUILD_PATH}}/api ./cmd/api/*

# Build both the service executables
build-all: build-indexer build-api

# Create a new migration
migrate-create name:
  migrate create -seq -ext sql -dir={{DB_MIGRATIONS_PATH}} {{name}}

# Apply all migrations
migrate-up:
  migrate -path={{DB_MIGRATIONS_PATH}} -database=$DB_CONN_STR up

# Rollback migrations (N steps, default 1)
migrate-down steps="1":
  migrate -path={{DB_MIGRATIONS_PATH}} -database=$DB_CONN_STR down {{steps}}

# Generate abi files from Solidity interfaces
# Add / remove the commands according to available contracts
gen-contract-abi:
  solc --abi {{CONTRACTS_SOL_INTERFACES_PATH}}/L1/* --pretty-json --json-indent 4 --overwrite --output-dir {{CONTRACTS_ABI_PATH}}/L1
  solc --abi {{CONTRACTS_SOL_INTERFACES_PATH}}/L2/* --pretty-json --json-indent 4 --overwrite --output-dir {{CONTRACTS_ABI_PATH}}/L2
  solc --abi {{CONTRACTS_SOL_INTERFACES_PATH}}/universal/* --pretty-json --json-indent 4 --overwrite --output-dir {{CONTRACTS_ABI_PATH}}/universal
  solc --abi {{CONTRACTS_SOL_INTERFACES_PATH}}/vendor/* --pretty-json --json-indent 4 --overwrite --output-dir {{CONTRACTS_ABI_PATH}}/vendor
  rm {{CONTRACTS_ABI_PATH}}/{L1,L2}/{ICrossDomainMessenger.abi,IStandardBridge.abi,IERC721Bridge.abi}

# Generate go contract bindings from abi files
# Add / remove the commands according to available contracts
gen-contract-bindings:
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L1/IL1CrossDomainMessenger.abi --pkg bindings --type IL1CrossDomainMessenger --out {{CONTRACTS_BINDINGS_PATH}}/IL1CrossDomainMessenger.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L1/IL1ERC721Bridge.abi --pkg bindings --type IL1ERC721Bridge --out {{CONTRACTS_BINDINGS_PATH}}/IL1ERC721Bridge.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L1/IL1StandardBridge.abi --pkg bindings --type IL1StandardBridge --out {{CONTRACTS_BINDINGS_PATH}}/IL1StandardBridge.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L1/IOptimismPortal2.abi --pkg bindings --type IOptimismPortal2 --out {{CONTRACTS_BINDINGS_PATH}}/IOptimismPortal2.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L1/ISystemConfig.abi --pkg bindings --type ISystemConfig --out {{CONTRACTS_BINDINGS_PATH}}/ISystemConfig.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L2/IL2CrossDomainMessenger.abi --pkg bindings --type IL2CrossDomainMessenger --out {{CONTRACTS_BINDINGS_PATH}}/IL2CrossDomainMessenger.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L2/IL2ERC721Bridge.abi --pkg bindings --type IL2ERC721Bridge --out {{CONTRACTS_BINDINGS_PATH}}/IL2ERC721Bridge.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L2/IL2StandardBridge.abi --pkg bindings --type IL2StandardBridge --out {{CONTRACTS_BINDINGS_PATH}}/IL2StandardBridge.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/L2/IL2ToL1MessagePasser.abi --pkg bindings --type IL2ToL1MessagePasser --out {{CONTRACTS_BINDINGS_PATH}}/IL2ToL1MessagePasser.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/universal/ICrossDomainMessenger.abi --pkg bindings --type ICrossDomainMessenger --out {{CONTRACTS_BINDINGS_PATH}}/ICrossDomainMessenger.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/universal/IERC721Bridge.abi --pkg bindings --type IERC721Bridge --out {{CONTRACTS_BINDINGS_PATH}}/IERC721Bridge.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/universal/IStandardBridge.abi --pkg bindings --type IStandardBridge --out {{CONTRACTS_BINDINGS_PATH}}/IStandardBridge.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/vendor/IERC1155.abi --pkg bindings --type IERC1155 --out {{CONTRACTS_BINDINGS_PATH}}/IERC1155.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/vendor/IERC165.abi --pkg bindings --type IERC165 --out {{CONTRACTS_BINDINGS_PATH}}/IERC165.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/vendor/IERC20.abi --pkg bindings --type IERC20 --out {{CONTRACTS_BINDINGS_PATH}}/IERC20.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/vendor/IERC721.abi --pkg bindings --type IERC721 --out {{CONTRACTS_BINDINGS_PATH}}/IERC721.go
  abigen --v2 --abi {{CONTRACTS_ABI_PATH}}/vendor/IERC20Metadata.abi --pkg bindings --type IERC20Metadata --out {{CONTRACTS_BINDINGS_PATH}}/IERC20Metadata.go

# Generate swagger docs for API
gen-docs:
  swag init \
  -g ./main.go \
  -d cmd/api,internal/api/handlers/withdrawals,\
  internal/api/handlers/response,\
  internal/api/dbmodels \
  -o {{SWAGGER_DOCS_PATH}}

  swag fmt
