生成api文件
goctl api go -api ./app/user/api/desc/user.api -dir ./tmp/app/user/api
goctl api go -api ./app/community/api/desc/community.api -dir ./app/community/api
goctl api go -api ./tmp/app/post/api/desc/post.api -dir ./tmp/app/post/api
goctl api go -api ./app/post/api/desc/post.api -dir ./tmp/app/post/api
goctl api go -api ./app/comment/api/desc/comment.api -dir ./app/comment/api
goctl api go -api ./app/vote/api/desc/vote.api -dir ./app/vote/api

生成rpc文件 进入对应目录
goctl rpc protoc ./app/user/rpc/pb/*.proto --go_out=./tmp/app/user/rpc --go-grpc_out=./tmp/app/user/rpc  --zrpc_out=./tmp/app/user/rpc --style=goZero
goctl rpc protoc ./app/user/rpc/pb/*.proto --go_out=./tmp/app/user/rpc --go-grpc_out=./tmp/app/user/rpc  --zrpc_out=./tmp/app/user/rpc --style=goZero
goctl rpc protoc ./app/community/rpc/pb/*.proto --go_out=./app/community/rpc --go-grpc_out=./app/community/rpc  --zrpc_out=./app/community/rpc --style=goZero
goctl rpc protoc ./app/post/rpc/pb/*.proto --go_out=./tmp/app/post/rpc --go-grpc_out=./tmp/app/post/rpc  --zrpc_out=./tmp/app/post/rpc --style=goZero
goctl rpc protoc ./app/post/rpc/pb/*.proto --go_out=./tmp/app/post/rpc --go-grpc_out=./tmp/app/post/rpc  --zrpc_out=./tmp/app/post/rpc --style=goZero
goctl rpc protoc ./app/comment/rpc/pb/*.proto --go_out=./app/comment/rpc --go-grpc_out=./app/comment/rpc  --zrpc_out=./app/comment/rpc --style=goZero
goctl rpc protoc ./app/comment/rpc/pb/*.proto --go_out=./tmp/app/comment/rpc --go-grpc_out=./tmp/app/comment/rpc  --zrpc_out=./tmp/app/comment/rpc --style=goZero
goctl rpc protoc ./app/vote/rpc/pb/*.proto --go_out=./tmp/app/vote/rpc --go-grpc_out=./tmp/app/vote/rpc  --zrpc_out=./tmp/app/vote/rpc --style=goZero


生成model
goctl model mysql ddl -src ./deploy/sql/user.sql -dir ./app/user/model -c
goctl model mysql ddl -src ./deploy/sql/community.sql -dir ./tmp/app/community/model -c
goctl model mysql ddl -src ./deploy/sql/post.sql -dir ./tmp/app/post/model -c
goctl model mysql ddl -src ./deploy/sql/comment.sql -dir ./app/comment/model -c
goctl model mysql ddl -src ./deploy/sql/vote.sql -dir ./app/vote/model -c

goctl model mysql ddl -src ./deploy/sql/post.sql -dir ./tmp/app/post/model -c
