package controller

import (
	"custom-command-line/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadFile() []model.CustomData {
	file, _ := ioutil.ReadFile("test.json")

	data := []model.CustomData{}

	_ = json.Unmarshal([]byte(file), &data)

	return data
}
func WriteFile() {
	data := []model.CustomData{
		model.CustomData{
			Name:        "hello",
			Help:        "this is my help",
			Description: "this is description",
			Str:         nil,
		},
		model.CustomData{

			Name:        "tittle1",
			Help:        "this is my help",
			Description: "this is description",
			Str:         nil,
		},
		model.CustomData{

			Name:        "tittle2",
			Help:        "this is my help",
			Description: "this is description",
			Str:         nil,
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile("test.json", file, 0644)
	if err != nil {
		fmt.Println("errors: ", err)
	} else {
		fmt.Println("successfully")
	}
}
func AddData(dataAdd model.CustomData) {
	file, _ := ioutil.ReadFile("test.json")
	data := []model.CustomData{}
	_ = json.Unmarshal([]byte(file), &data)
	for i := 0; i < len(data); i++ {
		if data[i].Name == dataAdd.Name {
			fmt.Println("can't not add")
			return
		}
	}
	data = append(data, dataAdd)
	fileWrite, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile("test.json", fileWrite, 0644)
	if err != nil {
		fmt.Println("errors: ", err)
	} else {
		fmt.Println("successfully")
	}

}
