package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CezarGarrido/FarmVue/ApiFarm/driver"
	entities "github.com/CezarGarrido/FarmVue/ApiFarm/entities"
	repo "github.com/CezarGarrido/FarmVue/ApiFarm/repo"
	"github.com/gorilla/mux"
)

// NewLogHandler ...
func NewAreaHandler(db *driver.DB) *Area {
	return &Area{
		repo: repo.NewSQLAreaRepo(db.SQL),
	}
}

// Log...
type Area struct {
	repo repo.AreaRepo
}

func (p *Area) Create(w http.ResponseWriter, r *http.Request) {
	area := entities.Area{}
	json.NewDecoder(r.Body).Decode(&area)
	newID, err := p.repo.Create(r.Context(), &area)
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	respondwithJSON(w, http.StatusCreated, map[string]interface{}{"message": "Successfully Created", "id": newID})
}
func (p *Area) GetAllByProprietarioID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	fmt.Println(id)
	payload, err := p.repo.GetAllByProprietarioID(r.Context(), int64(id))
	if err != nil {
		fmt.Printf("%+v\n", err.Error())
		respondWithError(w, http.StatusNoContent, "Content not found")
		return
	}
	respondwithJSON(w, http.StatusOK, payload)
}
// Delete a post
func (p *Area) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		fmt.Printf("%+v\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}
	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}
// Fetch all log data
/*func (p *Area) Fetch(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
		fmt.Println("teste")
		fmt.Fprint(os.Stderr, "request cancelled\n")
	case <-time.After(1 * time.Second):
		payload, err := p.repo.Fetch(ctx, 100)
		if err != nil {
			respondWithError(w, 500, err.Error())
			return
		}
		respondwithJSON(w, http.StatusOK, payload)
	}
}*/
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
