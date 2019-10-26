package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func createUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument count.")
	}

	user := user{}
	err := json.Unmarshal([]byte(args[0]), &user)
	if err != nil {
		return shim.Error(err.Error())
	}

	key, err := stub.CreateCompositeKey(prefixUser, []string{user.UserName})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Check if the user already exists
	userAsBytes, _ := stub.GetState(key)
	// User does not exist, attempting creation
	if len(userAsBytes) == 0 {
		userAsBytes, err = json.Marshal(user)
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(key, userAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

		// Return nil, if user is newly created
		return shim.Success(nil)
	}

	err = json.Unmarshal(userAsBytes, &user)
	if err != nil {
		return shim.Error(err.Error())
	}

	userResponse := struct {
		UserId   string `json:"userid"`
		UserName string `json:"username"`
	}{
		UserId:   user.UserId,
		UserName: user.UserName,
	}

	userResponseAsBytes, err := json.Marshal(userResponse)
	if err != nil {
		return shim.Error(err.Error())
	}
	// Return the username and the password of the already existing user
	return shim.Success(userResponseAsBytes)
}

func getUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument count.")
	}

	input := struct {
		UserId   string `json:"userid"`
		UserName string `json:"username"`
	}{}

	err := json.Unmarshal([]byte(args[0]), &input)
	if err != nil {
		return shim.Error(err.Error())
	}

	userKey, err := stub.CreateCompositeKey(prefixUser, []string{input.UserName})
	if err != nil {
		return shim.Error(err.Error())
	}
	userBytes, _ := stub.GetState(userKey)
	if len(userBytes) == 0 {
		return shim.Success(nil)
	}

	response := struct {
		UserId   string `json:"userid"`
		UserName string `json:"username"`
	}{}
	err = json.Unmarshal(userBytes, &response)
	if err != nil {
		return shim.Error(err.Error())
	}
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(responseBytes)
}
