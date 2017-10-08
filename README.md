# rethinkcli

#### Table of Contents

1. [Overview](#overview)
2. [Description - What this tool does and why it is useful](#description)
3. [Usage](#usage)
4. [Limitations - compatibility, etc.](#limitations)

## Overview

This is a very simple client for RethinkDB written in Golang. It allows to simply dump the contents of any table.

## Description

This client tool uses the [GoRethink](https://github.com/GoRethink/gorethink) driver v3. It can be run as both a docker image (recommended) or natively after compiling the go program.

Querying a RethinkDB database is not always straightforward. There is no client you can install similar to how one would install a sql client for an RDBMS system. This tool aims to make querying of RethinkDB data simpler by accepting very few parameters (db server host, database and table) and dumping all the contents of the given table. It is currently available to run it on CLI only & does not support advanced filtering options.

This tool is publicly available as a minimal docker image (only 2MB) as `anoop/rethingo`. Using this you can very quickly & easily run a RethinkDB client as `docker run anoop/rethingo 54.174.187.41:28015 test movies`.

## Usage

To use this tool, it is recommended to use it as a docker image / container. Simply build a docker image using the included `Dockerfile`. 

Clone this repo and run the following commands:
```
$ cd rethinkcli
$ docker build -t rethinkcli .
```
Now you can run this against an existing RethinkDB server as:
```
docker run rethinkcli <db-url> <db> <table>
```
For eg: [For a server running on the host `54.174.187.41`, with a table `movies` inside a database `test`].

Note: If the RethinkDB is running on the default port of `28015`, then you can omit the `:28015` part.

```
$ docker run rethinkcli 54.174.187.41:28015 test movies
 1:
             id: 0383bc8a-e0bf-4028-9a47-8a8f6826310d
           rank: 240
         rating: 8
          title: Jurassic Park
          votes: 402866
           year: 1993
           .....snip
           .....
253:
             id: fd41be4a-7f01-435d-9e50-0ab7d422666a
           rank: 80
         rating: 8.4
          title: Braveheart
          votes: 552365
           year: 1995

		====  Total rows returned: 253 ====
```

You can also run this as a simple go program. You need to have golang installed. To run it as a go program, clone this repo and then run the following commands:
```
$ cd rethinkcli
$ go get gopkg.in/gorethink/gorethink.v3
$ go build src/demo.go
$ ./demo <db-url> <db> <table>
```

## Limitations

This script currently does not support authentication or TLS certificates. These updates are coming soon.


