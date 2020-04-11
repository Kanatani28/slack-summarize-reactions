package structs

type ChannelMsgsJSON struct {
	Ok       bool             `json:"ok"`
	Messages []ChannelMessage `json:"messages"`
}

type ChannelMessage struct {
	Text      string     `json:"text"`
	User      string     `json:"user"`
	Reactions []Reaction `json:"reactions"`
}

type Reaction struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
	Count int      `json:"count"`
}
