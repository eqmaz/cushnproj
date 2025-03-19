### 1. Install oapi-codegen (required)
The REST API handlers are code-generated with openapi-generator. <br>
So, the openapi.yaml file is used to generate the API boilerplate code.<br>
(This API is implemented with Fiber because it's really fast).

```bash
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```
_In addition to generating the API boilerplate, we can also generate the API docs._ 

### 1.1 _(optional)_ Install 'Redoc' and/or 'widdershins'
This is for generating API documentation as either HTML or Markdown.<br> 
```bash
  npm install -g redoc-cli    # For generating HTML docs
  npm install -g widdershins  # For generating Markdown docs
```

Once installed, you can generate API html docs with:
```bash
  make gen-docs-html  #Creates HTML docs
  make gen-docs-md    #Creates Markdown docs
  make gen-docs-all   #Creates both HTML and Markdown docs
```

### 2. Install migrate tool for Go
This project has a local database for development purposes. <br>
You'll **need Docker installed to run the local database**. <br>
The database schema is managed with the 'migrate' tool. <br>
You'll need to start the local db container and first migration with structure and seeds.

```bash
  make migrate-install  # Installs the migrate tool
  make local-db-up      # Starts the local db docker container
  make migrate          # Runs the first migration with structure and seeds
```
Note that the ``make run`` command will also run any new migrations (and will ensure the local db is up).

### 3. Install slqc (optional) dev dependency for Model (re)generation
To save buckets of time when DB models are under heavy development you 
can use slqc to generate the data transfer objects (DTOs) from the database tables. <br>
```bash
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### 4. Install jq (optional) for local version-bump automation
```bash
  sudo apt install jq
  chmod +x ./cmd/version/version.sh
```
