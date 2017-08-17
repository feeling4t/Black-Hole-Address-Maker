// makeblackhole project main.go
package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"os"

	"github.com/crowsonkb/base58"
)

func printUsage() {
	fmt.Printf("\tMake a legal bitcoin-like BlackHole address\n")
	fmt.Printf("######## WARN: ANY COIN SEND TO A BLACKHOLE ADDRESS WILL NEVER GET BACK ########")
	fmt.Printf("\nUsage:\n\tmakeblackhole prefix\n")
	fmt.Printf("\tThe prefix string must NOT longer than 28\n")
	fmt.Printf("Example:\n\tmakeblackhole 1bitcoinDestroy\n\tThe address is: 1bitcoinDestroy11111111111111qJkq5")
	fmt.Printf("\n\tmakeblackhole LiteCoinDestroy\n\tThe address is: LiteCoinDestroy11111111111115LGn1H")
	fmt.Printf("\n\tmakeblackhole BLackCoinDestroy\n\tThe address is: BLackCoinDestroy1111111111113WGuzJ")
	fmt.Printf("\n\tmakeblackhole 1bitcoinDestroyXXXXXXXXXXXXX\n\tThe address is: 1bitcoinDestroyXXXXXXXXXXXXWvxdT7n\n")
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		printUsage()
		os.Exit(1)
	}
	prefix := flag.Arg(0)
	if len(prefix) > 28 {
		println("Error: prefix string too long")
		printUsage()
		os.Exit(2)
	}

	address := []byte("1111111111111111111111111111111111")
	copy(address, prefix)

	hex, err := base58.Bitcoin.Decode(string(address))
	if err != nil {
		fmt.Println("Error when decode:", err)
		os.Exit(3)
	}

	sum1 := sha256.Sum256(hex[:len(hex)-4])
	sum2 := sha256.Sum256(sum1[:])
	copy(hex[len(hex)-4:], sum2[0:4])
	fmt.Println("The address is: ", base58.Bitcoin.Encode(hex))
}
