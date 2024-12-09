package constants

import "fmt"

const BASEKEY = "resource:"
const URL_STRING = "code:url_"

func BuildUrlKey(url string) string {
	return fmt.Sprintf("%s%s%s", BASEKEY, URL_STRING, url)
}
