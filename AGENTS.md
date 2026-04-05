# AGENTS.md

## Project Overview
- **Type**: Go blog backend (Gin + GORM + Viper)
- **Entry**: `main.go`
- **Run**: `go run main.go`

## Critical Init Order
`main.go:9-17` - config must be loaded before database:
```
config.InitConfig() → config.InitDB() → router.SetupRouter()
```

## Config
- Location: `config/app.yaml`
- Loaded via Viper with `mapstructure` tags
- Database credentials are in the YAML file

## Directory Structure
```
config/    - config loading and DB init
controller/ - HTTP handlers
model/     - GORM models (User, Post, Comment)
middleware/ - auth, cors, log
util/      - JWT, response helpers
router/    - Gin route setup
```

## Dependencies
- gin v1.9.1
- gorm v1.31.1
- viper v1.18.2

## Build/Run
- No Makefile; use `go run main.go`
- No tests in this repo
- No lint/typecheck commands configured