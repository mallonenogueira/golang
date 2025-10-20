package repositories

import "go-api-crud/models"

type UserRepository struct {
	users map[string]models.User
}

func NewMemoryUserRepository() *UserRepository {
	users := make(map[string]models.User)

	user1, _ := models.NewUser("João Silva", "joao@example.com")
	user2, _ := models.NewUser("Maria Santos", "maria@example.com")

	users[user1.Id] = *user1
	users[user2.Id] = *user2

	return &UserRepository{
		users: users,
	}
}

func (u *UserRepository) GetAll() []models.User {
	users := make([]models.User, 0, len(u.users))

	for _, user := range u.users {
		users = append(users, user)
	}

	return users
}

func (u *UserRepository) GetById(id string) (models.User, bool) {
	user, exists := u.users[id]

	return user, exists
}

func (u *UserRepository) Update(user *models.User) bool {
	_, exists := u.users[user.Id]

	if !exists {
		return false
	}

	u.users[user.Id] = *user

	return true
}

func (u *UserRepository) Insert(user *models.User) {
	u.users[user.Id] = *user
}

func (u *UserRepository) Delete(id string) bool {
	_, exists := u.users[id]

	if !exists {
		return false
	}

	delete(u.users, id)

	return true
}
