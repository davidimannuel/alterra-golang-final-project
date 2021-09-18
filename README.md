alterra golang final project

App name : Keep Remind !!!
### structure folder
```
.
├── README.md
├── businesses
│   ├── context.go
│   └── note
│       ├── domain.go
│       └── usecase.go
├── configs
│   ├── config.go
│   └── constants.go
├── drivers
│   ├── postgresql
│   │   └── db.go
│   └── repositories
│       ├── label
│       └── note
│           ├── model.go
│           └── postgresql.go
├── files
│   ├── CleanArchitecture.jpeg
│   ├── db.sql
│   ├── list_api.txt
│   └── note.txt
├── go.mod
├── go.sum
├── helpers
│   └── str
│       └── time.go
├── server
│   ├── bootstraps
│   │   ├── app.go
│   │   ├── routes
│   │   │   └── v1
│   │   │       └── note.go
│   │   └── routes.go
│   ├── handlers
│   │   ├── handler.go
│   │   └── note
│   │       ├── http.go
│   │       ├── request
│   │       │   └── json.go
│   │       └── response
│   │           └── json.go
│   ├── main.go
│   └── middlewares
│       └── jwt.go
└── swagger
    ├── swagger.yaml
    └── swagger2ex.yaml
```