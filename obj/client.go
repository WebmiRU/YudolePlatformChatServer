package obj

import (
	"encoding/json"
	"slices"
	"sync"
)

var ClientMutex sync.Mutex
var Clients []*Client

type Client struct {
	Write     func(message json.RawMessage) error
	Subscribe []string
}

func ClientAppend(client *Client) {
	ClientMutex.Lock()
	Clients = append(Clients, client)
	ClientMutex.Unlock()
}

func ClientRemove(client *Client) {
	ClientMutex.Lock()
	var idx = slices.Index(Clients, client)
	Clients = slices.Delete(Clients, idx, idx+1)
	ClientMutex.Unlock()
}
