package reusable

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDefaultMongo(uri string, db string) error {
	return mgm.SetDefaultConfig(nil, db, options.Client().ApplyURI(uri))
}
