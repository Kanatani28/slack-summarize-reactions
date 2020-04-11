package structs

// UsersJSON users.listで取得できるJSON
type UsersJSON struct {
	Ok      bool   `json: "ok"`
	Members []User `json: "members"`
}

// User users.listで取得できるJSONのmembersフィールドの要素
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	RealName string `json:"real_name"`
}
