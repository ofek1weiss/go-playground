package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		redirectUrl, ok := pathsToUrls[request.RequestURI]
		if !ok {
			fallback.ServeHTTP(writer, request)
			return
		}
		http.Redirect(writer, request, redirectUrl, http.StatusFound)
	}
}

type yamlRecord struct {
	Path string `yaml:path`
	Url  string `yaml:url`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var config []yamlRecord
	err := yaml.Unmarshal(yml, &config)
	if err != nil {
		return nil, err
	}
	handler := func(writer http.ResponseWriter, request *http.Request) {
		for _, record := range config {
			if record.Path == request.RequestURI {
				http.Redirect(writer, request, record.Url, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(writer, request)
	}
	return handler, nil
}
