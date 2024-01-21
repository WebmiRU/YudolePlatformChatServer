package twitch

import (
	"YudolePlatofrmChatServer/obj"
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	SocketPing      = 20
	SocketWait      = 30
	SocketReconnect = 10
)

var re = regexp.MustCompile(`^(?:@([^\r\n ]*) +|())(?::([^\r\n ]+) +|())([^\r\n ]+)(?: +([^:\r\n ]+[^\r\n ]*(?: +[^:\r\n ]+[^\r\n ]*)*)|())?(?: +:([^\r\n]*)| +())?[\r\n]*$`)
var socket net.Conn
var lastDataReceived = time.Now()
var isIrcConnected = false

func processTags(tags string) map[string]string {
	result := make(map[string]string)

	if len(tags) == 0 {
		return result
	}

	for _, v := range strings.Split(tags, ";") {
		kv := strings.SplitN(v, "=", 2)

		if len(kv) == 2 {
			result[kv[0]] = kv[1]
		}
	}

	return result
}

func processSmiles(message IRCMessage) string {
	msg := []rune(message.Text)
	offset := 0

	if _, ok := message.Tags["emotes"]; !ok {
		return message.Text
	}

	if len(message.Tags["emotes"]) == 0 {
		return message.Text
	}

	for _, smile := range strings.Split(message.Tags["emotes"], "/") {
		smileIdFromTo := strings.Split(smile, ":")
		smileId := smileIdFromTo[0]

		for _, fromTo := range strings.Split(smileIdFromTo[1], ",") {
			smileFromTo := strings.Split(fromTo, "-")
			smileFrom, _ := strconv.Atoi(smileFromTo[0])
			smileTo, _ := strconv.Atoi(smileFromTo[1])
			smileText := msg[smileFrom+offset : smileTo+offset+1]
			smileReplacer := []rune(fmt.Sprintf("<img class=\"smile twitch\" src=\"https://static-cdn.jtvnw.net/emoticons/v2/%s/default/dark/1.0\" alt=\"%s\"/>", smileId, string(smileText)))
			msg = append(msg[:smileFrom+offset], append(smileReplacer, msg[smileTo+1+offset:]...)...)
			offset += smileFrom - smileTo + len(smileReplacer) - 1
		}
	}

	return string(msg)
}

func ircMessage(message string) IRCMessage {
	matches := re.FindStringSubmatch(message)
	tags := processTags(matches[1])
	login := strings.Split(matches[3], "!")[0]
	var nick string

	if v, ok := tags["display-name"]; ok {
		nick = v
	} else {
		nick = login
	}

	result := IRCMessage{
		Login:   login,
		Nick:    nick,
		Type:    matches[5],
		Channel: strings.Replace(matches[6], "#", "", 1),
		Text:    matches[8],
		Tags:    tags,
		Prefix:  matches[3],
	}

	//fmt.Println("IRC MESSAGE:", message)
	return result
}

func IrcPing() {
	for {
		if lastDataReceived.Add(SocketPing*time.Second).Before(time.Now()) && isIrcConnected {
			log.Println("Send PING")
			lastDataReceived = time.Now()

			if _, err := fmt.Fprintln(socket, "PING :tmi.twitch.tv"); err != nil {
				log.Println(err)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func Connect(out chan any) {
	var err error
	if socket, err = net.Dial("tcp", fmt.Sprintf("%s:%d", "irc.chat.twitch.tv", 6667)); err != nil {
		log.Println("Service TWITCH connection error: ", err)
	} else {
		log.Println("Connecting to twitch chat server successfull")
	}

	socket.SetReadDeadline(time.Now().Add(time.Second * SocketWait))

	fmt.Fprintln(socket, "CAP REQ :twitch.tv/commands twitch.tv/tags twitch.tv/membership")
	fmt.Fprintln(socket, fmt.Sprintf("PASS %s", "oauth:gex3upkd5yplcnpe0o0s0msls6781j"))
	fmt.Fprintln(socket, fmt.Sprintf("NICK %s", "YudoleBot"))
	fmt.Fprintln(socket, fmt.Sprintf("JOIN %s", "#ewolf34"))
	fmt.Fprintln(socket, fmt.Sprintf("JOIN %s", "#digitalcorp"))
	fmt.Fprintln(socket, fmt.Sprintf("JOIN %s", "#trossovich"))
	fmt.Fprintln(socket, fmt.Sprintf("JOIN %s", "#m_on_t"))

	isIrcConnected = true

	scanner := bufio.NewScanner(bufio.NewReader(socket))
	socket.SetReadDeadline(time.Now().Add(time.Second * SocketWait))

	for scanner.Scan() {
		lastDataReceived = time.Now()
		socket.SetReadDeadline(time.Now().Add(time.Second * SocketWait))
		msg := ircMessage(scanner.Text())

		fmt.Println(msg)

		switch strings.ToLower(msg.Type) {
		case "join":
			break

		case "part":

			break

		case "privmsg":
			fmt.Printf("INCOMING MESSAGE: [%s]: %s\n", msg.Nick, msg.Text)

			out <- obj.ChatMessage{
				Type:    "chat/message",
				Service: "twitch",
				Src:     msg.Text,
				Text:    "",
				Html:    processSmiles(msg),
				User: obj.User{
					Name: msg.Nick,
					Meta: obj.UserMeta{
						Avatar: "",
						Badges: msg.Tags,
					},
				},
			}
			break

		case "ping":
			log.Println("Received PING")
			if _, err := fmt.Fprintln(socket, "PONG :"+msg.Text); err != nil {
				log.Println("Service TWITCH connection error:", err)
			}
			break

		case "pong":
			log.Println("Received PONG")
			break

		case "roomstate":
			break

		case "userstate":
			break

		case "globaluserstate":
			break

		case "001", "002", "003", "004", "353", "366", "372", "375", "376", "cap":
			// Ignore this message types
			break

		default:
			log.Println("Service TWITCH unknown irc message type:", msg.Type)
		}
	}

	// TODO Reconnecting message send

	isIrcConnected = false
	defer Connect(out)
	defer socket.Close()
}
