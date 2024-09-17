# db

This module contains all logic, code, and definitions relating to the management and definitions of database schemas and models, and makes use of sqlc.

## Folder Structure

### schema

_Stores all database migration files generated using Goose._

### sqlc

_Stores all sqlc-related code and configs._

#### <project_name>

##### queries

_Stores all query files to be used in `sqlc generate` command to generate code based on the provided sql queries._

##### sqlc.yaml

_Defines the required sqlc-related configs for the specific application/service. All output is pushed to the respective module._
