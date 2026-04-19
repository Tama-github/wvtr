package client

type Waifu struct {
	ID                string `json:"id"`
	IdAl              uint   `json:"id_al"`
	NameUserPreferred string `json:"name_user_preferred"`
	ImageLarge        string `json:"image_large"`
	Rank              string `json:"rank"`
	// maybe later
	//CustomImage       string `json:"custom_image"`
}
