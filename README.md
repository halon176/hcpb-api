# HCrypto Price Bot API

⚠️ Current version is not ready for production. It is still under active development and may contain bugs. Use it at your own risk.

This web API was created to provide a logic and authentication layer for the [HCrypto Price Bot](https://github.com/halon176/h-crypto-price-bot). It currently collects data from the bot's calls and stores it in a database. It also provides a way to retrieve the data from the database to an interface.

The aim is to maintain this API as efficiently as possible to minimize the latency of the bot's calls. It is built entirely in pure Go, without any framework, to avoid overhead. It is built on top of a PostgreSQL database and uses the `github.com/lib/pqx` driver to connect to the database. JSON responses are generated directly by the DBMS to enhance performance.

## Security

The API uses an API key named `X-API-KEY` to authenticate requests. This key is specified in the `.env` file and protects all routes.
