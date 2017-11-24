# redashman

`redashman` is a query management tool for Redash

## Installation

```bash
$ go get -u github.com/ariarijp/redashman
```

## Usage

```bash
$ export REDASH_URL=http://localhost:5000
$ export REDASH_API_KEY=YOUR_API_KEY
$ 
$ # List queries
$ redashman query list 100 1 --url $REDASH_URL --api-key $REDASH_API_KEY
$ 
$ # Show a query
$ redashman query show 1 --url $REDASH_URL --api-key $REDASH_API_KEY
$ 
$ # Create a new query with text from a file
$ echo "SELECT NOW();" > new_query.sql
$ redashman query create new_query.sql --url $REDASH_URL --api-key $REDASH_API_KEY
$ 
$ # Modify a query with text from a file
$ echo "SELECT NOW(), CURDATE();" > modified_query.sql
$ redashman query modify 1 modified_query.sql --url $REDASH_URL --api-key $REDASH_API_KEY
$ 
$ # Fork a query from an existing one
$ redashman query fork 1 --url $REDASH_URL --api-key $REDASH_API_KEY
$ 
$ # Archive a query
$ redashman query archive 1 --url $REDASH_URL --api-key $REDASH_API_KEY
```

## License

MIT

## Author

[ariarijp](https://github.com/ariarijp)
