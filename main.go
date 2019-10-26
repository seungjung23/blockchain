package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const prefixUser = "user"

var logger = shim.NewLogger("main")

type SmartContract struct {
}

//var _ main.SmartContract = (*main.SmartContractCC)(nil)
var bcFunctions = map[string]func(shim.ChaincodeStubInterface, []string) pb.Response{
	"user_get_info": getUser,
	"user_create":   createUser,
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger := shim.NewLogger("Init")
	logger.Info("chaincode initialized")
	return shim.Success([]byte{})
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	logger.Infof("function name = %s", function)

	if function == "init" {
		return t.Init(stub)
	}
	bcFunc := bcFunctions[function]
	if bcFunc == nil {
		return shim.Error("Invalid invoke function.")
	}
	return bcFunc(stub, args)
}

func main() {
	logger.SetLevel(shim.LogInfo)

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
