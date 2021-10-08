#!bin/bash

lk login admin p4ssw0rd

lk create repo batest

lk create env batest/dev
lk create env batest/stg
lk create env batest/prd

lk create secret batest/dev/POSTGRES_PASSWORD $(lk rand 40)
lk create secret batest/dev/POSTGRES_USERNAME postgres
lk create secret batest/dev/POSTGRES_HOST phost-aws.com
lk create secret batest/dev/POSTGRES_PORT 5432

lk copy batest/dev stg
lk copy batest/dev prd
