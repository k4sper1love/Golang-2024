package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
type Response struct {
	Members []Member
}
type Member struct {
	Type string `json:"type"`
	Role string `json:"role"`
	Nickname string `json:"nickname"`
	FirstName string `json:"firstname"`
	SecondName string `json:"secondname"`
}

func prepareResponse() []Member{
	var members []Member
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "HObbit", FirstName: "Abay", SecondName: "Khassenov"})
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "Ax1Le", FirstName: "Sergey", SecondName: "Rykhtorov"})
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "electroNic", FirstName: "Denis", SecondName: "Sharipov"})
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "Perfecto", FirstName: "Ilia", SecondName: "Zalutskii"})
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "Boombl4", FirstName: "Kirill", SecondName: "Mikhaylov"})
	members = append(members, Member{Type: "Staff", Role: "Coach", Nickname: "groove", FirstName: "Konstantin", SecondName: "Pikiner"})
	members = append(members, Member{Type: "Staff", Role: "Analyst", Nickname: "F_1N", FirstName: "Ivan", SecondName: "Kochugov"})
	members = append(members, Member{Type: "Staff", Role: "Manager", Nickname: "Sweetypotz", FirstName: "Aleksandr", SecondName: "Shcherbakov"})
	return members
}

func main(){
	router := mux.NewRouter();
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/members", GetMembers).Methods("GET")
	membersRouter := router.PathPrefix("/members").Subrouter()
	membersRouter.HandleFunc("/{nickname}", GetMemberByNickname).Methods("GET")
	membersRouter.HandleFunc("/name/{firstName}-{secondName}", GetMembersByName).Methods("GET")
	membersRouter.HandleFunc("/type/{type}", GetMembersByType).Methods("GET")
	membersRouter.HandleFunc("/role/{role}", GetMembersByRole).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Api is running")
}

func GetMembers(w http.ResponseWriter, r *http.Request){
	members := prepareResponse()
	var response Response
	response.Members = members
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.Members)
}

func GetMemberByNickname(w http.ResponseWriter, r *http.Request){
	nickname := mux.Vars(r)["nickname"];
	members := prepareResponse();
	for _, member := range members {
		if member.Nickname == nickname{
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(member)

			return
		}
	}
	w.WriteHeader(http.StatusNotFound);
	http.NotFound(w, r)
}

func GetMembersByType(w http.ResponseWriter, r *http.Request) {
	memberType := mux.Vars(r)["type"]
	members := prepareResponse()
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

func GetMembersByName(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	firstName, secondName := vars["firstName"], vars["secondName"]
	members := prepareResponse()
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

func GetMembersByRole(w http.ResponseWriter, r *http.Request){
	role := mux.Vars(r)["role"]
	members := prepareResponse()
	var response Response
	for _, member := range members {
		if member.Role == role {
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