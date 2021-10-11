#!bin/bash

lk login admin p4ssw0rd

lk create repo abc

lk create env abc/dev
lk create env abc/stg
lk create env abc/prd

lk create secret abc/dev/POSTGRES_PASSWORD $(lk rand 40)
lk create secret abc/dev/POSTGRES_USERNAME postgres
lk create secret abc/dev/POSTGRES_HOST phost-aws.com
lk create secret abc/dev/POSTGRES_PORT 5432

lk copy env abc/dev stg
lk copy env abc/dev prd
