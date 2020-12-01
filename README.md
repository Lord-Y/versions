# versions-api

This api stand to centralize all your applications versions deployed in order to know when it has been deployed and his content.

## Definitions

Workload: it mean which team is related to this deployment
Platform: release to which platform your application has been deployed like development, staging, preproduction or production
Environment: on each platform, multiple environment can be deployed like development or integration

## How it works

You need to define OS environment variables like so:

```bash
export SQL_DRIVER=xxxx
export DB_URI must be like so: `USERNAME:PASSWORD@tcp(HOST:PORT)/DB_NAME?charset=utf8&autocommit=true&multiStatements=true&maxAllowedPacket=0&interpolateParams=true`
export DB_URI must be like so: `postgres://USERNAME:PASSWORD@HOST:PORT/DB_NAME?sslmode=disable`
# If you want to enable a redis
export REDIS_ENABLED=true
export REDIS_URI=redis://xxxxx@host:port/db
# or 
export REDIS_URI=redis://xxxxx@host:port
```

SQL_DRIVER can only be `mysql` or `postgres`
At startup, the DB initialization or migration will be handle

## Running port number

Port number can be change:

```bash
export APP_PORT=8080
```

## Tests

Unit testing are covered.

```bash
go test -v
```

## Queries

For POST request, both content-type 

```bash
# post requests with json
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.0.0","workload": "teamX", "environment":"production", "platform": "production", "changelogURL": "changelogURL", "raw": "rawContent"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamX", "environment":"production", "platform": "production", "changelogURL": "changelogURL", "raw": "rawContent"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.3.0","workload": "teamX", "environment":"dev", "platform": "development", "changelogURL": "changelogURL", "raw": "rawContent"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.2.0","workload": "teamX", "environment":"integration", "platform": "development", "changelogURL": "changelogURL", "raw": "rawContent"}'
# post requests with form format mode
curl -XPOST 0:8080/api/v1/versions/create -d 'version=1.1.0&workload=teamX&environment=production&platform=production&changelogURL=changelogURL&raw=rawContent'
curl -XPOST 0:8080/api/v1/versions/create -d 'version=1.3.0&workload=teamX&environment=production&platform=production&changelogURL=changelogURL&raw=rawContent'
curl -XPOST 0:8080/api/v1/versions/create -d 'version=1.2.0&workload=teamX&environment=production&platform=production&changelogURL=changelogURL&raw=rawContent'

...
....
.....
```

Here is an example of DB content:
```bash
+-------------+----------+---------------+---------------+---------+---------------+------------+---------------------+
| versions_id | workload | platform      | environment   | version | changelog_url | raw        | date                |
+-------------+----------+---------------+---------------+---------+---------------+------------+---------------------+
|           1 | teamX    | production    | production    | 1.0.0   | changelogURL  | rawContent | 2020-11-20 15:56:49 |
|           2 | teamX    | production    | production    | 1.1.0   | changelogURL  | rawContent | 2020-11-20 15:56:49 |
|           3 | teamX    | development   | dev           | 1.3.0   | changelogURL  | rawContent | 2020-11-20 15:56:49 |
|           4 | teamX    | development   | integration   | 1.2.0   | changelogURL  | rawContent | 2020-11-20 15:56:49 |
|           5 | teamY    | production    | production    | 1.0.0   | changelogURL  | rawContent | 2020-11-20 15:56:49 |
|           6 | teamY    | production    | production    | 1.1.0   | changelogURL  | rawContent | 2020-11-20 15:56:50 |
|           7 | teamY    | development   | dev           | 1.3.0   | changelogURL  | rawContent | 2020-11-20 15:56:50 |
|           8 | teamY    | development   | integration   | 1.2.0   | changelogURL  | rawContent | 2020-11-20 15:56:50 |
|           9 | teamZ    | production    | production    | 1.0.0   | changelogURL  | rawContent | 2020-11-20 15:56:50 |
|          10 | teamZ    | production    | production    | 1.1.0   | changelogURL  | rawContent | 2020-11-20 15:56:50 |
|          11 | teamZ    | development   | dev           | 1.3.0   | changelogURL  | rawContent | 2020-11-20 15:56:51 |
|          12 | teamZ    | development   | integration   | 1.2.0   | changelogURL  | rawContent | 2020-11-20 15:56:51 |
|          13 | teamZ    | development   | test          | 0.2.0   | changelogURL  | rawContent | 2020-11-20 15:56:51 |
|          14 | teamZ    | preproduction | preproduction | 1.1.0   | changelogURL  | rawContent | 2020-11-20 15:56:52 |
+-------------+----------+---------------+---------------+---------+---------------+------------+---------------------+
```