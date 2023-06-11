package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/lulzshadowwalker/saoirse/internal/helpers"
)

func Fetch(city string) (*Weather, error) {
	url := fmt.Sprintf("https://weatherapi-com.p.rapidapi.com/current.json?q=%s", city)

	req, _ := http.NewRequest("GET", url, nil)

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	config, err := helpers.ReadConfig[map[string]any](filepath.Join(cwd, "config/rapid-api-config.json"))
	if err != nil {
		log.Fatalf(`
error reading weather api config.
%q`, err)
	}

	req.Header.Add("X-RapidAPI-Key", (*config)["api-key"].(string))
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	weather := Weather{}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}
