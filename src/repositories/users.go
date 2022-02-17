package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) Create(user models.User) (uint64, error) {
	fmt.Println("Creating a new user", user)
	statement, err := u.db.Prepare(
		"INSERT INTO users (name, nickname, email, password) VALUES (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}

func (repository Users) Search(nameOrNickname string) ([]models.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

	lines, err := repository.db.Query(
		"SELECT id, name, nickname, email, password, created_at FROM users WHERE name LIKE ? OR nickname LIKE ?",
		nameOrNickname, nameOrNickname,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) FindById(id uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"SELECT id, name, nickname, email, password, created_at FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nickname,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, nickname = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nickname, user.Email, ID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM users WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Users) FindByEmail(email string) (models.User, error) {
	line, err := repository.db.Query(
		"SELECT id, password FROM users WHERE email = ?",
		email,
	)

	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err := line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repository Users) Follow(userID uint64, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Users) Unfollow(userID uint64, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Users) GetFollowers(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
	SELECT u.id, u.name, u.nickname, u.email, u.created_at 
	FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?
	`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var followers []models.User
	for lines.Next() {
		var follower models.User
		if err := lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nickname,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository Users) GetFollowing(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
	SELECT u.id, u.name, u.nickname, u.email, u.created_at 
	FROM users u INNER JOIN followers f ON u.id = f.user_id WHERE f.follower_id = ?
	`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var following []models.User
	for lines.Next() {
		var follow models.User
		if err := lines.Scan(
			&follow.ID,
			&follow.Name,
			&follow.Nickname,
			&follow.Email,
			&follow.CreatedAt,
		); err != nil {
			return nil, err
		}

		following = append(following, follow)
	}

	return following, nil
}

func (repository Users) GetPassword(userID uint64) (string, error) {
	line, err := repository.db.Query("select password from users where id = ?", userID)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository Users) UpdatePassword(userID uint64, password string) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(password, userID)
	if err != nil {
		return err
	}

	return nil
}
