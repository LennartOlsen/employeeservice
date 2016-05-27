package employee

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

type Controller struct {
	Repos *Repository
}

func (ctrl *Controller) GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx := context.Background()

	/** TODO : CREATE A QUERY STRING LANGUGE **/
	ECID := r.URL.Query().Get("employee_category_id")
	models, err := ctrl.Repos.ResolveFromQuery(ECID, ctx);

	if(err != nil){
		ctrl.errorResponse(w, err)
	} else {
		if models == nil {
			ctrl.successResponse(w, "{}")
		} else if b, err := json.Marshal(models); (err != nil){
			ctrl.errorResponse(w, err)
		} else {
			ctrl.successResponse(w, string(b))
		}
	}
}

func (ctrl *Controller) GetById(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx := context.Background()
	id := ps.ByName("id")
	if model, err := ctrl.Repos.GetById(id, ctx); err != nil {
		ctrl.errorResponse(w, err)
	} else {
		if b, err := json.Marshal(model); (err != nil){
			ctrl.errorResponse(w, err)
		} else {
			ctrl.successResponse(w, string(b))
		}
	}
}

func (ctrl *Controller) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx := context.Background()


	model, err := ctrl.Repos.DecodeModel(r.Body)

	if err != nil {
		ctrl.errorResponse(w, err)
	} else {
		if model, err := ctrl.Repos.Create(model, ctx); (err != nil) {
			ctrl.errorResponse(w, err)
		} else {
			if b, err := json.Marshal(model); (err != nil){
				ctrl.errorResponse(w, err)
			} else {
				ctrl.successResponse(w, string(b))
			}
		}
	}
}

func (ctrl *Controller) Put(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx := context.Background()

	model, err := ctrl.Repos.DecodeModel(r.Body)

	if err != nil {
		ctrl.errorResponse(w, err)
	} else {
		model.Id = ps.ByName("id");
		if model, err := ctrl.Repos.Put(model, ctx); (err != nil) {
			ctrl.errorResponse(w, err)
		} else {
			if b, err := json.Marshal(model); (err != nil){
				ctrl.errorResponse(w, err)
			} else {
				ctrl.successResponse(w, string(b))
			}
		}
	}
}

func (ctrl *Controller) successResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", response)
}

func (ctrl *Controller) errorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(500)

	if json, errJson := json.Marshal(err); (err != nil) {
		fmt.Fprintf(w, "%s", err)
		fmt.Fprintf(w, "%s", errJson)
	} else {
		fmt.Fprintf(w, "%s", json)
	}
}