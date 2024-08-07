package api

import (
    "net/http"
    "log"
    "fmt"
    "io"
    "encoding/json"
)

func (c *Client) ListLocationAreas() (LocationArea, error) {
    req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area", nil)
    if err != nil {
        log.Fatal(err)
        return LocationArea{}, err
    }

    res, err := c.httpClient.Do(req)
    
	if err != nil {
		log.Fatal(err)
        return LocationArea{}, err
	}

    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        return LocationArea{}, err
	}

    if err != nil {
        return LocationArea{}, err
    }


    var LocationAreaData LocationArea
    err = json.Unmarshal([]byte(body), &LocationAreaData)
    if err != nil {
        fmt.Println(err)
        return LocationArea{}, err
    }

    return LocationAreaData, nil
}
