# Prerequisite
1. Clone fabric-sample https://hyperledger-fabric.readthedocs.io/en/release-1.2/install.html
    ```bash
      curl -sSL http://bit.ly/2ysbOFE | bash -s 1.2.0
    ```
# Testing in dev mode
1. Copy hello-chaincode to `fabric-samples/chaincode`
1. Leveraging dev mode
    ```bash
      cd $FABRIC_SAMPLE/chaincode-docker-devmode/chaincode-docker-devmode
      docker-compose -f docker-compose-simple.yaml up
    ```
1. Compile chain code
1. Run peer
    ```bash
      CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./hello-chaincode
    ```

# Install chaincode
1. Leverage CLI container
    ```bash
      docker exec -it cli bash
    ```
    ``` bash
      peer chaincode install -p chaincodedev/chaincode/hello-chaincode -n mycc -v 0
      peer chaincode instantiate -n mycc -v 0 -c '{"Args":["a","10"]}' -C myc
    ```

# Using the app
1. Set
    ```bash
      peer hello-chaincode invoke -n mycc -c '{"Args":["set", "a", "20"]}' -C myc
    ```
1. Get
    ```bash
      peer hello-chaincode query -n mycc -c '{"Args":["query","a"]}' -C myc

    ```