package main


import (
  "os"
  "fmt"
  "flag"
  "time"
  "strconv"
  "gopkg.in/rana/ora.v4"
  pw "password"
)


func main() {

  var o_config Gconfig

  f_cfg := flag.String("conf", "./conf/gpass.conf", "gpass config")
  flag.Parse()

  if o_Err := o_config.LoadConfig(*f_cfg); o_Err != nil {
     fmt.Printf("Don't load config file. check it\n %s", o_Err)
     os.Exit(100)
  }

  o_Pw, o_Err := pw.NewGenerator(nil)

  if o_Err != nil {
     fmt.Println("Don't make pasword object : %s", o_Err)
     os.Exit(1)
  }

  s_Pw, o_Err := o_Pw.Generate(10, 1, 2, false, true)

  if o_Err != nil {
     fmt.Println("Don't make string pasword : %s", o_Err)
     os.Exit(1)
  }

  s_Db_url := fmt.Sprintf("%s/%s@%s:%s/%s", o_config.Repo.User,
     o_config.Repo.Pass, o_config.Repo.Ip, o_config.Repo.Port,
     o_config.Repo.Service)
     
  o_Env, o_Srv, o_Ses, o_Err := ora.NewEnvSrvSes(s_Db_url)
  if o_Err != nil {
     fmt.Println("oracle connect err!!")
  }
  defer o_Env.Close()
  defer o_Srv.Close()
  defer o_Ses.Close()

  o_Time := time.Now().AddDate(0, 1, 0)
  i_Month := int(o_Time.Month())

  fmt.Printf("%s년 %s월 password : %s", strconv.Itoa(o_Time.Year()), strconv.Itoa(i_Month), s_Pw)

  o_Stmt, o_Err := o_Ses.Prep(fmt.Sprintf(
     "update tb_db_pwd set mon_pwd='%s', upd_dt=sysdate where mon_id = %d",
        s_Pw, i_Month))

  defer o_Stmt.Close()
  i_rowsAffected, o_Err := o_Stmt.Exe()

  if i_rowsAffected == 1 {
     fmt.Println(", change ok")
  } else {
     fmt.Println(", change fail")
  }
}     
