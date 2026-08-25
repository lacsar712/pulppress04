基于 Go 实现的纸浆压榨线调度 Web 项目，一款后端服务，处理压区压力校验、出站指令与浆线批次落库。

# PulpPress Nip

纸浆压榨线 / 压区压力批次调度投递

Read `PROJECT.md` for the product scope, who uses it, and what the records mean. This is not a generic CMS.

## Run

```bash
set GOTOOLCHAIN=local
set CGO_ENABLED=0
go test ./... -count=1
go run ./cmd/server -addr 127.0.0.1:8080 -db ./data.sqlite
go run ./cmd/seed -db ./data.sqlite
```

Open http://127.0.0.1:8080/ for the console.
