生成api文件
goctl api go -api ./app/user/api/desc/user.api -dir ./app/user/api
goctl api go -api ./app/community/api/desc/community.api -dir ./app/community/api
goctl api go -api ./app/post/api/desc/post.api -dir ./app/post/api
goctl api go -api ./app/post/api/desc/post.api -dir ./tmp/app/post/api

生成rpc文件 进入对应目录
goctl rpc protoc ./app/user/rpc/pb/*.proto --go_out=./tmp/app/user/rpc --go-grpc_out=./tmp/app/user/rpc  --zrpc_out=./tmp/app/user/rpc --style=goZero
goctl rpc protoc ./app/community/rpc/pb/*.proto --go_out=./app/community/rpc --go-grpc_out=./app/community/rpc  --zrpc_out=./app/community/rpc --style=goZero
goctl rpc protoc ./app/post/rpc/pb/*.proto --go_out=./app/post/rpc --go-grpc_out=./app/post/rpc  --zrpc_out=./app/post/rpc --style=goZero


生成model
goctl model mysql ddl -src ./deploy/sql/user.sql -dir ./app/user/model -c
goctl model mysql ddl -src ./deploy/sql/community.sql -dir ./app/community/model -c
goctl model mysql ddl -src ./deploy/sql/post.sql -dir ./app/post/model -c
