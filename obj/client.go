package obj

import (
	"encoding/json"
	"slices"
	"sync"
)

var clientMutex sync.Mutex
var Clients []*Client

type Client struct {
	Write     func(message json.RawMessage) error
	Subscribe []string
}

func ClientAppend(client *Client) {
	clientMutex.Lock()
	Clients = append(Clients, client)
	clientMutex.Unlock()
}

func ClientRemove(client *Client) {
	clientMutex.Lock()
	var idx = slices.Index(Clients, client)
	Clients = slices.Delete(Clients, idx, idx+1)
	clientMutex.Unlock()
}
