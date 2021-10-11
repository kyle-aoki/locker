lk rename repo abc xyz

lk rename env xyz/dev DEV
lk rename env xyz/stg STG
lk rename env xyz/prd PRD

lk create secret xyz/DEV/SEC $(lk rand 40)
lk rename secret xyz/DEV/SEC CESER
