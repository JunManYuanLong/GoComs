package config

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	simplejson "github.com/bitly/go-simplejson"
)

func Load(path string, cfg interface{}) error {

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(cfg)
	err = json.Unmarshal(content, &cfg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cfg)
	return json.Unmarshal(content, cfg)
}

func JsonMapLoad(fileName string) *simplejson.Json {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	js, jserr := simplejson.NewJson(content)
	if jserr != nil {
		panic(jserr)
	}
	return js
}
