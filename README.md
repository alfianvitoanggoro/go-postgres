# go-postgres
Submodule from go-libs-playground for exploration library postgres on golang

# 📝 Structure
<pre>
go-postgres/
├── config/
│   └── config.go
├── db/
│   └── pool.go
├── internal/
│   ├── app
│   │    └── user/
│   │       ├── handler.go        # handler HTTP
│   │       ├── model.go
│   │       ├── repository.go
│   │       └── router.go         # router khusus user
│   └── router/
│       └── router.go             # register semua grup router (user, dll)
├── docker-compose.yml
├── Dockerfile
├── main.go
├── go.mod
├── go.sum
├── .env
├── .gitignore
├── .air.toml
└── readme.md
</pre>
