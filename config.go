package main

import (
   "io/ioutil"
   "os"
   "github.com/influxdata/toml"
   "fmt"
)

type Gconfig struct {
   Common struct {
      Homedir     string
   }

   Repo struct {
      Ip          string
      Port        string
      Service     string
      User        string
      Pass        string
   }
}

func (c *Gconfig) LoadConfig(v_file string) error {

   s_module := "[Config:LoadConfig]"


   if _, err := os.Stat(v_file); err != nil {
      fmt.Printf("E! %s %s is not exist\n", s_module, v_file)
      return err
   }

   f, err := os.Open(v_file)
   if err != nil {
      fmt.Printf("E! %s %s is open error\n", s_module, v_file)
      return err
   }
   defer f.Close()

   buf, err := ioutil.ReadAll(f)
   if err != nil {
      fmt.Printf("E! %s ioutil error : %s\n", s_module, err)
      return err
   }

   if err := toml.Unmarshal(buf, c); err != nil {
      fmt.Printf("E! %s toml.Unmarshal error : %s\n", s_module, err)
      return err
   }

   return nil
}
