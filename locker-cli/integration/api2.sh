#!bin/bash

lk create repo aaa
lk create env aaa/dev aaa/stg aaa/prd

lk create secret aaa/dev/PG_PW1 $(lk rand 40)
lk create secret aaa/dev/PG_PW2 $(lk rand 40)
lk create secret aaa/dev/PG_PW3 $(lk rand 40)

lk copy env aaa/dev stg prd

lk get aaa
lk get aaa/dev
lk get aaa/stg
lk get aaa/prd

lk delete repo aaa --force

lk get aaa
