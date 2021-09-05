### structure folder
```
.
├── README.md
├── business => businnes layer
│   ├── label
│   │   └── domain.go
│   ├── note
│   │   ├── domain.go
│   │   ├── repository
│   │   │   ├── orm.go
│   │   │   └── postgresql.go
│   │   └── usecase.go
│   └── user
├── driver => driver layer
│   ├── database
│   │   ├── orm
│   │   │   └── orm.go
│   │   └── postgresql
│   └── thirdparties
├── files
│   ├── CleanArchitecture.jpeg
│   ├── db.sql
│   ├── list_api.txt
│   └── note.txt
├── go.mod
├── go.sum
├── helper
├── pkg => thirdparties / libbrary folder
├── server
│   ├── bootstrap
│   │   └── routers.go
│   ├── handlers
│   ├── main.go
│   └── middleware
│       └── jwt.go
└── swagger
    ├── swagger.yaml
    └── swagger2ex.yaml
```