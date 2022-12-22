
#!/bin/bash
protoc --proto_path=../schema/ --go_out=plugins=grpc:../controller/pb ../schema/*.proto
