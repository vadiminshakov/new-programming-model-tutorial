Demo repo for article https://habr.com/ru/post/494880/

Download docker images
^^^^^^^^^^^^^^^^^^^^^^

We need four docker images in order for "dev mode" to run against the supplied
docker compose script.  If you installed the ``fabric-samples`` repo clone and
followed the instructions to `install samples, binaries and docker images <http://hyperledger-fabric.readthedocs.io/en/latest/install.html>`_, then
you should have the necessary Docker images installed locally.

Terminal 1 - Start the network
------------------------------

.. code:: bash

    cd chaincode && GOPROXY=direct go mod vendor -v && cd ../ && docker-compose -f docker-compose-simple.yaml up

The above starts the network with the ``SingleSampleMSPSolo`` orderer profile and
launches the peer in "dev mode".  It also launches two additional containers -
one for the chaincode environment and a CLI to interact with the chaincode.  The
commands for create and join channel are embedded in the CLI container, so we
can jump immediately to the chaincode calls.

.. note:: TLS is not enabled as it is not supported when running chaincode in dev mode.

Terminal 2 - Build & start the chaincode
----------------------------------------

.. code:: bash

  docker exec -it chaincode sh

You should see the following:

.. code:: sh

  /opt/gopath/src/chaincode $

Now, compile your chaincode:

.. code:: sh

  go build -mod=vendor -o chaincode

Now run the chaincode:

.. code:: sh

  CORE_CHAINCODE_ID_NAME=mycc:0 CORE_PEER_TLS_ENABLED=false ./chaincode -peer.address peer:7052

The chaincode is started with peer and chaincode logs indicating successful registration with the peer.
Note that at this stage the chaincode is not associated with any channel. This is done in subsequent steps
using the ``instantiate`` command.

Terminal 3 - Use the chaincode
------------------------------

Even though you are in ``--peer-chaincodedev`` mode, you still have to install the
chaincode so the life-cycle system chaincode can go through its checks normally.
This requirement may be removed in future when in ``--peer-chaincodedev`` mode.

We'll leverage the CLI container to drive these calls.

.. code:: bash

  docker exec -it cli bash

.. code:: bash

  peer chaincode install -p chaincodedev/chaincode -n mycc -v 0
  peer chaincode instantiate -n mycc -v 0 -c '{"Args":[]}' -C myc


Test chaincode with these commands
------------------------------
  
Create key:

      peer chaincode invoke -n mycc -c '{"Args":["Create", "KEY_1", "VALUE_1"]}' -C myc
  
Update key:

      peer chaincode invoke -n mycc -c '{"Args":["Update", "KEY_1", "VALUE_2"]}' -C myc
  
Read key:

      peer chaincode query -n mycc -c '{"Args":["Read", "KEY_1"]}' -C myc
  
Get metadata:

      peer chaincode query -n mycc -c '{"Args":["org.hyperledger.fabric:GetMetadata"]}' -C myc

Bad request:

      peer chaincode query -n mycc -c '{"Args":["BadRequest", "BadKey"]}' -C myc
