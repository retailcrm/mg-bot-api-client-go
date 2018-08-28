[![Build Status](https://img.shields.io/travis/retailcrm/mg-bot-api-client-go/master.svg?style=flat-square)](https://travis-ci.org/retailcrm/mg-bot-api-client-go)
[![GitHub release](https://img.shields.io/github/release/retailcrm/mg-bot-api-client-go.svg?style=flat-square)](https://github.com/retailcrm/mg-bot-api-client-go/releases)
[![GoLang version](https://img.shields.io/badge/GoLang-1.9%2C%201.10%2C%201.11-blue.svg?style=flat-square)](https://golang.org/dl/)


# retailCRM Message Gateway Bot API Go client

## Install

```bash
go get -u -v github.com/retailcrm/mg-bot-api-client-go
```

## Usage

```golang
package main

import (
	"fmt"
	"net/http"

	"github.com/retailcrm/mg-bot-api-client-go/v1"
)

func main() {
    var client = v1.New("https://token.url", "cb8ccf05e38a47543ad8477d49bcba99be73bff503ea6")
    message := MessageSendRequest{
        Scope:   "public",
        Content: "test",
        ChatID:  12,
    }

    data, status, err := c.MessageSend(message)
    if err != nil {
        t.Errorf("%d %v", status, err)
    }

    fmt.Printf("%v", data.MessageID)
}
```

## Documentation

* [GoDoc](https://godoc.org/github.com/retailcrm/mg-bot-api-client-go)
