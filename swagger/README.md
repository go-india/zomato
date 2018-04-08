# swagger

Swagger generates a [Go](http://golang.org/) client library for accessing [Zomato API](https://developers.zomato.com/api) using [Swagger Specifications](https://swagger.io/specification/) provided by [Zomato](https://developers.zomato.com/swagger.json).

## Requirements

* Unix or Linux based OS
* Go 

## Usage

To generate a client library from `swagger.json` file:

```bash
$ make
```

To clean previously generated client library:

```bash
$ make clean
```

To fetch new `swagger.json` from Zomato API servers:

```bash
$ make fetch
```

### Caveats
* The `swagger.json` [provided by Zomato](https://developers.zomato.com/swagger.json) is invalid, so we have included a fixed `swagger.json` file for the user in the folder.

