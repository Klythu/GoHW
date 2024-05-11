package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Service struct {
	data map[int]*User
}

type User struct {
	Name        string
	Age         int
	Friend_list []int
}

func (s *Service) refresh() {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		fmt.Print("Data was stolen")
	}
	var buff *map[int]*User
	err = json.Unmarshal(data, buff)
	if err != nil {
		fmt.Print("unredable")
	}
	s.data = *buff

}
func (s *Service) save() {
	buff, err := json.Marshal(s.data)
	if err != nil {
		fmt.Print("unwritable")
	}
	err = ioutil.WriteFile("./reserve/data.json", buff, os.FileMode(0644))
	if err != nil {
		fmt.Print("Data was NOT SAVED")
	}
}

type Deal struct {
	target string
	sourse string
}
type Dead struct {
	grave string
}

func main() {

	alogada := http.NewServeMux()
	srw := Service{map[int]*User{}}
	srw.refresh()
	alogada.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("alogada"))
	})
	alogada.HandleFunc("/create", srw.Create)
	alogada.HandleFunc("/make_friends", srw.Make_friends)
	alogada.HandleFunc("/user", srw.Delete)
	alogada.HandleFunc("/froends/", srw.Who)
	alogada.HandleFunc("/", srw.Older)
	http.ListenAndServe("localhost:8082", alogada)
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	s.refresh()
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var u *User
		if err := json.Unmarshal(content, u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		s.data[len(s.data)] = u
		s.save()
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Add User:" + u.Name))
	}
}
func (s *Service) Make_friends(w http.ResponseWriter, r *http.Request) {
	s.refresh()
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var deal *Deal
		if err := json.Unmarshal(content, deal); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		t, err := strconv.Atoi(deal.target)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		so, err := strconv.Atoi(deal.sourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.data[t].Friend_list[len(s.data[t].Friend_list)] = so
		s.data[so].Friend_list[len(s.data[so].Friend_list)] = t
		s.save()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Friends:" + s.data[t].Name + " " + s.data[so].Name))
	}
}
func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	s.refresh()
	if r.Method == "DELETE" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var dead *Dead
		if err := json.Unmarshal(content, dead); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		point, err := strconv.Atoi(dead.grave)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		for _, u := range s.data {
			for i, v := range u.Friend_list {
				if v == point {
					u.Friend_list = append(u.Friend_list[:i], u.Friend_list[i+1:]...)
				}
			}
		}
		u := s.data[point]
		delete(s.data, point)
		s.save()
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("No more User:" + u.Name))
	}
}
func (s *Service) Who(w http.ResponseWriter, r *http.Request) {
	s.refresh()
	if r.Method == "GET" {
		point := strings.Trim(r.URL.Path, "/froends/")
		search, err := strconv.Atoi(point)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		for i := range s.data[search].Friend_list {
			w.Write([]byte(s.data[i].Name))
		}
	}
}
func (s *Service) Older(w http.ResponseWriter, r *http.Request) {
	s.refresh()
	if r.Method == "PUT" {
		point := strings.Trim(r.URL.Path, "/")
		search, err := strconv.Atoi(point)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		var dead *Dead
		if err := json.Unmarshal(content, dead); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		s.data[search].Age, err = strconv.Atoi(dead.grave)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		s.save()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("возраст пользователя успешно обновлён"))
	}
}
