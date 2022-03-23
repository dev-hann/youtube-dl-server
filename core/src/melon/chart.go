package melon

type Chart struct {
	Title    string  `json:"title"`
	SubTitle string  `json:"sub_title"`
	ItemList []*Sing `json:"item_list"`
}
