#!bin/bash

lk login admin p4ssw0rd

lk create repo   abc
lk create env    abc/xyz
lk create secret abc/xyz/password $(lk rand 40)

lk delete env abc/xyz --force

lk get abc/xyz
