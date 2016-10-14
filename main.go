package main

import (
	"github.com/spf13/viper"
	"fmt"
	"bytes"
)


func main(  ) {

	const (
		//URI = "consul://localhost:8500/"
		//PATH = "/"
		PREFIX = ""
	)

	v := viper.New()


	v.AutomaticEnv()
	v.SetEnvPrefix(PREFIX)
	v.SetConfigType("yaml")



	var yamlExample = []byte(`
		Hacker: true
		name: steve
		hobbies: hobb-its
		clothing: "clothing"
		age: 35
	`)
	v.ReadConfig(bytes.NewBuffer(yamlExample))




	type Config struct {
		Hacker bool
		name string
		hobbies string
		clothing string
		age int
	}

	cf := Config{}

	err := v.Unmarshal(&cf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v == %v\n",v.Get("Hacker"), cf.Hacker )
	fmt.Printf("%v == %v\n",v.Get("name"), cf.name )
	fmt.Printf("%v == %v\n",v.Get("hobbies"), cf.hobbies )
	fmt.Printf("%v == %v\n",v.Get("clothing"), cf.clothing )
	fmt.Printf("%v == %v\n",v.Get("age"), cf.age )
}