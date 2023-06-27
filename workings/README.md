# go-coffeeshop-imitate

Imitating an event-driven microservices coffee shop application.

## Dev notes

- [Doc flow](https://libraries.excalidraw.com/): Write pretty documentation (flow, design, infrastructure)
- [Sql migration](https://github.com/golang-migrate/migrate): Migrate db
  - Questions: How about versions ?
  - How to use ?
    - generating migration

      ```bash
        make migrategen NAME=${NAME}
      ```

    - version up

      ```bash
        make migrateup
      ```

    - version up
    
      ```bash
        make migratedown
      ```

- [Generate sql-boilerplate by sqlc](https://github.com/kyleconroy/sqlc): Generate sql query and model for avoiding boilerplate
  
  ```bash
    make sqlc
  ```

- [buf](https://github.com/bufbuild/buf): gRPC generation for decreasing complexity of natural protoc-gen-go, besides that got benefits from versioning control

  - buf installation
    ```bash
    brew install bufbuild/buf/buf
    ```
  
  - buf lint and generate stuff
    ```bash
    make bufall
    ```
  
## Technical debt

- How to add logging and tracing ?
- How to add observability (metric) ?
- How to add unit + integration tests ?
- [grpc] How to add retry mechanism in grpc_client for overcome hiccup network issues ?
- [grpc] How to add circuit breaker for push up system reliability ?

      