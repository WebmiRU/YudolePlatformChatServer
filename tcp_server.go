package main

//import (
//	"YudolePlatofrmChatServer/types"
//	"encoding/json"
//	"fmt"
//	"log"
//	"net"
//	"os"
//	"slices"
//	"sync"
//)
//
//var tcpSubscribersMutex sync.Mutex
//var tcpSubscribers = make(map[net.Conn][]string)
//
//func tcpServerStart() {
//	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 5801))
//	if err != nil {
//		fmt.Println("Listening error:", err)
//		os.Exit(1)
//	}
//
//	defer server.Close()
//
//	for {
//		conn, err := server.Accept()
//		if err != nil {
//			log.Println("Error accepting:", err.Error())
//			os.Exit(1)
//		}
//
//		go tcpAccept(conn)
//	}
//}
//
//func tcpAccept(conn net.Conn) {
//	decoder := json.NewDecoder(conn)
//
//	for {
//		var message json.RawMessage
//		if err := decoder.Decode(&message); err != nil {
//			break
//		}
//
//		var base types.Base
//		json.Unmarshal(message, &base)
//
//		switch base.Type {
//		case "subscribe":
//			var subscribe types.Subscribe
//			if err := json.Unmarshal(message, &subscribe); err != nil {
//				log.Println("JSON message decode error:", err)
//				continue
//			}
//
//			tcpSubscribersMutex.Lock()
//			for _, v := range subscribe.Events {
//				if slices.Index(tcpSubscribers[conn], v) == -1 {
//					tcpSubscribers[conn] = append(tcpSubscribers[conn], v)
//				}
//			}
//			tcpSubscribersMutex.Unlock()
//			break
//
//		case "unsubscribe":
//			var unsubscribe types.Unsubscribe
//			if err := json.Unmarshal(message, &unsubscribe); err != nil {
//				log.Println("JSON message decode error:", err)
//				continue
//			}
//
//			tcpSubscribersMutex.Lock()
//			for _, v := range unsubscribe.Events {
//				idx := slices.Index(tcpSubscribers[conn], v)
//
//				if slices.Index(tcpSubscribers[conn], v) >= 0 {
//					tcpSubscribers[conn] = slices.Delete(tcpSubscribers[conn], idx, idx+1)
//				}
//			}
//			tcpSubscribersMutex.Unlock()
//			break
//
//		default:
//			broadcast(message, base.Type)
//			break
//		}
//	}
//
//	tcpSubscribersMutex.Lock()
//	delete(tcpSubscribers, conn)
//	tcpSubscribersMutex.Unlock()
//}
