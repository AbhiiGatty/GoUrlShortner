# GoUrlShortner :rocket:
This is a url shortening microservice written in Go lang

> Note: The Docker image that we are going to build is already optimized where we utilize muti-stage build to bring the image size down from 300MB+ to 13.1MB for the GoLang application

## Requirements :pencil:
* An machine with following installed in the OS
  1. GoLang
  2. Docker

## Getting Started :runner:
1. Fork this directory under your namespace
2. Run the following commands
```bash
$ mkdir ~/workspace/
$ cd ~/workspace/
# Note: In the command below 'AbhiiGatty' should be changed to your respective username
$ git clone git@github.com:AbhiiGatty/GoUrlShortner.git
$ cd ~/workspace/GoUrlShortner/
```

## Local Development :computer:
1. We need to build the images, start and interact with the containers
> We have a Makefile which executes a set of instructions that should make setup a lot easier
```bash
# This will be RUN ONLY ONCE in your machine
$ docker network create docker-internal-network
# This will build both webapp and database application without using any cached layers 
$ make build_all

# If we want to rebuild the Postgres application
$ make build_db
# If we want to rebuild the GoLang application
$ make build_app

# When we want to just start the services/containers
$ make start_all

# When we want to attach to the container running the Postgres application 
$ make attach_to_db
# When we want to attach to the container running the GoLang application 
$ make attach_to_app
```
2. Setup the postgres database with the required changes
```bash
# Run to attach to the container running the Postgres application 
$ make attach_to_db

# The following commands are run inside Postgres application container
$ psql -U postgres

# Create a custom user for the project
CREATE ROLE gus_user;
ALTER USER "gus_user" WITH PASSWORD 'gus_user@123';
ALTER ROLE "gus_user" WITH LOGIN;
# Create the database
CREATE DATABASE go_url_shortner;
# Connect to the database
\c go_url_shortner;
# Create table
CREATE TABLE url_map (
	id serial PRIMARY KEY,
	"fullUrl" VARCHAR (2048) UNIQUE NOT NULL,
	"shortUrlCode" VARCHAR (7) NOT NULL UNIQUE,
	"createdAt" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
# Give gus_user permission to access all data
GRANT USAGE ON SCHEMA public TO gus_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO gus_user;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO gus_user;

```
3. Execute the GoLang application
```bash
# Run to attach to the container running the GoLang application 
$ make attach_to_app

# The following commands are run inside GoLang application container
$ ./GoUrlShortner
{"level":"info","msg":"Successfully connected from database -go_url_shortner","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: e4b09c9 for URL: http://12factor.net/","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: 8450f5a for URL: http://25.media.tumblr.com/tumblr_m3ej2dwCGM1qzw1qyo1_500.gif","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: b92b9ce for URL: http://38.media.tumblr.com/a0e4af72980aec0a152c42cfca06e3b1/tumblr_n8h39wBOC31sg98gmo1_400.gif","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: d211e55 for URL: http://67.media.tumblr.com/3b6e20716fa7b29ddb1467c9f3ad3e86/tumblr_o89e596wTT1v54eezo1_500.gif","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: 6a01e00 for URL: http://67.media.tumblr.com/a88971bfc9c5d47929cc29da38e00c7c/tumblr_nwqry416uK1u9b9ceo1_500.gif","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: 0c8ce33 for URL: http://67.media.tumblr.com/e9bfb9dd48ecbd85fa673d8f44c06549/tumblr_mrjofjetW11rdw1plo1_500.gif","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: 05c2d2a for URL: http://adambrown.info/p/wp_hooks/hook","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: 6efd3d7 for URL: http://adambrown.info/p/wp_hooks/hook/core_upgrade_preamble?version=4.4\u0026file=wp-admin/update-core.php","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: b87220d for URL: http://adambrown.info/p/wp_hooks/hook/wp_upgrade","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: b5f66d4 for URL: http://adodb.sourceforge.net/","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Inserted ShortUrlCode: 31f6546 for URL: http://afamilyinbaghdad.blogspot.com/","time":"2022-01-31T12:58:30Z"}
{"level":"info","msg":"Successfully disconnected to database - go_url_shortner","time":"2022-01-31T12:58:30Z"}
```
4. Check the Postgres application
```bash
# Run to attach to the container running the Postgres application 
$ make attach_to_db

# The following commands are run inside Postgres application container
$ psql -U postgres

# Connect to the database
postgres=# \c go_url_shortner;

go_url_shortner=# \d url_map;
                                        Table "public.url_map"
    Column    |           Type           | Collation | Nullable |               Default
--------------+--------------------------+-----------+----------+-------------------------------------
 id           | integer                  |           | not null | nextval('url_map_id_seq'::regclass)
 fullUrl      | character varying(2048)  |           | not null |
 shortUrlCode | character varying(7)     |           | not null |
 createdAt    | timestamp with time zone |           |          | CURRENT_TIMESTAMP
Indexes:
    "url_map_pkey" PRIMARY KEY, btree (id)
    "url_map_fullUrl_key" UNIQUE CONSTRAINT, btree ("fullUrl")
    "url_map_shortUrlCode_key" UNIQUE CONSTRAINT, btree ("shortUrlCode")


go_url_shortner=# select * from url_map limit 5;
 id |                                            fullUrl                                             | shortUrlCode |           createdAt
----+------------------------------------------------------------------------------------------------+--------------+-------------------------------
 23 | http://12factor.net/                                                                           | e4b09c9      | 2022-01-31 12:58:30.06341+00
 24 | http://25.media.tumblr.com/tumblr_m3ej2dwCGM1qzw1qyo1_500.gif                                  | 8450f5a      | 2022-01-31 12:58:30.079175+00
 25 | http://38.media.tumblr.com/a0e4af72980aec0a152c42cfca06e3b1/tumblr_n8h39wBOC31sg98gmo1_400.gif | b92b9ce      | 2022-01-31 12:58:30.080768+00
 26 | http://67.media.tumblr.com/3b6e20716fa7b29ddb1467c9f3ad3e86/tumblr_o89e596wTT1v54eezo1_500.gif | d211e55      | 2022-01-31 12:58:30.082043+00
 27 | http://67.media.tumblr.com/a88971bfc9c5d47929cc29da38e00c7c/tumblr_nwqry416uK1u9b9ceo1_500.gif | 6a01e00      | 2022-01-31 12:58:30.083398+00
(5 rows)

go_url_shortner=#

```

#### Credits :handshake:
- https://www.chicagoitsystems.com/plain-text-file-of-urls/
