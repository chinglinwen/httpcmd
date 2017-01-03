package main

import (
        "flag"
        "fmt"
        "log"
        "net/http"
        "os"
        "os/exec"
        "strings"
)

var example = `
Example:
  $ curl localhost:9002 -F cmd="./a.sh hello"
  hello
  error will output here too(if it has error)
  exit status 1
  $ curl localhost:9002 -F cmd="./b.sh hello"                          
  hello
  $
`

func main() {
        flag.Usage = func() {
                fmt.Println("Usage of httpcmd:")
                flag.PrintDefaults()
                fmt.Printf(example)
        }

        port := flag.String("port", "9002", "Port number.")
        version := flag.Bool("v", false, "Show version.")

        flag.Parse()

        //Display version info.
        if *version {
                fmt.Println("httpcmd version=1.0, 2017-1-3")
                os.Exit(0)
        }

        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":"+*port, nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
        c := r.FormValue("cmd")
        cmds := strings.Fields(c)
        if len(cmds) <= 0 {
                fmt.Fprintf(w, "no command received\n")
                return
        }

        var cmd *exec.Cmd
        if len(cmds) == 1 {
                cmd = exec.Command(cmds[0])
        } else {
                cmd = exec.Command(cmds[0], cmds[1:]...)
        }

        cmd.Stdout = w
        cmd.Stderr = w
        err := cmd.Run()
        if err != nil {
                fmt.Fprintf(w, "%v\n", err)
        }
}