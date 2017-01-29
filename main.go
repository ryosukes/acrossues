package main

import (
    "fmt"
    "os"
    "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "agi"
    app.Usage = "get github issues across repositrys"
    app.Action = func(c *cli.Context) error {
    fmt.Println("get github issues across repositrys")
        return nil
    }

    app.Run(os.Args)
}
