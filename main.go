package main

import (
	testAPI "github.com/ameight/testserv/pkg/api"
)

const (
	first  = 1
	second = 2
)

func main() {

}

func foo(api *testAPI.GetInfoRequest) testAPI.GetInfoResponse {
	api.GetNumber() 

	return testAPI.GetInfoResponse{
		Number1: api.GetNumber()+1,
		Number2: ,
	}	
}