package mapcode

import "fmt"

func Search(myMap map[string]string, key string) (string, error) {
	if _, found := myMap[key]; !found {
		return "", fmt.Errorf("%s key is not present", key)
	}
	return myMap[key], nil
}
