proto:
	protoc -I=. --go_out=./web ./crud.proto
	protoc -I=. --go_out=./backapp ./crud.proto
