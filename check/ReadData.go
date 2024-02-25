package check

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Item struct {
    Name  string
    Price int
}
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func InsertJsonData(){
			jsonList := []string{
		`{"name": "Alice", "age": 30}`,
		`{"name": "Bob", "age": 35}`,
		`{"name": "rick", "age": 45}`,
		`{"name": "bred", "age": 55}`,
		`{"name": "marty", "age": 65}`,
	}
	 for i, jsonData := range jsonList {
        var person Person
        if err := json.Unmarshal([]byte(jsonData), &person); err != nil {
            fmt.Printf("Error parsing JSON for entry %d: %v\n", i, err)
            continue
        }
		    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
	}
	}
func InsertJsonFromAnotherData(){
	file, err := os.Open("check/names.json")
		if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		var person Person
		if err := json.Unmarshal([]byte(line), &person); err != nil {
			fmt.Printf("Error parsing JSON on line %d: %v\n", lineNumber, err)
			continue
		}
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
	}
}
func InsertData(){
		file, err := os.Open("check/d.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)  
	for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ",")
		fmt.Print(parts, "----parts")
        if len(parts) != 2 {
            fmt.Println("Invalid line format:", line)
            continue
        }
		name := strings.TrimSpace(parts[0])
		mPrice:= strings.TrimSpace(parts[1])
        price, err := strconv.Atoi(mPrice)
		        if mPrice == "" {
            fmt.Println("Empty price:", line)
            continue
        }
        if name == "" {
            fmt.Println("Empty price:", line)
            continue
        }
        if err != nil {
            fmt.Println("Error parsing price:", err)
            continue
        }
        item := Item{Name: name, Price: price}
		fmt.Print(item)
}
}