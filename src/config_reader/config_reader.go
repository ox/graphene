package config_reader

import (
  "io/ioutil"
  "encoding/json"
  )

func ReadConfigFile(path string) map[string]interface{} {
  b, err := ioutil.ReadFile(path)
  if err != nil {
    panic(err)
  }

  var f interface{}
  err = json.Unmarshal(b, &f)
  if err != nil {
    panic(err)
  }

  return f.(map[string]interface{})
}
