package main_test

import (
	"testing"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/stretchr/testify/assert"
)

//var SmartContract Chaincode

// const (
// 	inputJSON = `{"UserId":"seungjung23", "UserName":"SeungJung"}`
// 	userJSON  = "[" + inputJSON + "]"
// )

func TestInit(t *testing.T) { //Init 체크
	stub := shim.NewMockStub("SmartContract", new(SmartContract))
	if !assert.NotNil(t, stub) {
		res := stub.MockInit(util.GenerateUUID(), nil)
		assert.True(t, res.Status < shim.ERRORTHRESHOLD)
	}
}

// func TestCreate(t *testing.T) {
// 	stub := shim.NewMockStub("test", new(SmartContract))
// 	if !assert.NotNil(t, stub) { //인스턴스화 체크
// 		return
// 	}

// 	if !assert.True(
// 		t,
// 		stub.MockInvoke(
// 			util.GenerateUUID(),
// 			getBytes("createUser", inputJSON),
// 		).Status < shim.ERRORTHRESHOLD,
// 	) {
// 		return
// 	}

// 	res := stub.MockInvoke(util.GenerateUUID(), getBytes("getUser"))

// 	_ = assert.True(t, res.Status < shim.ERRORTHRESHOLD) && assert.JSONEq(t, userJSON, string(res.Payload))
// }

// func getBytes(function string, args ...string) [][]byte {
// 	bytes := make([][]byte, 0, len(args)+1)
// 	bytes = append(bytes, []byte(function))
// 	for _, s := range args {
// 		bytes = append(bytes, []byte(s))
// 	}
// 	return bytes
// }
