package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	team "github.com/k4sper1love/go-2024/pkg/cloud9"
)

type Response struct {
	Members []team.Member `json:"members"`
}

func healthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Author: Sergey Yelkin 22B030534\nThis is an app for obtaining information about the members of the CLOUD9 esports team in the Counter Strike discipline in JSON format")
}

func members(w http.ResponseWriter, r *http.Request){
	members := team.GetMembers()

	var response Response
	response.Members = members

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.Members)
}

func membersByType(w http.ResponseWriter, r *http.Request) {
	memberType := mux.Vars(r)["type"]
	members := team.GetMembers()

	var response Response
	for _, member := range members {
		if member.Type == memberType {
			response.Members = append(response.Members, member)
		}
	}

	if response.Members != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response.Members)
		return
	}
	http.NotFound(w, r)
}

func memberByNickname(w http.ResponseWriter, r *http.Request){
	nickname := mux.Vars(r)["nickname"]
	members := team.GetMembers()

	for _, member := range members {
		if member.Nickname == nickname{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(member)
			return
		}
	}
	http.NotFound(w, r)
}

func memberByName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	firstName, secondName := vars["firstName"], vars["secondName"]
	members := team.GetMembers()

	for _, member := range members {
		if member.FirstName == firstName && member.SecondName == secondName {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(member)
			return
		}
	}
	http.NotFound(w, r)
}