# jumangok
## Description
jumangok is a go package to carry out Japanese morphological analysis with jumanpp.

You can try REPL or serve easy-to-use API.

It is faster than the default script/server.rb a little bit.

## Features
- repl
- http server
- JSON-API
- metatag is parsed (developing).

## Installation and Usages
### install with go command:
First, install jumanpp.

After that, if you have already installed go command, the installation is almost done.

    go get github.com/u3paka/jumangok
    jumangok serve

### install and serve with docker:
if you don't want to install jumannpp and go directly your environment, you can use a docker image!

    docker run -p 12000:12000 u3paka/jumangok jumangok serve

### install and try REPL with docker:
    
    docker run -it u3paka/jumangok jumangok repl


## Usage as a golang package
### import
    import (
        "github.com/u3paka/jumangok/jmg"
    )

### client-usage

    in := "風が語りかけます。美味い美味すぎる! 饅頭。埼玉銘菓。"
    ws, err := jmg.NewClient("localhost:12000").Jumanpp(context.Background(), in)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println(ws)

### Word struct data and tools (developing)
    //Extract function
    ews := jmg.Extract(ws, func(w *Word) bool {
		if w.HasDomain("料理・食事") {
		    return true
		}
		return false
	})
    fmt.Println(ews[0].Surface) // 饅頭


## Link
黒橋・河原研究室 様　日本語形態素解析システム JUMAN++ 

http://nlp.ist.i.kyoto-u.ac.jp/index.php?JUMAN%2B%2B

## Author
u3paka

## Licence
MIT