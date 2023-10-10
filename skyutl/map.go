package skyutl

import (
	"github.com/elliotchance/orderedmap"
)

//GetMapValues function
func GetMapValues(mp map[string]interface{}) []interface{} {
	var result = []interface{}{}
	for _, value := range mp {
		result = append(result, value)
	}

	return result
}

//GetOrderedMapValues function
func GetOrderedMapValues(mp *orderedmap.OrderedMap) []interface{} {
	var result = []interface{}{}
	for _, key := range mp.Keys() {
		value, _ := mp.Get(key)
		result = append(result, value)
	}

	return result
}
