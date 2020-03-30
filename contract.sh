peer chaincode invoke -n mycc -c '{"Args":["Create", "KEY_1", "VALUE_1"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["Update", "KEY_1", "VALUE_2"]}' -C myc
peer chaincode query -n mycc -c '{"Args":["Read", "KEY_1"]}' -C myc