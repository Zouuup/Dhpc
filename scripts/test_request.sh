#!/bin/sh

Decentd tx data add-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 decent1sr55zk85uerru8tu0m7l7aujha5f8chcxh0073 ethereum any 150 --from dataProvider1 -y --gas 500237 
Decentd tx data add-data 0xc236eb8aed128275c8f224de2f1dd202c0ab5830 decent1sr55zk85uerru8tu0m7l7aujha5f8chcxh0073 ethereum any 250 --from dataProvider1 -y --gas 500237
Decentd tx data add-data 0xc136eb8aed128275c8f224de2f1dd202c0ab5830 decent1sr55zk85uerru8tu0m7l7aujha5f8chcxh0073 ethereum any 350 --from dataProvider1 -y --gas 500237
Decentd tx data add-data 0xc536eb8aed128275c8f224de2f1dd202c0ab5830 decent1sr55zk85uerru8tu0m7l7aujha5f8chcxh0073 ethereum any 450 --from dataProvider1 -y --gas 500237

Decentd tx data add-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 decent1g6veqgqucfsx987uz9r03q8hesy8798c5e5c7l ethereum any 150 --from dataProvider2 -y --gas 500237
Decentd tx data add-data 0xc236eb8aed128275c8f224de2f1dd202c0ab5830 decent1g6veqgqucfsx987uz9r03q8hesy8798c5e5c7l ethereum any 250 --from dataProvider2 -y --gas 500237
Decentd tx data add-data 0xc136eb8aed128275c8f224de2f1dd202c0ab5830 decent1g6veqgqucfsx987uz9r03q8hesy8798c5e5c7l ethereum any 350 --from dataProvider2 -y --gas 500237
Decentd tx data add-data 0xc536eb8aed128275c8f224de2f1dd202c0ab5830 decent1g6veqgqucfsx987uz9r03q8hesy8798c5e5c7l ethereum any 450 --from dataProvider2 -y --gas 500237

Decentd tx data add-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 decent16stssnndzktenrpjpy0sk84pvqh8800u0h80s3 ethereum any 150 --from dataProvider3 -y --gas 500237
Decentd tx data add-data 0xc236eb8aed128275c8f224de2f1dd202c0ab5830 decent16stssnndzktenrpjpy0sk84pvqh8800u0h80s3 ethereum any 250 --from dataProvider3 -y --gas 500237
Decentd tx data add-data 0xc136eb8aed128275c8f224de2f1dd202c0ab5830 decent16stssnndzktenrpjpy0sk84pvqh8800u0h80s3 ethereum any 350 --from dataProvider3 -y --gas 500237
Decentd tx data add-data 0xc536eb8aed128275c8f224de2f1dd202c0ab5830 decent16stssnndzktenrpjpy0sk84pvqh8800u0h80s3 ethereum any 450 --from dataProvider3 -y --gas 500237

Decentd tx request create-treasury `Decentd keys show treasury -a` -y --from alice

provider1Key=`Decentd keys show dataProvider1 -a`
provider2Key=`Decentd keys show dataProvider2 -a`
provider3Key=`Decentd keys show dataProvider3 -a`

rightAddress=`Decentd q data show-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 $provider1Key ethereum | grep -i hash | sed 's/.*: //'`,`Decentd q data show-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 $provider2Key ethereum | grep -i hash | sed 's/.*: //'`,`Decentd q data show-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 $provider3Key ethereum | grep -i hash | sed 's/.*: //'`
oneExtra=`Decentd q data show-data 0xc536eb8aed128275c8f224de2f1dd202c0ab5830 $provider3Key ethereum | grep -i hash | sed 's/.*: //'`
oneNotFound=25ac0fb34dce2b232dc95fc204075af9
rightPlusExtra=$rightAddress,$oneExtra
rightPlusNotFound=$rightAddress,$oneNotFound


Decentd tx request create-request-record aea7cfd3-1e63-4f90-9311-1c32dacf41cb 0xa0fb4b5e51e55405e5b9e6c5defe381409e99204a3a23b5ef994e3951ded26e0 ethereum `Decentd keys show dapp -a` 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 `Decentd keys show oracle -a` 12345678 --from oracle -y --gas 500237

Decentd tx request create-miner-response BobRandomUUID aea7cfd3-1e63-4f90-9311-1c32dacf41cb f547df54e746a70c78056a40fda9f019 $rightAddress --from bob -y --gas 500237

Decentd tx request create-miner-response AliceRandomUUID aea7cfd3-1e63-4f90-9311-1c32dacf41cb 2348653802f49d37e64b9635ab61261c $rightPlusExtra --from alice -y --gas 500237

Decentd tx request create-miner-response Arthurrandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 416568850c9075e2b5656f97131d3e82 $rightPlusNotFound --from arthur -y --gas 500237

Decentd tx request create-miner-response Carlrandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 8b2b19bcd23502ff3a8177ef3ae4b8aa $rightAddress --from carl -y --gas 500237

Decentd tx request create-miner-response Issacrandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 8b2b19bcd23502ff3a8177ef3ae4b8ab $rightAddress --from issac -y --gas 500237

Decentd tx request create-miner-response Stanrandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 8b2b19bcd23502ff3a8177ef3ae4b8ac $rightAddress --from stan -y --gas 500237



#echo "Updating answers"
Decentd tx request update-miner-response BobRandomUUID aea7cfd3-1e63-4f90-9311-1c32dacf41cb 123 43578934 --from bob -y --gas 500237

Decentd tx request update-miner-response Alicerandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 123 65765765 --from alice -y --gas 500237

Decentd tx request update-miner-response Arthurrandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 123 44298392 --from arthur -y --gas 500237

Decentd tx request update-miner-response Carlrandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 5 65765893 --from carl -y --gas 500237

Decentd tx request update-miner-response Issacrandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 5 65765893 --from issac -y --gas 500237

Decentd tx request update-miner-response Stanrandomuuid aea7cfd3-1e63-4f90-9311-1c32dacf41cb 5 65765893 --from stan -y --gas 500237


# Add treasury address similar to how we added oracles 
