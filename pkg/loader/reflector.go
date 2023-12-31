package loader

import (
	"github.com/johnfercher/go-arch/pkg"
	"reflect"
)

var mapValue = make(map[string]interface{})
var mapReflection = reflect.ValueOf(mapValue)
var mapType = mapReflection.Type()

func BuildTree(key string, data interface{}) *pkg.Node {
	node := &pkg.Node{
		Key: key,
	}

	if data == nil {
		return node
	}

	dataReflection := reflect.ValueOf(data)
	if dataReflection.IsZero() {
		return node
	}

	if isMap(dataReflection) {
		innerMap := data.(map[string]interface{})
		for key, value := range innerMap {
			innerNode := BuildTree(key, value)
			node.Nodes = append(node.Nodes, innerNode)
		}
	} else {
		values := data.([]interface{})
		for _, value := range values {
			stringValue, ok := value.(string)
			if ok {
				node.Values = append(node.Values, stringValue)
			}
		}
	}

	return node
}

func isMap(value reflect.Value) bool {
	return value.Type() == mapType
}
