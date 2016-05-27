package employee

import (
	"golang.org/x/net/context"
	"google.golang.org/cloud/datastore"
	"github.com/satori/go.uuid"
	"time"
	"encoding/json"
	"io"
	"log"
)

type Repository struct {
	Client *datastore.Client
}
/** NOT AT ALL THE RIGHT WAY!!!! **/
func (repos *Repository ) ResolveFromQuery (ECID string, ctx context.Context) ([]Model, error){
	if(ECID != ""){
		return repos.GetByEmployeeCategoryId(ECID, ctx);
	} else {
		return repos.GetAll(ctx);
	}
}

func (repos *Repository ) DecodeModel(reader io.Reader) (Model, error) {
	var model Model
	err := json.NewDecoder(reader).Decode(&model)

	return model, err;
}

func (repos *Repository ) GetAll(ctx context.Context) ([]Model, error) {
	var models []Model

	query := datastore.NewQuery(DatastoreEntity).Order("created")

	_, err := repos.Client.GetAll(ctx, query, &models)

	if err != nil {
		return nil, err
	}

	return models, nil
}

func (repos *Repository ) Create(model Model, ctx context.Context) (*Model, error) {

	/*Make sure that the user has a create time*/
	model.Created = time.Now()

	if model.Id == "" {
		v := uuid.NewV4()
		model.Id = v.String();
	}

	return repos.save(&model, ctx);


}

func (repos *Repository ) Put( model Model, ctx context.Context ) (*Model, error) {
	if oldModel, err := repos.GetById(model.Id, ctx); err != nil {
		return nil, err;
	} else {
		model.Created = oldModel.Created
		model.Updated = time.Now()

		return repos.save(&model, ctx);
	}
}
/*func (repos *Repository ) Patch( model Employee, ctx context.Context ) (*Employee, error) {
	model.Updated = time.Now()

	return repos.save(&model, ctx);
}*/

func (repos *Repository ) GetById(id string, ctx context.Context) (*Model, error) {
	models, err := repos.getByKeyValue("id", id, "", ctx)
	if err != nil {
		return nil, err
	}

	return &models[0], nil
}

func (repos *Repository ) GetByEmployeeCategoryId(id string, ctx context.Context) ([]Model, error) {
	return repos.getByKeyValue("employee_category_id", id, "", ctx)
}
func (repos *Repository ) getByKeyValue(key string, value string, comparator string, ctx context.Context) ([]Model, error) {
	var models []Model
	if(comparator == ""){
		comparator = "="
	}
	log.Println(value);
	q := datastore.NewQuery(DatastoreEntity).Filter(key + " " + comparator, value)

	_, err := repos.Client.GetAll(ctx, q, &models)
	log.Println(len(models));
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (repos *Repository) save(model *Model, ctx context.Context) (*Model, error) {
	_, err := repos.Client.Put(ctx, model.key(ctx), model)

	if err != nil {
		return nil, err
	}

	return model, nil
}