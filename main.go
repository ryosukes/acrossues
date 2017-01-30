package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli"
    "golang.org/x/oauth2"
    "github.com/google/go-github/github"
    "github.com/BurntSushi/toml"
)

type Config struct {
    User UserConfig
    Repositories []RepositoryConfig
}

type UserConfig struct {
    Token string    `toml:"token"`
}

type RepositoryConfig struct {
    Owner   string  `toml:"owner"`
    Name    string  `toml:"name"`
}

var config Config

func main() {
    app := cli.NewApp()
    app.Name = "agi"
    app.Usage = "get github issues across repositrys"
    app.Action = func(c *cli.Context) error {
        fmt.Println("get github issues across repositrys")
        return nil
    }

    loadConfig()
    findIssues()

    app.Run(os.Args)
}

func findIssues() {
    var token  = config.User.Token

    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(oauth2.NoContext, ts)
    client := github.NewClient(tc)
    issues, _, err := client.Issues.ListByRepo(config.Repositories[0].Owner, config.Repositories[0].Name, nil)

    if err != nil {
        panic(err)
    }

    fmt.Println(issues)
}

func loadConfig() {
    _, err := toml.DecodeFile("./config.tml", &config)
    if err != nil {
        panic(err)
    }
}
