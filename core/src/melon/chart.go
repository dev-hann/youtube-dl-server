package melon

type Chart struct {
	Top     []*Sing `json:"top,omitempty"`
	Ballade []*Sing `json:"ballade,omitempty"`
	Dance   []*Sing `json:"dance,omitempty"`
	Hiphop  []*Sing `json:"hiphop,omitempty"`
	Rnb     []*Sing `json:"rnb,omitempty"`
	Indie   []*Sing `json:"indie,omitempty"`
	Rock    []*Sing `json:"rock,omitempty"`
	Trot    []*Sing `json:"trot,omitempty"`
	Folk    []*Sing `json:"folk,omitempty"`
}
