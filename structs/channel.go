package structs

// ChannelsJSON channels.listで取得できるJSON
type ChannelsJSON struct {
	Ok       bool      `json:"ok"`
	Channels []Channel `json:"channels"`
}

// Channel channels.listで取得できるJSONのchannelsフィールドの要素
type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
