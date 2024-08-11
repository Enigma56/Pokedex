package api

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io"
    "log"
)

type AreaDetails struct {
    PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) ListAreaDetails(area string) (AreaDetails, error) {
    url := "https://pokeapi.co/api/v2/location-area/" + area
    dat, ok := c.cache.Get(url)
    if ok {
        fmt.Println("cache hit!")
        var areaDetails AreaDetails
        err := json.Unmarshal(dat, &areaDetails)
        if err != nil {
            fmt.Println(err)
            return AreaDetails{}, err
        }
        
        return areaDetails, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
        return AreaDetails{}, err
    }

    res, err := c.httpClient.Do(req)
    
	if err != nil {
		log.Fatal(err)
        return AreaDetails{}, err
	}

    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if res.StatusCode > 399 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
        return AreaDetails{}, err
	}

    if err != nil {
        return AreaDetails{}, err
    }

    err = c.cache.Add(url, body)
    fmt.Printf("Added url to cache: %s\n", url)
    if err != nil {
        return AreaDetails{}, err
    }

    var areaDetails AreaDetails
    err = json.Unmarshal([]byte(body), &areaDetails)
    if err != nil {
        fmt.Println(err)
        return AreaDetails{}, err
    }

    return areaDetails, nil
}
