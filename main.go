package favorites

import (
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"google.golang.org/api/oauth2/v2"
)

const (
	CLIENT_ID_SERVER     = "notdefined.apps.googleusercontent.com"
	CLIENT_SECRET_SERVER = "notdefined"
	CLIENT_ID_IOS        = "CLIENT_ID_IOS"
	CLIENT_ID_ANDROID    = "CLIENT_ID_ANDROID"
)

var (
	Scopes    = []string{endpoints.EmailScope, oauth2.UserinfoProfileScope, oauth2.PlusMeScope}
	ClientIds = []string{CLIENT_ID_SERVER, CLIENT_ID_IOS, endpoints.APIExplorerClientID}
	Audiences = []string{CLIENT_ID_SERVER, CLIENT_ID_IOS}
)

func init() {
	favsApi, err := endpoints.RegisterService(&FavoritesService{}, "favorites", "v1", "Favorites API", false)
	if err != nil {
		log.Fatalf("Register service: %v", err)
	}

	register := func(api *endpoints.RPCService, orig, name, method, path, desc string, auth bool) {
		m := api.MethodByName(orig)
		if m == nil {
			log.Fatalf("Missing method %s", orig)
		}
		i := m.Info()
		i.Name, i.HTTPMethod, i.Path, i.Desc = name, method, path, desc
		if auth {
			i.Scopes, i.ClientIds, i.Audiences = Scopes, ClientIds, Audiences
		}
	}

	register(favsApi, "FavsFindRecord", "favorites.findRecord", "GET", "favorites/{name}", "Find a specific Favorite list.", false)
	register(favsApi, "FavsUpdate", "favorites.update", "PUT", "favorites", "Update a specific Favorite list.", false)
	register(favsApi, "FavsDelete", "favorites.delete", "DELETE", "favorites/{name}", "Delete a specific Favorite list.", false)

	endpoints.HandleHTTP()
}
