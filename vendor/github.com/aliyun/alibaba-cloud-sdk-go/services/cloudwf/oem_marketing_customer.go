package cloudwf

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// OemMarketingCustomer invokes the cloudwf.OemMarketingCustomer API synchronously
// api document: https://help.aliyun.com/api/cloudwf/oemmarketingcustomer.html
func (client *Client) OemMarketingCustomer(request *OemMarketingCustomerRequest) (response *OemMarketingCustomerResponse, err error) {
	response = CreateOemMarketingCustomerResponse()
	err = client.DoAction(request, response)
	return
}

// OemMarketingCustomerWithChan invokes the cloudwf.OemMarketingCustomer API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/oemmarketingcustomer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OemMarketingCustomerWithChan(request *OemMarketingCustomerRequest) (<-chan *OemMarketingCustomerResponse, <-chan error) {
	responseChan := make(chan *OemMarketingCustomerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OemMarketingCustomer(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// OemMarketingCustomerWithCallback invokes the cloudwf.OemMarketingCustomer API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/oemmarketingcustomer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OemMarketingCustomerWithCallback(request *OemMarketingCustomerRequest, callback func(response *OemMarketingCustomerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OemMarketingCustomerResponse
		var err error
		defer close(result)
		response, err = client.OemMarketingCustomer(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// OemMarketingCustomerRequest is the request struct for api OemMarketingCustomer
type OemMarketingCustomerRequest struct {
	*requests.RpcRequest
	Bid requests.Integer `position:"Query" name:"Bid"`
}

// OemMarketingCustomerResponse is the response struct for api OemMarketingCustomer
type OemMarketingCustomerResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateOemMarketingCustomerRequest creates a request to invoke OemMarketingCustomer API
func CreateOemMarketingCustomerRequest() (request *OemMarketingCustomerRequest) {
	request = &OemMarketingCustomerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "OemMarketingCustomer", "cloudwf", "openAPI")
	return
}

// CreateOemMarketingCustomerResponse creates a response to parse from OemMarketingCustomer response
func CreateOemMarketingCustomerResponse() (response *OemMarketingCustomerResponse) {
	response = &OemMarketingCustomerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}