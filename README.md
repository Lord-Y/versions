# versions-api

This api permit to centralize all your applications versions deployed in order to know when it has been deployed and his content.

## Definitions

Workload: it mean which team is related to this deployment
Platform: release to which platform your application has been deployed like development, staging, preproduction or production
Environment: on each platform, multiple environment can be deployed like development or integration

## How it works

You need to define OS environment variables like so:

```bash
# for mysql
export SQL_DRIVER=mysql
export DB_URI must be like so: `USERNAME:PASSWORD@tcp(HOST:PORT)/DB_NAME?charset=utf8&autocommit=true&multiStatements=true&maxAllowedPacket=0&interpolateParams=true&parseTime=true`
# for postgres
export SQL_DRIVER=postgres
export DB_URI must be like so: `postgres://USERNAME:PASSWORD@HOST:PORT/DB_NAME?sslmode=disable`
# If you want to enable a redis
export REDIS_ENABLED=true
export REDIS_URI=redis://:xxxxx@host:port/db
# or 
export REDIS_ENABLED=true
export REDIS_URI=redis://:xxxxx@host:port
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
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.0.0","workload": "teamX", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "ongoing"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamX", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.3.0","workload": "teamX", "environment":"dev", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "failed"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.2.0","workload": "teamX", "environment":"integration", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "ongoing"}'

curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.0.0","workload": "teamY", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamY", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.3.0","workload": "teamY", "environment":"dev", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "ongoing"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.2.0","workload": "teamY", "environment":"integration", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'

curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.0.0","workload": "teamZ", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamZ", "environment":"production", "platform": "production", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "ongoing"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.3.0","workload": "teamZ", "environment":"dev", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.2.0","workload": "teamZ", "environment":"integration", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "0.2.0","workload": "teamZ", "environment":"test", "platform": "development", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'
curl -XPOST 0:8080/api/v1/versions/create -H 'Content-type: application/json' -d '{"version": "1.1.0","workload": "teamZ", "environment":"preproduction", "platform": "preproduction", "changelogURL": "https://jsonplaceholder.typicode.com/", "raw": "{\"a\":\"b\"}", "status": "deployed"}'

# post requests with form format mode
curl -XPOST 0:8080/api/v1/versions/create -d 'version=1.1.0&workload=teamX&environment=production&platform=production&changelogURL=https://jsonplaceholder.typicode.com/&raw=rawContent&status=ongoing'
curl -XPOST 0:8080/api/v1/versions/create -d 'version=1.3.0&workload=teamX&environment=production&platform=production&changelogURL=https://jsonplaceholder.typicode.com/&raw=rawContent&status=deployed'
curl -XPOST 0:8080/api/v1/versions/create -d 'version=1.2.0&workload=teamX&environment=production&platform=production&changelogURL=https://jsonplaceholder.typicode.com/&raw=rawContent&status=failed'
```

Each `POST` return the id of the deployment like so `{"versionId":14}`.
To update the status of the deployment:
```bash
curl -XPOST 0:8080/api/v1/versions/update/status -H 'Content-type: application/json' -d '{"versionId": "14","status": "deployed"}'
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