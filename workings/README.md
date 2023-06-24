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
      