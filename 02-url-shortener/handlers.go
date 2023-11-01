package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

type shortener struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

func mapHandler(paths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if val, ok := paths[path]; ok {
			http.Redirect(w, r, val, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func yamlHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse yaml
	var yamlShortener []shortener
	err := yaml.Unmarshal(yamlData, &yamlShortener)
	if err != nil {
		return nil, err
	}

	// Convert yaml to map
	urlPaths := make(map[string]string)
	for _, y := range yamlShortener {
		urlPaths[y.Path] = y.Url
	}

	// Invoke map handler
	return mapHandler(urlPaths, fallback), nil
}

func jsonHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse JSON data
	jsonShortener := []shortener{}
	err := json.Unmarshal(jsonData, &jsonShortener)
	if err != nil {
		return nil, err
	}

	// Convert it into a map
	pathsMap := make(map[string]string)
	for _, j := range jsonShortener {
		pathsMap[j.Path] = j.Url
	}

	// Invoke map handler
	return mapHandler(pathsMap, fallback), nil
}
