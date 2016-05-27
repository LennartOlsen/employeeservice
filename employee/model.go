package employee

import (
	"time"
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
)

const DatastoreEntity string = "Employee"

type Model struct {
	datastoreKey int64 `json:"-" datastore:"-"`
	Id string `json:"id" datastore:"id"`
	Name string `json:"name" datastore:"name"`
	Email string `json:"email" datastore:"email"`
	Image string `json:"image_id" datastore:"image_id"`
	EmployeeCategory string `json:"employee_category_id" datastore:"employee_category_id"`
	Title string `json:"title" datastore:"title"`
	Created time.Time `json:"created" datastore:"created"`
	Updated time.Time `json:"updated" datastore:"updated"`
	Deleted time.Time `json:"deleted" datastore:"deleted"`
}

func (model *Model) key(c context.Context) *datastore.Key {
	// if there is no Id, we want to generate an "incomplete"
	// one and let datastore determine the key/Id for us
	if model.datastoreKey == 0 {
		return datastore.NewIncompleteKey(c, DatastoreEntity, nil)
	}

	// if Id is already set, we'll just build the Key based
	// on the one provided.
	return datastore.NewKey(c, DatastoreEntity, "", model.datastoreKey, nil)
}