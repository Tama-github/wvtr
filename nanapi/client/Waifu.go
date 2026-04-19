package client

type CharachterAL struct {
	IdAl              int    `json:"id_al"`
	NameUserPreferred string `json:"name_user_preferred"`
	ImageLarge        string `json:"image_large"`
	Rank              string `json:"rank"`
}

type Waifu struct {
	ID         string        `json:"id"`
	Charachter *CharachterAL `json:"character"`
}

type JoinWC struct {
	ID                string `json:"id"`
	IdAl              int    `json:"id_al"`
	NameUserPreferred string `json:"name_user_preferred"`
	ImageLarge        string `json:"image_large"`
	Rank              string `json:"rank"`
}
