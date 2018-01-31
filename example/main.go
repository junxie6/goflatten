package main

import (
	"encoding/json"
	"fmt"
	"github.com/junxie6/goflatten"
)

type Person struct {
	ID            uint
	Name          string
	PermissionArr []string
	HomeArr       []Home
	Table         [][]Column
}

type Home struct {
	City      string
	YearBuilt uint
	RoomArr   []Room
}

type Room struct {
	Name string
	Size float64
}

type Column struct {
	Name  string
	Width int
}

func main() {
	// data holds the flatten key and value result.
	data := make(map[string]interface{})

	p1 := Person{
		ID:            999,
		Name:          "Jun",
		PermissionArr: []string{"VIEW", "EDIT"},
		HomeArr: []Home{
			Home{
				City:      "Vancouver0",
				YearBuilt: 1999,
				RoomArr: []Room{
					Room{
						Name: "Living room",
						Size: 500,
					},
					Room{
						Name: "Master room",
						Size: 600,
					},
				},
			},
			Home{
				City:      "Vancouver1",
				YearBuilt: 2000,
			},
		},
		Table: [][]Column{
			[]Column{
				Column{
					Name:  "Row 0 Col 0",
					Width: 10,
				},
				Column{
					Name:  "Row 0 Col 1",
					Width: 11,
				},
			},
			[]Column{
				Column{
					Name:  "Row 1 Col 0",
					Width: 20,
				},
				Column{
					Name:  "Row 1 Col 1",
					Width: 21,
				},
			},
		},
	}

	goflatten.Flatten(&p1, data, "", false)

	ObjectToJSON(data, true)
}

func ObjectToJSON(v interface{}, isIndent bool) {
	var byteArr []byte
	var err error

	if isIndent == true {
		byteArr, err = json.MarshalIndent(v, "", "    ")
	} else {
		byteArr, err = json.Marshal(v)
	}

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("%s\n", string(byteArr))
}
