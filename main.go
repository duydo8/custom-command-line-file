package main

import (
	"custom-command-line/controller"
	"custom-command-line/model"
	"fmt"
	"github.com/abiosoft/ishell/v2"
)

func init() {
	controller.WriteFile()

	data := model.CustomData{
		Name:        "tittle3",
		Help:        "say hello",
		Description: "say hello",
		Str:         nil,
	}
	controller.AddData(data)

}
func exitFunc(c *ishell.Context) {
	c.Stop()
}
func main() {
	shell := ishell.New()
	data := controller.ReadFile()
	shell.AddCmd(&ishell.Cmd{
		Name: data[0].Name,
		Help: data[0].Help,
		Func: func(c *ishell.Context) {
			c.ShowPrompt(false)
			defer c.ShowPrompt(true)
			c.Println("demo parameter")
			c.Print("Username: ")
			username := c.ReadLine()
			c.Print("Password: ")
			password := c.ReadPassword()
			c.Println("username:", username, "\n", "password:", password)

		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "demo-argument",
		Help: "demo-argument",
		Func: func(c *ishell.Context) {
			x := c.Args[0:]
			c.Println(x)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "switch",
		Help: "demo-switch",
		Func: func(c *ishell.Context) {
			x := c.Args[0:]
			i := 0
			var check, check2 bool
			for i < len(x) {
				for j := 0; j < len(data); j++ {

					switch x[i] {
					case data[j].Name:
						fmt.Println(data[j].Description)
						check2 = true
						break
					default:
						check = true
					}
				}
				i++
			}
			if check == true && check2 == false {
				c.Println("no thing matches")
			}
		},
	})
	for i := 1; i < len(data)-1; i++ {
		shell.AddCmd(&ishell.Cmd{
			Name: data[i].Name,
			Help: data[i].Help,
			Func: func(c *ishell.Context) {
				c.Println(data[i].Description)
			},
		})
	}
	shell.Run()
}
