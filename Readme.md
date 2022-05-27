# LEARINING API

<img align="center" width="180px" src="https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/c7d894cb-8d37-4495-a454-89c868b12375/dcycwca-813a3b2d-1eae-4f6a-beab-27f1264b364b.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcL2M3ZDg5NGNiLThkMzctNDQ5NS1hNDU0LTg5Yzg2OGIxMjM3NVwvZGN5Y3djYS04MTNhM2IyZC0xZWFlLTRmNmEtYmVhYi0yN2YxMjY0YjM2NGIucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.KijY-p4GWjczqKcWqY3xgRmvPgK8SUgbHDdHsDIQvYc">

MAIL for send mail

## Getting started

1. Download swag by using:

```sh
$ go get -u github.com/swaggo/swag/cmd/swag
```

2. Run `swag init` in the project's root folder which contains the `main.go` file. This will parse your comments and generate the required files (`docs` folder and `docs/docs.go`).

```sh
$ swag init
```

3. Pull Image mysql:

```sh
$ docker pull mysql
```

4. Run Image mysql:

```sh
$ docker run --name=learning -e MYSQL_ROOT_PASSWORD=P@ssw0rd -e MYSQL_DATABASE=learning -p 3306:3306 -d mysql
```

```sh
$ go get -u github.com/swaggo/swag/cmd/swag
```

5. Run `go run main.go`
