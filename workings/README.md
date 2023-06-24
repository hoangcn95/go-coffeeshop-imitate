# go-coffeeshop-imitate

Imitating an event-driven microservices coffee shop application.

## Dev notes

- [doc flow](https://libraries.excalidraw.com/): Write pretty documentation (flow, design, infrastructure)
- [sql migration](https://github.com/golang-migrate/migrate): Migrate db
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
      