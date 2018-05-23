package main

import "fmt"

func main() {
	var m1 = map[string]string{
		"name":     "DTC",
		"class":    "english",
		"location": "Hanzhou",
	}
	m2 := map[string]string{
		"name":     "DTC",
		"class":    "english",
		"location": "Hanzhou",
	}
	m3 := make(map[string]string)
	fmt.Println(m1, m2, m3)

	for k, v := range m1 {
		fmt.Println(k, v)
	}

	m2["name"] = "dtc"
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	m4 := m2
	for k, v := range m4 {
		fmt.Println(k, v)
	}

	fmt.Println(".........................")
	fmt.Println(m2["lastname"])
	v, exist := m2["lastname"]
	fmt.Println(v, exist)
	v, exist = m2["name"]
	fmt.Println(v, exist)

	fmt.Println(".........................")
	delete(m2, "name")
	for k, v := range m2 {
		fmt.Println(k, v)
	}

}
