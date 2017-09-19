# Buffalo example - GoWroc simple blog

## Running the database

```
./run-postgres.sh
```

## Creating databases

```
buffalo db create -a
```

The password is defined in `../run-postgres.sh` script.

## Starting the development server

```
buffalo dev
```

The application should be available on http://localhost:3000.

## Internationalisation

To access the site in Polish you should use one of two ways:

- Send `Accept-Language` HTTP header with a value of `pl-pl`
- Define a cookie on the page:

	```
	document.cookie="lang=pl-pl"
	```

## Generating database entities

To generate a post model

```
buffalo db generate model post title:string abstract:text content:text
```

To generate a post resource (with CRUD GUI):

```
buffalo generate resource post title:string abstract:text content:text
```

## Authentication

To generate an authentication intergration with Github:

```
buffalo generate goth github
```

In order to set up the integration you should create `.env` file (basing on `.env.sample`) with appropriate Github properties.

## REPL console

To access Buffalo's REPL console:

```
buffalo console
...
gore>
```

## Building

To builld a single binary:

```
buffalo build
```

This should create a single binary file in `bin/` directory.
