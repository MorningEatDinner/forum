生成api文件
goctl api go -api ./app/user/api/desc/user.api -dir ./app/user/api

生成rpc文件 进入对应目录
goctl rpc protoc ./app/user/rpc/pb/*.proto --go_out=./tmp/app/user/rpc --go-grpc_out=./tmp/app/user/rpc  --zrpc_out=./tmp/app/user/rpc --style=goZero


生成model
goctl model mysql ddl -src ./deploy/sql/user.sql -dir ./app/user/model -c
