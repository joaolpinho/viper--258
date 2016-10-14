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

	var err error = nil

	v := viper.New()


	v.AutomaticEnv()
	v.SetEnvPrefix(PREFIX)
	v.SetConfigType("yaml")



	var yamlExample = []byte(`
		Hacker: true ""
		Name: steve
		Hobbies: hobb-its
		Clothing: "clothing"
		Age: 35
	`)
	err = v.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		fmt.Println(err)
	}



	type Config struct {
		Hacker   bool
		Name     string
		Hobbies  string
		Clothing string
		Age      int
	}

	cf := Config{}

	err = v.Unmarshal(&cf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v == %v\n",v.Get("Hacker"), cf.Hacker )
	fmt.Printf("%v == %v\n",v.Get("Name"), cf.Name )
	fmt.Printf("%v == %v\n",v.Get("Hobbies"), cf.Hobbies )
	fmt.Printf("%v == %v\n",v.Get("Clothing"), cf.Clothing )
	fmt.Printf("%v == %v\n",v.Get("Age"), cf.Age )
}