docker-compose -f docker-compose-simple.yaml up -d
docker exec -e CORE_CHAINCODE_ID_NAME=mycc:0 -e CORE_PEER_TLS_ENABLED=false -e GOPROXY="https://goproxy.io" chaincode sh -c "go mod vendor -v && go build -v -mod=vendor -o chaincode && chmod 777 ./chaincode"
docker exec -e CORE_CHAINCODE_ID_NAME=mycc:0 -e CORE_PEER_TLS_ENABLED=false chaincode sh -c "./chaincode -peer.address peer:7052"
sleep 5
docker exec cli sh -c "peer chaincode install -p /opt/gopath/src/chaincodedev -n mycc -v 0"
docker exec -e ARG='{"Args":[]}' cli sh -c "peer chaincode instantiate -n mycc -v 0 -c '$ARG' -C myc"