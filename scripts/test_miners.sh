#!/bin/sh

# Test

Dhpcd tx data add-data 0x97e13fccfe2edab7e8077d0f2ac69fe984cfe7f7 dhpc1sr55zk85uerru8tu0m7l7aujha5f8chcxh0073 ethereum any 500 110 --from dataProvider1 -y --log_level error  --gas 500237

Dhpcd tx data add-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 dhpc1sr55zk85uerru8tu0m7l7aujha5f8chcxh0073 ethereum any 500 220 --from dataProvider1 -y --log_level error  --gas 500237 

Dhpcd tx data add-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 dhpc1g6veqgqucfsx987uz9r03q8hesy8798c5e5c7l ethereum any 500 330 --from dataProvider2 -y --log_level error  --gas 500237

Dhpcd tx data add-data 0xc436eb8aed128275c8f224de2f1dd202c0ab5830 dhpc16stssnndzktenrpjpy0sk84pvqh8800u0h80s3 ethereum any 500 440 --from dataProvider3 -y --log_level error  --gas 500237

Dhpcd tx data add-data 0xc8372ff584f621041b6B107050a16c4EcAa57F5A dhpc1g6veqgqucfsx987uz9r03q8hesy8798c5e5c7l ethereum any 500 550 --from dataProvider2 -y --log_level error  --gas 500237


Dhpcd tx request create-treasury `Dhpcd keys show treasury -a` -y --log_level error  --from alice

Dhpcd tx request create-allowed-oracles main `Dhpcd keys show oracle -a` --from alice -y --log_level error 

Dhpcd tx request create-request-record aea7cfd3-1e63-4f90-9311-1c32dacf41cb 0xa0fb4b5e51e55405e5b9e6c5defe381409e99204a3a23b5ef994e3951ded26e0 ethereum `Dhpcd keys show dapp -a` 0xc8372ff584f621041b6B107050a16c4EcAa57F5A `Dhpcd keys show oracle -a` 17446887 --from oracle -y --log_level error  --gas 500237

