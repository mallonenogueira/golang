package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Server struct {
	users map[string]User
	mu    sync.RWMutex
}

func NewServer() *Server {
	return &Server{
		users: make(map[string]User),
	}
}

// GET /users - Lista todos os usuários
// GET /users/{id} - Busca um usuário específico
func (s *Server) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s.getUsers(w, r)
	case http.MethodPost:
		s.createUser(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

// GET /users ou /users/{id}
func (s *Server) getUsers(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/users")
	path = strings.TrimPrefix(path, "/")

	// Se tem ID na URL, busca usuário específico
	if path != "" {
		s.getUserByID(w, r, path)
		return
	}

	// Lista todos os usuários
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GET /users/{id}
func (s *Server) getUserByID(w http.ResponseWriter, r *http.Request, id string) {
	s.mu.RLock()
	user, exists := s.users[id]
	s.mu.RUnlock()

	if !exists {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// POST /users
func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if user.ID == "" || user.Name == "" {
		http.Error(w, "ID e Name são obrigatórios", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.users[user.ID] = user
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// PUT /users/{id}
func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id := path

	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	if _, exists := s.users[id]; !exists {
		s.mu.Unlock()
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	user.ID = id
	s.users[id] = user
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DELETE /users/{id}
func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id := path

	if id == "" {
		http.Error(w, "ID é obrigatório", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	if _, exists := s.users[id]; !exists {
		s.mu.Unlock()
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	delete(s.users, id)
	s.mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) router(w http.ResponseWriter, r *http.Request) {
	// Roteamento manual
	if strings.HasPrefix(r.URL.Path, "/users") {
		if r.Method == http.MethodPut {
			s.updateUser(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			s.deleteUser(w, r)
			return
		}
		s.handleUsers(w, r)
		return
	}

	http.NotFound(w, r)
}

func main() {
	server := NewServer()

	server.users["1"] = User{ID: "1", Name: "João Silva", Email: "joao@example.com"}
	server.users["2"] = User{ID: "2", Name: "Maria Santos", Email: "maria@example.com"}

	http.HandleFunc("/", server.router)

	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
