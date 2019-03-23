package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/u3paka/jumangok/jmg"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "repl",
			Aliases: []string{"r"},
			Usage:   "start jumanpp REPL(almost same as vanilla-jumanpp)",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "jumanpp",
					Value: "jumanpp",
					Usage: "place of libjumanpp",
				},
				cli.StringSliceFlag{
					Name:  "option",
					Value: &cli.StringSlice{},
					Usage: "jumanpp option",
				},
			},
			Action: func(clc *cli.Context) error {
				jcl, err := jmg.NewService(clc.String("jumanpp"), clc.StringSlice("option")...)
				if err != nil {
					log.Fatal(err)
					return err
				}
				s := bufio.NewScanner(os.Stdin)
				fmt.Println("JUMAN++[REPL] sentences + Enter")
				for s.Scan() {
					fmt.Println(jcl.RawParse(s.Text()))
				}
				if s.Err() != nil {
					// non-EOF error.
					log.Fatal(s.Err())
				}
				select {}
			},
		},
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "start jumanpp go-server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "jumanpp",
					Value: "jumanpp",
					Usage: "place of libjumanpp",
				},
				cli.StringSliceFlag{
					Name:  "option",
					Value: &cli.StringSlice{},
					Usage: "jumanpp option",
				},
				cli.IntFlag{
					Name:  "port",
					Value: 12000,
					Usage: "a port number to serve",
				},
			},
			Action: func(clc *cli.Context) error {
				jcl, err := jmg.NewService(clc.String("jumanpp"), clc.StringSlice("option")...)
				if err != nil {
					return err
				}
				// 公式と同じ ただしsocket -> http
				http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						fmt.Println(err)
						return
					}
					in := string(body)
					fmt.Println(in)
					out := jcl.RawParse(in)
					fmt.Fprint(w, out)
				})
				http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
					if r.Method != "POST" {
						return
					}
					body, err := ioutil.ReadAll(r.Body)
					if err != nil {
						fmt.Println(err)
						return
					}
					in := string(body)
					fmt.Println(in)
					out, err := jcl.GetWords(in)
					if err != nil {
						fmt.Println(err)
						return
					}
					outjson, err := json.Marshal(out)
					if err != nil {
						fmt.Println(err)
					}
					w.Header().Set("Content-Type", "application/json")
					fmt.Fprint(w, string(outjson))
				})
				return http.ListenAndServe(":"+strconv.Itoa(clc.Int("port")), nil)
			},
		},
	}
	app.Run(os.Args)
}
