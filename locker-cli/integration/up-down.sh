#!bin/bash

lk create repo 1 2 3 4 5
lk create env 1/dev stg prd
lk create env 2/dev stg prd
lk create env 3/dev stg prd
lk create env 4/dev stg prd
lk create env 5/dev stg prd

lk create secret 1/dev/P1 $(lk rand 40)
lk create secret 1/stg/P1 $(lk rand 40)
lk create secret 1/prd/P1 $(lk rand 40)

lk create secret 2/dev/P1 $(lk rand 40)
lk create secret 2/stg/P1 $(lk rand 40)
lk create secret 2/prd/P1 $(lk rand 40)

lk create secret 3/dev/P1 $(lk rand 40)
lk create secret 3/stg/P1 $(lk rand 40)
lk create secret 3/prd/P1 $(lk rand 40)

lk create secret 4/dev/P1 $(lk rand 40)
lk create secret 4/stg/P1 $(lk rand 40)
lk create secret 4/prd/P1 $(lk rand 40)

lk create secret 5/dev/P1 $(lk rand 40)
lk create secret 5/stg/P1 $(lk rand 40)
lk create secret 5/prd/P1 $(lk rand 40)

lk delete repo --force 1 2 3 4 5
