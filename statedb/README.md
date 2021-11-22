# StateDB

An historical database for tabular data with loader for EOSIO state table & permissions.

## Installation

Install through [dfuse for EOSIO](..)

## Features

StateDB supports parallel ingestion by doing a first sharding pass on the chain history, splitting the different tables
being mutated in different shards, then ingesting linearly each table.

This allows ingestion of the whole history in a few hours.

## Documentation

See the `/v0/state` endpoints under https://docs.dfuse.io/reference/eosio/rest/

## Contributing

*Issues and PR related to FluxDB belong to this repository*

See the dfuse-wide [contribution guide](https://github.com/dfuse-io/dfuse#contributing)
if you wish to contribute to this code base.

## Technical Implementation

### Tablets

* ContractStateTablet - prefix `cst` hex identifier `0xB000`:
  row key format: `cst:<contract>:<table>:<scope>:<block_num>:<dbop_primary_key>`, row value type: ContractStateValue
  contains DBOps payer and new data
* AuthLinkTablet - prefix `al` hex identifier `0xB100`:
  row key format: `al:<account>:<block_num>:<link_auth_code>:<link_auth_type>`, row value type: AuthLinkValue contains
  Permission
* ContractTableScopeTablet - prefix `ctscp` hex identifier `0xB200`:
  row key format: `ctscp:<contract>:<table_name>:<block_num>:<scope>`, row value type: ContractTableScopeValue contains
  TableOps payer
* KeyAccountTablet - prefix `ka` hex identifier `0xB300`:
  row key format: `ka:<public_key>:<block_num>:<account+permission>`, row value none

### Singlets

* ContractABISinglet - prefix `abi` hex identifier `0xA000`:
  row key format: `abi:<contract>:<block_num>`, row value type: ContractABIValue contains RawABI
* TabletIndex - prefix `idx` hex identifier `0xFFFF`:
