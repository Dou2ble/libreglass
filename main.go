package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
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

	url := fmt.Sprintf("%s?%s", IcemanProdTrackerUrl+"/getNearestStops", params.Encode())

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
	url := fmt.Sprintf("%s?%s", IcemanProdTrackerUrl+"/getSalesInfoByStop", params.Encode())

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

// // https://iceman-prod.azurewebsites.net/api/tracker/getVisitedStops/?lastTimeCalled=23:16
// func getVisitedStops(lastTimeCalled time.Time) {
// 	var result IcemanResponse[SalesInfo]

// 	params := url.Values{}
// 	params.Add("stopId", fmt.Sprint(stopId))
// 	url := fmt.Sprintf("%s?%s", IcemanProdTrackerUrl+"/getSalesInfoByStop", params.Encode())
// }

func main() {
	res, err := getNearestSTops(11.0274, 55.3618, 24.1935, 69.0605, math.MaxInt32)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
	fmt.Println(len(res.Data))

	// fmt.Println()
	// fmt.Println()

	// resp, err := getSalesInfoByStop(3160106)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(resp)
}
