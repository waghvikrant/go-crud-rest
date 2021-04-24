package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"sample-backend/model"
	"sample-backend/persistence"
	"strconv"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var e model.Employee
	body, err := ioutil.ReadAll(io.Reader(r.Body))

	if err != nil {
		logger.Error(err)
	}
	if err := r.Body.Close(); err != nil {
		logger.Error(err)
	}
	if err := json.Unmarshal(body, &e); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			logger.Error(err)
		}
	}

	err = persistence.Add(e)

	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ListEmployees(w http.ResponseWriter, r *http.Request) {
	e, err := persistence.List()

	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		logger.Error(err)
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["eid"])

	if err != nil {
		logger.Error(err)
	}
	e, err := persistence.Get(uint64(id))

	if err != nil {
		logger.Error(err)
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		logger.Error(err)
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var e model.Employee

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["eid"])

	if err != nil {
		logger.Error(err)
	}
	e, err = persistence.Get(uint64(id))
	if err != nil {
		logger.Error(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	body, err := ioutil.ReadAll(io.Reader(r.Body))

	if err != nil {
		logger.Error(err)
	}
	if err := r.Body.Close(); err != nil {
		logger.Error(err)
	}
	if err := json.Unmarshal(body, &e); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			logger.Error(err)
		}
	}
	e.ID = uint64(id)
	err = persistence.Update(e)

	if err != nil {
		logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["eid"])

	if err != nil {
		logger.Error(err)
	}
	err = persistence.Delete(uint64(id))

	if err != nil {
		logger.Error(err)
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}
