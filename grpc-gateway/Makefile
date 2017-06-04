gen_proto:
	protoc --go_out=${GOPATH}/src proto/service.proto

gen_proto2:
	protoc -I. -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:${GOPATH}/src proto/service.proto --grpc-gateway_out=logtostderr=true:.

gen_server_tls:
	openssl req -x509 -newkey rsa:4096 -keyout server/server-key.pem -out server/server-cert.pem -days 365 -nodes -subj '/CN=localhost'

run/server:
	go run server/main.go

run/client:
	go run client/main.go