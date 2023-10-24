# versions [![CircleCI](https://circleci.com/gh/Lord-Y/versions/tree/main.svg?style=svg)](https://circleci.com/gh/Lord-Y/versions?branch=main)

`versions` is a single binary that hold both frontend UI and API that permit to centralize all your applications versions deployed in order to know when it has been deployed and his content.

## Definitions

`Workload`: it mean which team is related to this deployment.

`Platform`: release to which platform your application has been deployed like development, staging, preproduction or production.

`Environment`: on each platform, multiple environment can be deployed like development or integration.

## How it works

You need to define OS environment variables like so:

```bash
# for mysql
export SQL_DRIVER=mysql
export DB_URI="USERNAME:PASSWORD@tcp(HOST:PORT)/DB_NAME?charset=utf8&autocommit=true&multiStatements=true&maxAllowedPacket=0&interpolateParams=true&parseTime=true"
# for postgres
export SQL_DRIVER=postgres
export DB_URI="postgres://USERNAME:PASSWORD@HOST:PORT/DB_NAME?sslmode=disable"
# If you want to enable a redis
export REDIS_ENABLED=true
export REDIS_URI=redis://:xxxxx@host:port/db
# or 
export REDIS_ENABLED=true
export REDIS_URI=redis://:xxxxx@host:port
```

SQL_DRIVER can only be `mysql` or `postgres`
At startup, the DB initialization or migration will be handle.

Here are the supported versions:
- postgres 10 and above
- mariadb 10.x and above
- mysql 5.7 and above

## Running port number

Port number can be change:

```bash
export APP_PORT=8080
```

## Tests

Unit testing are covered.

```bash
go test -v ./... -coverprofile=coverage.out
# go tool cover -html=coverage.out
# go tool cover -func=coverage.out
```

## Queries

For POST request, both content-type 

```bash
# post requests with json
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.0.0","workload": "teamX", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "ongoing"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamX", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.3.0","workload": "teamX", "environment":"dev", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "failed"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.2.0","workload": "teamX", "environment":"integration", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "ongoing"}'

curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.0.0","workload": "teamY", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamY", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.3.0","workload": "teamY", "environment":"dev", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "ongoing"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.2.0","workload": "teamY", "environment":"integration", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'

curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.0.0","workload": "teamZ", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamZ", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "ongoing"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.3.0","workload": "teamZ", "environment":"dev", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.2.0","workload": "teamZ", "environment":"integration", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "0.2.0","workload": "teamZ", "environment":"test", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamZ", "environment":"preproduction", "platform": "preproduction", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'

# post requests with form format mode
curl -XPOST 0:8080/api/v1/create -d 'version=1.1.0&workload=teamX&environment=production&platform=production&changelogURL=https://jsonplaceholder.typicode.com/&raw=rawContent&status=ongoing'
curl -XPOST 0:8080/api/v1/create -d 'version=1.3.0&workload=teamX&environment=production&platform=production&changelogURL=https://jsonplaceholder.typicode.com/&raw=rawContent&status=deployed'
curl -XPOST 0:8080/api/v1/create -d 'version=1.2.0&workload=teamX&environment=production&platform=production&changelogURL=https://jsonplaceholder.typicode.com/&raw=rawContent&status=failed'
```

Each `POST` return the id of the deployment like so `{"versionId":14}`.
To update the status of the deployment:
```bash
curl -XPOST 0:8080/api/v1/update/status -H 'Content-type: application/json' -d '{"versionId": "14","status": "deployed"}'
```

Here is an example of DB content:
```bash
+-------------+----------+---------------+---------------+---------+---------------------------------------+---------------+----------+---------------------+
| versions_id | workload | platform      | environment   | version | changelog_url                         | raw           | status   | date                |
+-------------+----------+---------------+---------------+---------+---------------------------------------+---------------+----------+---------------------+
|           1 | teamX    | production    | production    | 1.0.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | ongoing  | 2020-12-02 19:58:10 |
|           2 | teamX    | production    | production    | 1.1.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:10 |
|           3 | teamX    | development   | dev           | 1.3.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | failed   | 2020-12-02 19:58:10 |
|           4 | teamX    | development   | integration   | 1.2.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | ongoing  | 2020-12-02 19:58:10 |
|           5 | teamY    | production    | production    | 1.0.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:10 |
|           6 | teamY    | production    | production    | 1.1.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:10 |
|           7 | teamY    | development   | dev           | 1.3.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | ongoing  | 2020-12-02 19:58:10 |
|           8 | teamY    | development   | integration   | 1.2.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:10 |
|           9 | teamZ    | production    | production    | 1.0.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:10 |
|          10 | teamZ    | production    | production    | 1.1.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | ongoing  | 2020-12-02 19:58:10 |
|          11 | teamZ    | development   | dev           | 1.3.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:10 |
|          12 | teamZ    | development   | integration   | 1.2.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:10 |
|          13 | teamZ    | development   | test          | 0.2.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:10 |
|          14 | teamZ    | preproduction | preproduction | 1.1.0   | https://jsonplaceholder.typicode.com/ | {\"a\":\"b\"} | deployed | 2020-12-02 19:58:11 |
+-------------+----------+---------------+---------------+---------+---------------------------------------+---------------+----------+---------------------+
```

## Get latest deployment

Sometimes, you may need to fetch dynamically latest deployment.
You can fetch it like so:
```bash
curl '0:8080/api/v1/read/environment/latest?workload=teamX&environment=production&platform=production'
```
It will query latest version with `status` equal to `deployed` or `completed`.

You can also use:
```bash
curl '0:8080/api/v1/read/environment/latest/whatever?workload=teamX&environment=production&platform=production'
```
It will fetch the latest version deployed without checking the `status` of the deployment.

## Metrics

Golang metrics are built-in in our API.

We also add last 10 days deployments by status on `curl 0:8080/api/v1/stats/latest` but also on `curl 0:9101/metrics |grep versions_last_deployments_by_status`

## Enable logging on health(z) paths

To enable logging on health(z) path, set environment variable to any value for example: `APP_SKIP_PATH_DISABLED=true`.

## What does it looks like

The home page show you the last deployments from all workloads:

![Home](./screenshots/home.png "Home")

The statistic page show you how much deployments have by done:

![Staticstics](./screenshots/staticstics.png "Staticstics")

The platform page show you deployments from selected plaform:

![Workloads platforms](./screenshots/workloads_platforms.png "Workloads platforms")

The environment page show you deployments from selected environment:

![Workloads environments](./screenshots/workloads_environments.png "Workloads environments")

The raw page show you the content of the version deployed:

![Raw version](./screenshots/raw_version.png "Raw version")

## Contributing

In order to contribute to our project, we need to configure git hooks.

You need first do install [golangci-lint](https://golangci-lint.run/usage/install/)

Then, enable the hook in our project:
```bash
git config core.hooksPath .githooks
```
When you need to run both `ui` and `api` you need to open 2 shells:
```bash
# api shell
export SQL_DRIVER=postgres
export DB_URI="postgres://USERNAME:PASSWORD@HOST:PORT/DB_NAME?sslmode=disable"
export APP_PORT=8081
go run main.go

# ui shell
cd ui
export API_GATEWAY=http://127.0.0.1:8081
npm run dev
```

## TODO

Use [vue-google-charts](https://github.com/devstark-com/vue-google-charts) when it will be compatible with Vue 3.
