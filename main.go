package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli"
    "golang.org/x/oauth2"
    "github.com/google/go-github/github"
)

func main() {
    app := cli.NewApp()
    app.Name = "agi"
    app.Usage = "get github issues across repositrys"
    app.Action = func(c *cli.Context) error {
        fmt.Println("get github issues across repositrys")
        return nil
    }

    findIssues()

    app.Run(os.Args)
}

func findIssues() {
    const token = "hoge"

    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(oauth2.NoContext, ts)
    client := github.NewClient(tc)
    issues, _, err := client.Issues.ListByRepo("fuga", "piyo", nil)

    if err != nil {
        panic(err)
    }

    fmt.Println(issues)
}
