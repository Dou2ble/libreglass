package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type IcemanResponse[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
}

type Stop struct {
	StopId   int     `json:"stopId"`
	Lon      float64 `json:"longitude"`
	Lat      float64 `json:"latitude"`
	NextDate string  `json:"nextDate"`
	NextTime string  `json:"nextTime"`
	Distance float64 `json:"distance"`
	RouteId  int     `json:"routeId"`
}

type SalesInfo struct {
	SalesmanName     string  `json:"salesmanName"`
	PhoneNumber      string  `json:"phoneNumber"`
	DepotName        string  `json:"depotName"`
	DepotEmail       string  `json:"depotEmail"`
	StreetAddress    string  `json:"streetAddress"`
	City             string  `json:"city"`
	Comment          string  `json:"comment"`
	Cancelled        bool    `json:"cancelled"`
	CancelledMessage *string `json:"cancelledMessage"`
}

const IcemanProdTrackerUrl = "https://iceman-prod.azurewebsites.net/api/tracker"

func getNearestSTops(minLon, minLat, maxLon, maxLat float64, limit int32) (IcemanResponse[[]Stop], error) {
	var result IcemanResponse[[]Stop]

	params := url.Values{}

	params.Add("minLong", fmt.Sprint(minLon))
	params.Add("minLat", fmt.Sprint(minLat))
	params.Add("maxLong", fmt.Sprint(maxLon))
	params.Add("maxLat", fmt.Sprint(maxLat))
	params.Add("limit", fmt.Sprint(limit))

	url := IcemanProdTrackerUrl + "/getNearestStops?" + params.Encode()

	response, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func getSalesInfoByStop(stopId int) (IcemanResponse[SalesInfo], error) {
	var result IcemanResponse[SalesInfo]

	params := url.Values{}
	params.Add("stopId", fmt.Sprint(stopId))
	url := IcemanProdTrackerUrl + "/getSalesInfoByStop?" + params.Encode()

	response, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func stopsEta(stopId int, routeId int) (IcemanResponse[string], error) {
	var result IcemanResponse[string]

	params := url.Values{}
	params.Add("stopId", fmt.Sprint(stopId))
	params.Add("routeId", fmt.Sprint(routeId))
	url := IcemanProdTrackerUrl + "/getSalesInfoByStop?" + params.Encode()

	response, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

// TODO: this function should not return a string. but ATM i don't know what the data type is
func getVisitedStops(lastTimeCalled *time.Time, routeIds []int) (string, error) {
	var result string

	// if lastTimeCalled dose exsist then we should format it and save it in a new variable otherwise it should be empty
	lastTimeCalledString := ""
	if lastTimeCalled != nil {
		lastTimeCalledString = lastTimeCalled.Format("15:04")
	}

	params := url.Values{}
	params.Add("lastTimeCalled", lastTimeCalledString)
	url := IcemanProdTrackerUrl + "/getVisitedStops?" + params.Encode()

	// Convert the array of ints to a JSON-encoded request body
	requestBody, err := json.Marshal(map[string][]int{"routeIds": routeIds})
	if err != nil {
		return result, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return result, nil
	}

	return string(body), nil
}

// // https://iceman-prod.azurewebsites.net/api/tracker/liverouteinfo/37816
// func liveRouteInfo(routeId int) {

// }

func main() {
	uwu := []int{13, 12, 12, 12}
	result, err := getVisitedStops(nil, uwu)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
