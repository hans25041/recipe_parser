package recipe_parser

import (
	"fmt"
	"net/url"
	"strings"
)

type HostNameNotMinimalistBakerError string

func (h HostNameNotMinimalistBakerError) Error() string {
	return fmt.Sprintf("Host name was not minimalistbaker.com: %q", string(h))
}

func ParseUrl(raw_url string) (*url.URL, error) {
	recipe_url, err := url.Parse(raw_url)
	if err != nil {
		panic(err)
	}
	hostname := getHostName(recipe_url)
	return recipe_url, HostNameNotMinimalistBakerError(hostname)
}

func getHostName(recipe_url *url.URL) string {
	hostname_parts := strings.Split(recipe_url.Hostname(), ".")
	return strings.Join(hostname_parts[len(hostname_parts)-2:], ".")
}
