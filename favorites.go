package favorites

import (
	"google.golang.org/appengine/datastore"
)

type Person struct {
	Name     string `json:"name" xml:"name" datastore:"-" goon:"id"`
	Password []byte `json:"password" xml:"password"`
}

type Favorite struct {
	Parent *datastore.Key `json:"-" xml:"-" datastore:"-" goon:"parent"`
	Order  int64          `json:"-" xml:"-" goon:"id"`
	Type   string         `json:"type" xml:"type"`
	Value  string         `json:"value" xml:"value"`
}

type FavoriteSingle struct {
	Person    *Person     `json:"person" xml:"person"`
	Favorites []*Favorite `json:"favorite" xml:"favorite"`
}

type FavoriteFindRequest struct {
	Name string `json:"name" xml:"name"`
}
