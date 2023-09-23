package pkg

import (
	"errors"
	"os"
	"strings"
)

func Generate(dir string, apiName string, node *Node) error {
	if node.Key != "structure" {
		return errors.New("invalid sub tree")
	}

	content := dir + "/" + apiName
	err := os.Mkdir(content, os.ModePerm)
	if err != nil {
		return err
	}

	for _, inner := range node.Nodes {
		err := generate(content, inner)
		if err != nil {
			return err
		}
	}

	return nil
}

func generate(dir string, node *Node) error {
	content := dir + "/" + node.Key
	err := os.Mkdir(content, os.ModePerm)
	if err != nil {
		return err
	}

	for _, inner := range node.Nodes {
		err := generate(content, inner)
		if err != nil {
			return err
		}
	}

	for _, value := range node.Values {
		err := createFile(dir, node.Key, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func createFile(path string, dir string, value string) error {
	if value == "main.go" {
		bytes, err := os.ReadFile("templates/code/main.txt")
		if err != nil {
			return err
		}

		return os.WriteFile(path+"/"+dir+"/main.go", bytes, os.ModePerm)
	}

	if value == "entity" {
		bytes, err := os.ReadFile("templates/code/entity.txt")
		if err != nil {
			return err
		}

		stringValue := string(bytes)
		stringValue = strings.ReplaceAll(stringValue, "{package}", dir)
		stringValue = strings.ReplaceAll(stringValue, "{struct}", "Entity")

		return os.WriteFile(path+"/"+dir+"/entity.go", []byte(stringValue), os.ModePerm)
	}

	if value == "entity" {
		bytes, err := os.ReadFile("templates/code/entity.txt")
		if err != nil {
			return err
		}

		stringValue := string(bytes)
		stringValue = strings.ReplaceAll(stringValue, "{package}", dir)
		stringValue = strings.ReplaceAll(stringValue, "{struct}", "Entity")

		return os.WriteFile(path+"/"+dir+"/entity.go", []byte(stringValue), os.ModePerm)
	}

	return nil
}
