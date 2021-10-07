lk login admin p4ssw0rd

lk create repo book-api

lk create env  book-api/dev
lk create env  book-api/stg
lk create env  book-api/prd

lk create secret book-api/dev/POSTGRES_PASSWORD $(lk rand 40)
lk create secret book-api/stg/POSTGRES_PASSWORD $(lk rand 40)
lk create secret book-api/prd/POSTGRES_PASSWORD $(lk rand 40)
