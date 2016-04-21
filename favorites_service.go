package favorites

import (
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"github.com/mjibson/goon"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

type FavoritesService struct{}

// GET http://localhost:8080/posts/ahdkZXZ-ZmVkZXJhdGlvbi1zZXJ2aWNlc3IVCxIIcHJvZmlsZXMYgICAgICAgAoM
func (ps *FavoritesService) FavsFindRecord(c context.Context, r *FavoriteFindRequest) (*FavoriteSingle, error) {
	g := goon.FromContext(c)
	p := &Person{Name: r.Name}
	if err := g.Get(p); err == datastore.ErrNoSuchEntity {
		return nil, endpoints.NewNotFoundError("Name not found: %s", r.Name)
	} else if err != nil {
		return nil, endpoints.NewInternalServerError("Datastore: %v", err)
	}

	f := []*Favorite{}
	q := datastore.NewQuery("Favorite").Ancestor(g.Key(p)).Order("Order")
	if _, err := g.GetAll(q, &f); err == datastore.ErrNoSuchEntity {
		log.Infof(c, "No Favorites for: %s", r.Name)
	} else if err != nil {
		return nil, endpoints.NewInternalServerError("Datastore: %v", err)
	}

	return &FavoriteSingle{
		Person:    p,
		Favorites: f,
	}, nil
}

func (ps *FavoritesService) FavsUpdate(c context.Context, r *FavoriteSingle) (*FavoriteSingle, error) {
	g := goon.FromContext(c)
	k, err := g.Put(r.Person)
	if err != nil {
		return nil, endpoints.NewInternalServerError("Datastore Put Person: %v", err)
	}

	for i, f := range r.Favorites {
		f.Order = int64(i + 1)
		f.Parent = k
	}
	if _, err := g.PutMulti(&r.Favorites); err != nil {
		return nil, endpoints.NewInternalServerError("Datastore Put Favorites: %v", err)
	}

	return r, nil
}

func (ps *FavoritesService) FavsDelete(c context.Context, r *FavoriteFindRequest) (*FavoritesService, error) {
	g := goon.FromContext(c)
	q := datastore.NewQuery("").Ancestor(g.Key(&Person{Name: r.Name})).KeysOnly()
	k, err := g.GetAll(q, nil)
	if err != nil {
		return nil, endpoints.NewInternalServerError("Delete Favs: %v", err)
	}

	if err := g.DeleteMulti(k); err != nil {
		return nil, endpoints.NewInternalServerError("Delete Favs: %v", err)
	}

	return &FavoritesService{}, nil
}
