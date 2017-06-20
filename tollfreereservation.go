/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// NumberPortabilityChaincode is a Smart Contract between CSPs for porting In or Porting Out Customers and settling the billing across the CSPs


type NumberPortabilityChaincode struct {
}


type Reserve struct {
	TollFreeno string
	ServiceProvider string
	status string
	AssignedDate string
}
// Init method will be called during deployment.

func (t *NumberPortabilityChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init Chaincode...")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	fmt.Println("Init Chaincode...done")

	return nil, nil
}






// args should be Number, serviceProviderOld, serviceProviderNew

func (t *NumberPortabilityChaincode) RegulatorQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3 arguments")
    }

    key = args[0]+args[1]+args[2]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query NumberPortability Chaincode... end") 
    return valAsbytes, nil 

}


// args should be Number, serviceProviderOld, serviceProviderNew

func (t *NumberPortabilityChaincode) EntitlementFromRecipientCSPQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 3 {
        return nil, errors.New("Incorrect number of arguments. Expecting 3 arguments")
    }

    key = args[0]+args[1]+args[2]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query EntitlementFromRecipientCSPQuery ... end") 
    return valAsbytes, nil 

}



// Invoke Function

func (t *NumberPortabilityChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
      
	 fmt.Println("Invoke NumberPortability Chaincode... start") 

	
	// Handle different functions UserAcceptance
	if function == "Reserve" {
		return t.Reserve(stub, args)
	}else{
	    return nil, errors.New("Invalid function name. Expecting 'EligibilityConfirm' or 'UsageDetailsFromDonorCSP' or 'EntitlementFromRecipientCSP' but found '" + function + "'")
	}
	
	
	fmt.Println("Invoke Numberportability Chaincode... end") 
	
	return nil,nil;
}




// Query to get CSP Service Details

func (t *NumberPortabilityChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query NumberPortability Chaincode... start") 

	
	if function == "EntitlementFromRecipientCSPQuery" {
		return t.EntitlementFromRecipientCSPQuery(stub, args)
	} 
	
	if function == "RegulatorQuery" {
		return t.RegulatorQuery(stub, args)
	} 
	
	if function == "RegulatorQuery1" {
		return t.RegulatorQuery1(stub, args)
	} 
	
	// else We can query WorldState to fetch value
	
	var key, jsonResp string
    var err error

    if len(args) < 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }
	fmt.Println(len(args))
	if len(args) == 3 {
	   key = args[0]+args[1]+args[2]
	} else if len(args) == 2 {
	   key = args[0]+args[1]
	} else {
	   key = args[0]
	}

    
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query NumberPoratbility Chaincode... end") 
    return valAsbytes, nil 
  
	
}

//Reserve Invoke function
func (t *NumberPortabilityChaincode) Reserve(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    fmt.Println("Reserve Information invoke Begins...")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	// Check the Reseve paramater, if true then update world state with new status
        var status1 string
	var key string
	var Acceptance string
	//var value string
	key = args[0]
	//value = args[1]
	Acceptance = args[2]
	if(Acceptance == "true"){
	 status1 = "RequestInitiated"
	} 
	
	ReserveObj := Reserve{TollFreeno: args[0],ServiceProvider: args[1], status: status1, AssignedDate: args[3]}
   fmt.Println("Reserve Details Structure ",ReserveObj)
	//value, e := json.Marshal(ReserveObj)
	//if e != nil {
		//return nil, e
	//}
	//err := stub.PutState(key,[]byte(value)
	//err := stub.PutState(key,[]byte(fmt.Sprintf("%s",value)))
	err := stub.PutState(key,[]byte(fmt.Sprintf("%s",ReserveObj)))
	if err != nil {
		return nil, err
	}
	
	fmt.Println("Reserve Information invoke ends...")
	return nil, nil
}

func (t *NumberPortabilityChaincode) RegulatorQuery1(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2 arguments")
    }

    key = args[0]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    } else if len(valAsbytes) == 0{
	    jsonResp = "{\"Error\":\"Failed to get Query for " + key + "\"}"
        return nil, errors.New(jsonResp)
	}

	fmt.Println("Query NumberPortability Chaincode... end") 
    return valAsbytes, nil 

}




func main() {
	err := shim.Start(new(NumberPortabilityChaincode))
	if err != nil {
		fmt.Println("Error starting NumberPortabilityChaincode: %s", err)
	}
}
