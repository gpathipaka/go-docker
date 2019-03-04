#build#
#protobuf 
protoc --go_out=plugins=micro:. proto/vessel/vessel.proto

#Docker Build
docker build -t vessel-service .

#docker Run
docker run -p 50052:50052 -e MICRO_SERVICE_ADDRESS="50051 -e MICRO_REGISTRY=mdns vessel-service