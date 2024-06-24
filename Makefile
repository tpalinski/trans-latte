proto:
	protoc -I=. --go_out=./web ./crud.proto
	protoc -I=. --go_out=./backapp ./crud.proto
	protoc -I=. --python_out=./pricing ./crud.proto
