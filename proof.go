package main

import (
	"fmt"
	"bufio"
  	"os"
	"math"
	"strings"
	"strconv"
	"crypto/sha256"
)

var tree = make([][]string, 0)

func main() {
	fmt.Println("\n==================================================================================================================================================")
	fmt.Println("=========================================================== Program to return the ================================================================")
	fmt.Println("================================================== Proof of Membership in a Merkle Tree ==========================================================")
	fmt.Println("==================================================================================================================================================")
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Println("\nInput the set of transaction hashes for a block (in the format `[tx1 -> tx2 -> tx2]`): ")
		scanner.Scan()
		transactions := parse_transactions_input(scanner.Text())
		if len(transactions) == 0 {
			fmt.Println("Invalid set of transaction hashes entered, try again.")
			continue
		}
		generate_merkle_root(transactions)

		fmt.Println("\nInput the transaction hash (without quotes) of the transaction for which the proof is needed: ")
		scanner.Scan()
		transaction_to_find := scanner.Text()
		fmt.Println("Input the index of this transaction (starting from 0): ")
		scanner.Scan()
		claimed_index, err := strconv.Atoi(scanner.Text())
		if claimed_index >= len(transactions) || err != nil {
			fmt.Println("\nInvalid index entered, try again.")
			continue
		}

		proof, exists := prove_membership(transaction_to_find, claimed_index)
		if exists {
			fmt.Println("\n********************************  Voila,", transaction_to_find, "exists!! *********************************")
			fmt.Println("Here's the proof of it's membership in the merkle tree generated from the set of the given transaction hashes:\n", proof)
		} else {
			fmt.Println("\nSorry,", transaction_to_find, "does not exist.\nHere's the proof that it does not belong to the merkle tree generated from the set of the given transaction hashes:\n", proof)
		}
		fmt.Println("\n\n====================================================================================================================================================\n")
  	}
}

func parse_transactions_input(input_str string) (transactions []string) {
	replacer := strings.NewReplacer(" ", "", "[", "", "]", "", "\"", "", "“", "", "”", "", "'", "", "`", "")
	parsed_input_str := replacer.Replace(input_str)
	if len(parsed_input_str) == 0 {
		return []string{}
	}
 	return strings.Split(parsed_input_str, "->")
}

func generate_merkle_root(txs []string) {
	size := len(txs)	

	leaf_node_count := math.Log2(float64(size))
	for leaf_node_count != float64(int64(leaf_node_count)) {
		txs = append(txs, txs[size-1])
		size += 1
		leaf_node_count = math.Log2(float64(size))
	}
	
	height := 0
	for size >= 1 {
		size /= 2
		height += 1
	}

    res := make([][]string, height)
	res[0] = txs
	for i := 1; i < height ; i++ {
		for j := 0; j < len(res[i-1]) ; j+=2 {
			left := res[i-1][j]
			right := res[i-1][j+1]
			parent := fmt.Sprintf("%x", sha256.Sum256([]byte(left + right)))
			res[i] = append(res[i], parent)
		}
	}
	tree = res
}


func give_proof_of_membership(tx string) ([]string) {
	proof := make([]string, 0)
	search_str := tx
	for i := 0; i < (len(tree) - 1); i++ {
		for j := 0; j < len(tree[i]); j++ {
			current_element := tree[i][j]
			if search_str == current_element {
				// if index is odd, return path to element at index-1 (left_neighbour)
				// if index is even, return path to element at index+1 (right_neighbour)
				if j%2 == 1 {
					left_neighbour := tree[i][j-1]

					proof = append(proof, left_neighbour)
					search_str = fmt.Sprintf("%x", sha256.Sum256([]byte(left_neighbour + current_element)))
				} else {
					right_neighbour := tree[i][j+1]

					proof = append(proof, right_neighbour)
					search_str = fmt.Sprintf("%x", sha256.Sum256([]byte(current_element + right_neighbour)))
				}
			}
		}
	}
	proof = append(proof, tree[len(tree)-1][0])
	return proof
}

func prove_membership(tx string, index int) (proof []string, exists bool) {
	if tree[0][index] != tx {
		// if index is odd, return proof of membership of element at index-1
		// if index is even, return proof of membership of element at index+1
		if index%2 == 1 {
			return give_proof_of_membership(tree[0][index-1]), false
		} else {
			return give_proof_of_membership(tree[0][index+1]), false
		}
	} 
	return give_proof_of_membership(tree[0][index]), true	
}