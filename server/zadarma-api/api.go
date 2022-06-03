package zadarma_api

import (
	"net/http"
	"net/url"

	"encoding/json"
	"log"
)

type apiClient struct {
	Key    string
	Secret string
}

const ApiUrl = "https://api.zadarma.com/"

func (api apiClient) CallMethod(methodName string, params url.Values, methodType string) interface{} {
	client := http.Client{}

	fullURL := BuildAPIUrl(methodName, params)
	sign := Sign(api, methodName, params)
	request, err := http.NewRequest(methodType, fullURL, nil)
	if err != nil {
		log.Printf("Error when creating newRequest for \"%s\": %s", fullURL, err.Error())
		return nil
	}

	request.Header.Set("Authorization", api.Key+":"+sign)

	response, requesting_err := client.Do(request)
	if requesting_err != nil {
		log.Printf("Error when requesting %s: %s", fullURL, requesting_err.Error())
		return nil
	}

	var result interface{}
	decoder := json.NewDecoder(response.Body)
	decoder.Decode(&result)
	return result
}

func (api apiClient) ChangeCallerID(sip, callerID string) interface{} {
	params := make(url.Values)
	params.Add("id", sip)
	params.Add("number", callerID)
	return api.CallMethod("/v1/sip/callerid/", params, http.MethodPut)
}

func (api apiClient) DirectNumbers() interface{} {
	return api.CallMethod("/v1/direct_numbers/", nil, http.MethodGet)
}

func (api apiClient) Callback(from, to, sip string, isPredicted bool) interface{} {
	params := make(url.Values)
	params.Add("from", from)
	params.Add("to", to)
	if sip != "" {
		params.Add("sip", sip)
	}
	if isPredicted {
		params.Add("isPredicted", "1")
	}

	return api.CallMethod("/v1/request/callback/", params, http.MethodGet)
}

func (api apiClient) GetTotalCalls(callerId string) interface{} {
	params := make(url.Values)
	params.Add("id", callerId)
	return api.CallMethod("/v1/zcrm/calls", params, http.MethodGet)
}
