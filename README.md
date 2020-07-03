# Xfinite Blockchain Assignment
Merkle proof of a transaction's existence in a given set of ordered blockchain transactions

## Assignment Requirements
1. Given a set of transaction hashes for a block as [ “ab12” -> “bn56” -> “lk87” -> ”op92” ] , write the code for the simple merkle root of the transactions above. The code should have a function generate_merkle_root()

2. Given a transaction hash “lm46” at index 1, prove that it does not belong to the above merkle tree.

## Package reference
Following functions are available in the package:
1. **generate_merkle_root:** Creates the merkle tree and generates the merkle root from a given set of ordered transaction hashes.

2. **prove_membership:** Takes a transaction hash and an index and returns a bool to denote whether of not the transaction exists at the given index. If it exists, the proof is given. If not, the proof path of the actual element at that index is given. 

2. **generate_proof_of_membership:** Takes an included transaction hash and returns its proof path.

## Running
Go is required to run this code.

```shell
$ git clone https://github.com/muditmarda/xfinite-assignment.git
$ go run proof.go
```

### NOTE:
* In Bitcoin, while generating the merkle root, the transaction hashes are converted from little-endian to big-endian before generating the merkle root which after computation is converted back to little-endian and returned. This has not been implemented here and this is just a simple implementation of merkle root generation and proof.
* For merkleization, trees have to be binary balanced, and that the number of transactions to be stored have to be some exponent of two. When that is not the case, the last leaf node is duplicated to make it as such.