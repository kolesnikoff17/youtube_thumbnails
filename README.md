# youtube_thumbnails

gRPC client and server for getting video thumbnails from YouTube.

## Notes about task

* Go version for this project is 1.18.2 (unfortunately, 
I don't have sudo permission on my PC to update it).
* I choose Redis as a repository for this project.

## Requirements

Docker/docker-compose.

## Description

Small gRPC proxy server for downloading thumbnails from YouTube. 
Client is a command line utility, with additional flag support ([see below](#supported-flags)), 
which takes YouTube video id or link ([example](#example)) and saves picture 
in ```client/thumbs``` directory. All responses cache on server side, default TTL 
for cache is one hour.

## How to start

```bash
$ make
```
This command creates two containers: gRPC Server on 8080 port and Redis with no exposed port on localhost,
and compiles client application (```app``` executable in project's root directory).
To stop containers you can type:
```bash
$ make down
or
$ docker-compose down
```
To run unit tests:
```bash
$ make test
or
$ go test ./...
```
For coverage report:
```bash
$ make report
```
but I highly recommend to use Goland IDE for this purpose.

## Supported flags

```
--async             : download multiple files concurrently
--force-update      : ignore existing Redis cache and download thumbnail from YouTube API
--verbose           : report about success download in stdout
```

## Example

```bash
$ ./app --async --force-update --verbose https://www.youtube.com/watch?v=nviEkurZlao 'youtube.com/watch?v=mMCTX58yZZY' jfBkqXE5qjQ "https://www.youtube.com/watch?v=V2UIOLJCbqU"
```
## Cache

Why did I choose Redis over Sqlite:
* Speed - constant time CRUD operations in Redis against O(n) (O(log(n)) 
with btree indexes) in SQL solutions.
* Persistence - good dumps configuration and ability to set TTL for cached data out of the box.
* Prevalence - Redis is widespread in web development for caching data, so another developer, 
who may work on this API (if someone would ever), may easily make out 
solutions that was implemented.
* NoSQL - key-value repository is more suitable for my problem than relations with attributes. 
Redis command syntax is less verbose than SQL. Also, in vacancy description you mentioned experience 
with NoSQL databases would be a plus.
