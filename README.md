# GoUrlShortner
This is a url shortening microservice written in Go lang

> Note: The Docker image that we are going to build is already optimized where we utilize muti-stage build to bring the image size down from 300MB+ to 13.1MB for the GoLang application

## Requirements
* A machine with following installed
    1. GoLang
    2. Docker

## Getting Started
1. Fork this directory under your namespace
2. Run the following commands
```bash
$ mkdir ~/workspace/
$ cd ~/workspace/
# Note: In the command below 'AbhiiGatty' should be changed to your respective username
$ git clone git@github.com:AbhiiGatty/GoUrlShortner.git
$ cd ~/workspace/GoUrlShortner/
```

## Local Development
1. We need to build the images, start and interact with the containers
> We have a Makefile which executes a set of instructions that should make setup a lot easier
```bash
# This will build both webapp and database application [RUN ONLY ONCE]
$ make build_all
# If we want to rebuild the Postgres application
$ make build_db
# If we want to rebuild the GoLang application
$ make build_app
# When we want to just start the services
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
```

### Credits
- https://www.chicagoitsystems.com/plain-text-file-of-urls/