## Introduction

- Blog System developed in the Go language for practice.



## Technology used

- Golang，Gin
- MySQL，Gorm
- Swagger
- validator，JWT



## API Description

| URI | Description |
| --- | ------------ |
| /auth?app_key=xxx&app_secret=yyy | Auth |
| /api/v1/tag [get] | Get multiple tags |
| /api/v1/tags [post] | Add tags |
| /api/v1/tags/{id} [put] | Update tag |
| /api/v1/tags/{id} [delete] | Delete tag |
| /api/v1/article/{id} [get] | Get article |
| /api/v1/article [post] | Create article |
| /api/v1/article/{id} [put] | Update article |
| /api/v1/article/{id} [delete] | Delete article |
| /groups/{group_id}/group-users | グループにユーザーを追加する|
| /groups/{group_id}/group-users | グループに所属するユーザーを一覧取得する|
| /groups/{group_id}/group-users/{user_id} | グループからユーザーを脱退させる |
| /upload/file -F file=@"/Users/chengkun02/Downloads/a.txt" -F type=1 | Upload file |

```go
.
├── READEME.md
├── blog.sql
├── configs
│   └── config.yaml
├── docs
│   ├── docs.go
│   ├── sql
│   │   └── blog.sql
│   ├── swagger.json
│   └── swagger.yaml
├── global
│   ├── db.go
│   ├── setting.go
│   └── validator.go
├── go.mod
├── go.sum
├── internal
│   ├── dao
│   │   ├── auth.go
│   │   ├── dao.go
│   │   └── tag.go
│   ├── middleware
│   │   ├── access_log.go
│   │   ├── app_info.go
│   │   ├── context_timeout.go
│   │   ├── jwt.go
│   │   ├── limiter.go
│   │   ├── recovery.go
│   │   └── translations.go
│   ├── model
│   │   ├── article.go
│   │   ├── article_tag.go
│   │   ├── auth.go
│   │   ├── model.go
│   │   └── tag.go
│   ├── routers
│   │   ├── api
│   │   └── router.go
│   └── service
│       ├── article.go
│       ├── auth.go
│       ├── service.go
│       ├── tag.go
│       └── upload.go
├── main.go
├── pkg
│   ├── app
│   │   ├── app.go
│   │   ├── form.go
│   │   ├── jwt.go
│   │   └── pagination.go
│   ├── convert
│   │   └── convert.go
│   ├── email
│   │   └── email.go
│   ├── errcode
│   │   ├── common_code.go
│   │   ├── errcode.go
│   │   └── module_code.go
│   ├── limiter
│   │   ├── limiter.go
│   │   └── method_limiter.go
│   ├── logger
│   │   └── logger.go
│   ├── setting
│   │   ├── section.go
│   │   └── setting.go
│   ├── upload
│   │   └── file.go
│   ├── util
│   │   └── md5.go
│   └── validator
│       └── custom_validator.go
└── storage
    └── uploads
        ├── 0cc175b9c0f1b6a831c399e269772661.txt
        └── 92eb5ffee6ae2fec3ad71c777531578f.txt

24 directories, 53 files

```

