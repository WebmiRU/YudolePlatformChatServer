package twitch_client

type IRCMessage struct {
	Login   string
	Nick    string
	Tags    map[string]string
	Text    string
	Channel string
	Type    string
	Prefix  string
}
