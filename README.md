[![Build Status](https://img.shields.io/travis/retailcrm/mg-bot-api-client-go/master.svg?logo=travis&style=flat-square)](https://travis-ci.org/retailcrm/mg-bot-api-client-go)
[![GitHub release](https://img.shields.io/github/release/retailcrm/mg-bot-api-client-go.svg?style=flat-square)](https://github.com/retailcrm/mg-bot-api-client-go/releases)
[![GoLang version](https://img.shields.io/badge/go-1.8%2C%201.9%2C%201.10%2C%201.11-blue.svg?style=flat-square)](https://golang.org/dl/)
[![Godoc reference](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/retailcrm/mg-bot-api-client-go)


# Message Gateway Bot API Go client

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

## Websocket Example

```golang
package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/retailcrm/mg-bot-api-client-go/v1"
)

func main() {
	var client = v1.New("https://token.url", "cb8ccf05e38a47543ad8477d49bcba99be73bff503ea6")

	url, headers, err := client.WsMeta([]string{"message_new"})
	if err != nil {
		log.Fatal("wsMeta:", err)
	}

	wsConn, _, err := websocket.DefaultDialer.Dial(url, headers)
	if err != nil {
		log.Fatal("dial:", err)
	}

	for {
		var wsEvent v1.WsEvent
		err = wsConn.ReadJSON(&wsEvent)
		if err != nil {
			log.Fatal("ReadJSON:", err)
		}

		var eventData v1.WsEventMessageNewData
		err = json.Unmarshal(wsEvent.Data, &eventData)
		if err != nil {
			log.Fatal("Unmarshal:", err)
		}

		if !strings.HasPrefix(eventData.Message.Content, "Hello") {
			continue
		}

		message := v1.MessageSendRequest{
			Scope:   "public",
			Content: "Bonjour!",
			ChatID:  eventData.Message.ChatID,
		}

		_, status, err := client.MessageSend(message)
		if err != nil {
			log.Fatalf("%d %v", status, err)
		}
	}
}
```

### Documentation

* [English](https://help.retailcrm.pro/Developers/MgBot)
* [Russian](https://help.retailcrm.ru/Developers/MgBot)
