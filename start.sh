export ARG='{"Args":[]}'
docker-compose -f docker-compose-simple.yaml up -d
cd chaincode && GOPROXY=direct go mod vendor -v
sudo chmod -R +0777 ../chaincode
docker exec chaincode sh -c "GO111MODULE=on go build -v -mod=vendor -o chaincode"
docker exec chaincode ls
docker exec -e CORE_CHAINCODE_ID_NAME=mycc:0 -e CORE_PEER_TLS_ENABLED=false chaincode sh -c "./chaincode -peer.address peer:7052"
sleep 5
docker exec cli sh -c "peer chaincode install -p /opt/gopath/src/chaincodedev/chaincode -n mycc -v 0"
docker exec cli sh -c "peer chaincode instantiate -n mycc -v 0 -c '$(ARG)' -C myc"