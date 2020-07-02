# xfinite-assignment
Merkle proof of a transaction's existence in a given set of ordered blockchain transactions

## Assignment Requirements:
1. Given a set of transaction hashes for a block as [ “ab12” -> “bn56” -> “lk87” -> ”op92” ] , write the code for the simple merkle root of the transactions above. The code should have a function generate_merkle_root()

2. Given a transaction hash “lm46” at index 1, prove that it does not belong to the above merkle tree.

## Running

Go is required to run this code.

```shell
$ git clone https://github.com/muditmarda/xfinite-assignment.git
$ go run proof.go
```
