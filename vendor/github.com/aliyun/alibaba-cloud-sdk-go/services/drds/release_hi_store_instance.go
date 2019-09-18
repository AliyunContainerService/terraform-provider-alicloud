package drds

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

// ReleaseHiStoreInstance invokes the drds.ReleaseHiStoreInstance API synchronously
// api document: https://help.aliyun.com/api/drds/releasehistoreinstance.html
func (client *Client) ReleaseHiStoreInstance(request *ReleaseHiStoreInstanceRequest) (response *ReleaseHiStoreInstanceResponse, err error) {
	response = CreateReleaseHiStoreInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseHiStoreInstanceWithChan invokes the drds.ReleaseHiStoreInstance API asynchronously
// api document: https://help.aliyun.com/api/drds/releasehistoreinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseHiStoreInstanceWithChan(request *ReleaseHiStoreInstanceRequest) (<-chan *ReleaseHiStoreInstanceResponse, <-chan error) {
	responseChan := make(chan *ReleaseHiStoreInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseHiStoreInstance(request)
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

// ReleaseHiStoreInstanceWithCallback invokes the drds.ReleaseHiStoreInstance API asynchronously
// api document: https://help.aliyun.com/api/drds/releasehistoreinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseHiStoreInstanceWithCallback(request *ReleaseHiStoreInstanceRequest, callback func(response *ReleaseHiStoreInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseHiStoreInstanceResponse
		var err error
		defer close(result)
		response, err = client.ReleaseHiStoreInstance(request)
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

// ReleaseHiStoreInstanceRequest is the request struct for api ReleaseHiStoreInstance
type ReleaseHiStoreInstanceRequest struct {
	*requests.RpcRequest
	HistoreInstanceId string `position:"Query" name:"HistoreInstanceId"`
	DrdsPassword      string `position:"Query" name:"DrdsPassword"`
	DrdsInstanceId    string `position:"Query" name:"DrdsInstanceId"`
}

// ReleaseHiStoreInstanceResponse is the response struct for api ReleaseHiStoreInstance
type ReleaseHiStoreInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateReleaseHiStoreInstanceRequest creates a request to invoke ReleaseHiStoreInstance API
func CreateReleaseHiStoreInstanceRequest() (request *ReleaseHiStoreInstanceRequest) {
	request = &ReleaseHiStoreInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "ReleaseHiStoreInstance", "drds", "openAPI")
	return
}

// CreateReleaseHiStoreInstanceResponse creates a response to parse from ReleaseHiStoreInstance response
func CreateReleaseHiStoreInstanceResponse() (response *ReleaseHiStoreInstanceResponse) {
	response = &ReleaseHiStoreInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}