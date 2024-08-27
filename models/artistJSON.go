package models

type Artists struct {
	Artist []Artist
}

type Artist struct {
	ID           int64     `json:"id"`
	Image        string    `json:"image"`
	Name         string    `json:"name"`
	Members      []string  `json:"members`
	CreationDate int64     `json:"creationDate"`
	FirstAlbum   string    `json:"firstAlbum"`
	Locations    string    `json:"locations"`
	ConcertDates string    `json:"concertDates"`
	Relations    string    `json:"relations"`
}

type ArtistTemplate struct {
	Artist       *Artist
	Locations    Locations
	Relations    Relations
	ConcertDates ConcertDates
}

type ConcertDates struct {
	ID       int64 `json:"id"`
	ConcertDates []string `json:"concertDates"`
}

type Locations struct {
	ID        int64 `json:"id"`
	Locations []string `json:"locations"`
}

type Relations struct {
	ID        int64 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type  Pagedata struct {
	Result string
}