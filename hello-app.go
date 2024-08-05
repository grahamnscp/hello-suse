package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
)

var Hostname, oserr = os.Hostname()

func main() {
  if oserr != nil {
    fmt.Println("io.hostname error: ", oserr, " hostname value:", Hostname)
  }

  http.HandleFunc("/", index)
  http.HandleFunc("/hello", helloHandler)
  http.HandleFunc("/hello/", helloHandler)
  http.HandleFunc("/suse", suseHandler)
  http.HandleFunc("/suse/", suseHandler)
  http.HandleFunc("/docker", dockerHandler)
  http.HandleFunc("/docker/", dockerHandler)
  http.HandleFunc("/whale", whaleHandler)
  http.HandleFunc("/whale/", whaleHandler)
  err := http.ListenAndServe(":8080", nil)

  if err != nil {
    fmt.Println("Serve Http:", err)
  }
}

func index(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"\n")
  fmt.Fprint(w," _          _ _\n")
  fmt.Fprint(w,"| |__   ___| | | ___\n")
  fmt.Fprint(w,"| '_ \\ / _ \\ | |/ _ \\\n")
  fmt.Fprint(w,"| | | |  __/ | | (_) |\n")
  fmt.Fprint(w,"|_| |_|\\___|_|_|\\___/\n")
  fmt.Fprint(w,"\n")
  dt := time.Now()
  fmt.Fprint(w,"[", dt.Format("01-02-2006 15:04:05.00"), "] ", "Container hostname: ", Hostname, "\n")
  fmt.Println("[",dt.Format("01-02-2006 15:04:05.00"),"defaultHandler]","Container hostname:",Hostname)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"\n")
  fmt.Fprint(w," _           _    _\n")
  fmt.Fprint(w,"| |     ___ | |  | |    ___\n")
  fmt.Fprint(w,"| |___ / _ \\| |  | |   / _ \\\n")
  fmt.Fprint(w,"|  _  |  __/| |__| |__| (_) |\n")
  fmt.Fprint(w,"|_| |_|\\___/ \\___|\\___|\\___/\n")
  fmt.Fprint(w,"\n")
  dt := time.Now()
  fmt.Fprint(w,"[", dt.Format("01-02-2006 15:04:05.00"), "] ", "Container hostname: ", Hostname, "\n")
  fmt.Println("[",dt.Format("01-02-2006 15:04:05.00"),"helloHandler]","Container hostname:",Hostname)
}

func suseHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"\n")
  fmt.Fprint(w,"  ____  _   _ ____  _____ \n")
  fmt.Fprint(w," / ___|| | | / ___|| ____|\n")
  fmt.Fprint(w," \\___ \\| | | \\___ \\|  _|  \n")
  fmt.Fprint(w,"  ___) | |_| |___) | |___ \n")
  fmt.Fprint(w," |____/ \\___/|____/|_____|\n")
  fmt.Fprint(w,"\n")
  dt := time.Now()
  fmt.Fprint(w,"[", dt.Format("01-02-2006 15:04:05.00"), "] ", "Container hostname: ", Hostname, "\n")
  fmt.Println("[",dt.Format("01-02-2006 15:04:05.00"),"suseHandler]","Container hostname:",Hostname)
}

func dockerHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"\n")
  fmt.Fprint(w,"     _            _\n")
  fmt.Fprint(w,"  __| | ___   ___| | _____ _ __\n")
  fmt.Fprint(w," / _  |/ _ \\ / __| |/ / _ \\ '__|\n")
  fmt.Fprint(w,"| (_| | (_) | (__|   <  __/ |\n")
  fmt.Fprint(w," \\____|\\___/ \\___|_|\\_\\___|_|\n")
  fmt.Fprint(w,"\n")
  dt := time.Now()
  fmt.Fprint(w,"[", dt.Format("01-02-2006 15:04:05.00"), "] ", "Container hostname: ", Hostname, "\n")
  fmt.Println("[",dt.Format("01-02-2006 15:04:05.00"),"dockerHandler]","Container hostname:",Hostname)
}

func whaleHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"\n")
  fmt.Fprint(w,"                   ##\n")
  fmt.Fprint(w,"             ## ## ##        ==\n")
  fmt.Fprint(w,"          ## ## ## ## ##    ===\n")
  fmt.Fprint(w,"      /`````````````````\\___/ ===\n")
  fmt.Fprint(w," ~~~ {~~ ~~~~ ~~~ ~~~~ ~~~ ~~/~ === ~~~\n")
  fmt.Fprint(w,"      \\______ o           __/\n")
  fmt.Fprint(w,"        \\    \\         __/\n")
  fmt.Fprint(w,"         \\____\\_______/\n")
  fmt.Fprint(w,"\n")
  dt := time.Now()
  fmt.Fprint(w,"[", dt.Format("01-02-2006 15:04:05.00"), "] ", "Container hostname: ", Hostname, "\n")
  fmt.Println("[",dt.Format("01-02-2006 15:04:05.00"),"whaleHandler]","Container hostname:",Hostname)
}

