# Sensedia Challenge API

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

You need to have installed the following software:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)
- [Python3](https://www.python.org/downloads/)
- [Pip](https://pip.pypa.io/en/stable/installing/)
- [Pipenv](https://pipenv.readthedocs.io/en/latest/install/#installing-pipenv)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Sqlx](https://github.com/launchbadge/sqlx/blob/main/sqlx-cli/README.md)

### Installing

After installing the prerequisites, you need to install the dependencies of the project.

```bash
pip install psycopg2
```

After installing the dependencies, you need to build the database container. Make sure you have your .env file with the correct values, as exemplified in the .env.sample.

```bash
 docker compose up -d
```

After building the database container, you need to run the migrations.

```bash
make run_migrations
```

After that, you need to run the application.

```bash
make run
```

You can also run the application with live reload using air or the make command provided below in the Makefile section.

## MakeFile

```bash
make build
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Live reload the application. If you already have air installed but our command isn't working, you can run it directly with air.

```bash
make watch
```

Run the test suite

```bash
make test
```

Clean up binary from the last build

```bash
make clean
```

Run the migrations

```bash
make run_migrations
```

Run the migrations rollback, you need to run this once for every migration you want to rollback.

```bash
make rollback_migrations
```

Populate the database with fake data

```bash
make populate_db
```
