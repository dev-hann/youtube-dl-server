package melon

type Top struct {
	Title    string  `json:"title"`
	SubTitle string  `json:"sub_title"`
	ItemList []*Sing `json:"item_list"`
}
