package api

import (
    "net/http"
    "log"
    "fmt"
    "io"
    "encoding/json"

    //"github.com/Enigma56/pokedex/internal/cache"
)


func (c *Client) ListLocationAreas(pageURL *string) (LocationArea, error) {
    dat, ok := c.cache.Get(*pageURL)
    if ok {
        fmt.Println("cache hit!")
        var LocationAreaData LocationArea
        err := json.Unmarshal(dat, &LocationAreaData)
        if err != nil {
            fmt.Println(err)
            return LocationArea{}, err
        }
        
        return LocationAreaData, nil
    }

    req, err := http.NewRequest("GET", *pageURL, nil)
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

    err = c.cache.Add(*pageURL, body)
    fmt.Printf("Added url to cache: %s\n", *pageURL)
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
