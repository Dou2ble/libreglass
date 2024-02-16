from types import GenericAlias
import requests
from typing import Any, Optional, Generic, TypeVar, TypedDict
from datetime import time

from dataclasses import dataclass, asdict
from dacite import from_dict

T = TypeVar("T")

class IcemanResponse(TypedDict, Generic[T]):
    statusCode: int
    message: str
    data: T


class Stop(TypedDict):
    stopId: int
    longitude: float
    latitude: float
    nextDate: str
    nextTime: str
    distance: float
    routeId: int

class VisitedStop(TypedDict):
    stopId: int
    reason: int # what is this?
    visitedTime: str
    longitude: float
    latitude: float
    routeId: int

    # i have never seen these not being null
    idleTime: Optional[Any]
    distance: Optional[Any]

class SalesInfo(TypedDict):
    salesmanName: str
    phoneNumber: str
    depotName: str
    depotEmail: str
    streetAddress: str
    city: str
    comment: str
    cancelled: bool
    cancelledMessage: Optional[str]

ICEMAN_PROD_TRACKER_URL = "https://iceman-prod.azurewebsites.net/api/tracker"
GET_NEAREST_STOPS_URL = f"{ICEMAN_PROD_TRACKER_URL}/getNearestStops"
GET_SALES_INFO_BY_STOP_URL = f"{ICEMAN_PROD_TRACKER_URL}/getSalesInfoByStop"
GET_VISITED_STOPS_URL = f"{ICEMAN_PROD_TRACKER_URL}/getVisitedStops/"

def get_nearest_stops(min_longitude: float, min_latitude: float, max_longitude: float, max_latitude: float, limit: int) -> IcemanResponse[list[Stop]]:
    parameters = {
        "minLong": min_longitude,
        "minLat": min_latitude,
        "maxLong": max_longitude,
        "maxLat": max_latitude,
        "limit": limit
    }

    response = requests.get(GET_NEAREST_STOPS_URL, params=parameters)
    response.raise_for_status()

    response_data = response.json()
    typed_response = IcemanResponse[list[Stop]](**response_data)

    return typed_response

def get_sales_info_by_stop(stop_id: int) -> IcemanResponse[SalesInfo]:
    parameters = {
        "stopId": stop_id
    }

    response = requests.get(GET_SALES_INFO_BY_STOP_URL, params=parameters)
    response.raise_for_status()

    response_data = response.json()
    typed_response = IcemanResponse[SalesInfo](**response_data)

    return typed_response

def stops_eta(stop_id: int, route_id: int):
    raise NotImplementedError

def get_visited_stops(last_time_called: Optional[time], route_ids: list[int]) -> IcemanResponse[list[VisitedStop]]:
    last_time_called_string = ""
    if last_time_called:
        last_time_called_string = last_time_called.strftime("%H:%M")

    parameters = {
        "lastTimeCalled": last_time_called_string,
    }

    payload = {
        "routeIds": route_ids
    }

    response = requests.post(GET_VISITED_STOPS_URL, params=parameters, json=payload)
    response.raise_for_status()

    response_data = response.json()
    typed_response = IcemanResponse[list[VisitedStop]](**response_data)

    return typed_response





if __name__ == "__main__":
    # stops = get_nearest_stops(11.027368, 55.595919, 23.903332, 69.036518, 5)
    # sales_info = get_sales_info_by_stop(stops["data"][0]["stopId"])
    route_ids = [23190,7154,26340,15777,20462,20456,37362,14788,18093,26478,12981,13144,12980]
    visited_stops = get_visited_stops(None, route_ids)

    print(visited_stops)
    # print(sales_info)

