package airstack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	client, _ := AirstackClient("API_KEY")
	assert.NotNil(t, client)
}

func TestSendRequest(t *testing.T) {

	client, _ := AirstackClient("API_KEY")
	assert.NotNil(t, client)

	testQuery := `
	query MyQuery($name1: Address!, $name2: Address!) {
		TokenBalances(
		  input: {filter: {tokenAddress: {_eq: $name1}}, blockchain: ethereum}
		) {
		  TokenBalance {
			chainId
			blockchain
			amount
			lastUpdatedTimestamp
			id
		  }
		}
		test2: TokenTransfers(
		  input: {filter: {tokenAddress: {_eq: $name2}}, blockchain: ethereum}
		) {
		  TokenTransfer {
			chainId
			blockchain
			id
			amount
		  }
		}
	  }
	`
	variables := map[string]interface{}{
        "name1": "0x1130547436810db920fa73681c946fea15e9b758",
        "name2": "0xf4eced2f682ce333f96f2d8966c613ded8fc95dd",
    }

	response, err := client.executeQuery(testQuery, variables)
	fmt.Println(response)
	assert.Nil(t, err)
}
