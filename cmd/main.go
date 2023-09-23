package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
)

var mapValue = make(map[string]interface{})
var mapReflection = reflect.ValueOf(mapValue)
var mapType = mapReflection.Type()

var arrValue = make([]interface{}, 0)
var arrReflection = reflect.ValueOf(arrValue)
var arrType = arrReflection.Type()

var stringValue string
var stringReflection = reflect.ValueOf(stringValue)
var stringType = stringReflection.Type()

type Node struct {
	Key    string
	Nodes  []*Node
	Values []string
}

func (n *Node) String(depth int) string {
	var content string

	var space string
	for i := 0; i < depth; i++ {
		space += "--"
	}

	content += fmt.Sprintf("%s%s\n", space, n.Key)
	content += fmt.Sprintf("%s%v\n", space, n.Values)

	for _, node := range n.Nodes {
		content += node.String(depth + 1)
	}

	return content
}

func main() {
	bytes, err := os.ReadFile("templates/hexagonal.yml")
	if err != nil {
		log.Fatal(err.Error())
	}

	data := make(map[string]interface{})

	err = yaml.Unmarshal(bytes, &data)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(data)

	node := &Node{}

	for key, value := range data {
		inner := buildTree(key, value)
		node.Nodes = append(node.Nodes, inner)
	}

	value := node.String(0)
	fmt.Println(value)
}

func buildTree(key string, data interface{}) *Node {
	node := &Node{
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
			innerNode := buildTree(key, value)
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

func isArr(i interface{}) bool {
	iReflection := reflect.ValueOf(i)
	return iReflection.Type() == arrType
}

func isString(i interface{}) bool {
	iReflection := reflect.ValueOf(i)
	return iReflection.Type() == stringType
}
