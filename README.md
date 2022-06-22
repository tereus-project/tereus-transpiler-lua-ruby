# Tereus Lua to Ruby transpiler

Converts Lua to Ruby.

## Command-line usage

You can use the following command to convert a Lua file to Ruby:

```shell
go run . <inputlua>
```

## Worker usage

The transpiler is designed to listen on a queue for Lua files to convert.

First, you need to copy the `.env.example` file to `.env` and fill in the values.

The Kafka endpoint and S3 bucket should be shared with the [Tereus API](https://github.com/tereus-project/tereus-api).

You can then start the worker with:

```shell
➜  go run .
INFO[0000] Connecting to Kafka...
INFO[0000] Connecting to MinIO...
INFO[0000] Starting transpiler job listener...
DEBU[0006] Job '48549e3e-a181-49f5-b204-0fc56df3e319' started
DEBU[0006] Downloading job files...
DEBU[0006] Downloading file 'mainlua'
DEBU[0006] Remixing file 'mainlua'
DEBU[0006] Uploading file 'mainrb'
DEBU[0007] Job '48549e3e-a181-49f5-b204-0fc56df3e319' completed
```

The worker will get transpilation submissions from a queue and update the status and result in another one.
