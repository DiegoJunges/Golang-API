package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, nil
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorID)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (repository Posts) FindById(postID uint64) (models.Post, error) {
	lines, err := repository.db.Query(`
		SELECT p.*, u.nickname FROM 
		posts p inner join users u
		on u.id = p.author_id where p.id = ?`,
		postID,
	)
	if err != nil {
		return models.Post{}, err
	}
	defer lines.Close()

	var post models.Post

	if lines.Next() {
		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository Posts) Find(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		SELECT distinct p.*, u.nickname FROM posts p 
		inner join users u on u.id = p.author_id
		inner join followers f on p.author_id = f.user_id
		where u.id = ? or f.follower_id = ?
		order by 1 desc`,
		userID, userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Posts) Update(postID uint64, post models.Post) error {
	statement, err := repository.db.Prepare(
		"UPDATE posts SET title = ?, content = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(post.Title, post.Content, postID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Posts) Delete(postID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM posts WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(postID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Posts) FindByUser(userID uint64) ([]models.Post, error) {
	lines, err := repository.db.Query(`
		SELECT p.*, u.nickname FROM posts p 
		join users u on u.id = p.author_id
		where p.author_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var posts []models.Post

	for lines.Next() {
		var post models.Post

		if err = lines.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNickname,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
func (repository Posts) Like(postID uint64) error {
	statement, err := repository.db.Prepare(
		"UPDATE posts SET likes = likes + 1 WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(postID)
	if err != nil {
		return err
	}

	return nil
}

func (repository Posts) Dislike(postID uint64) error {
	statement, err := repository.db.Prepare(`
	UPDATE posts SET likes =
	CASE 
		WHEN likes > 0 THEN likes - 1 
		ELSE 0 AND	
		WHERE id = ?`,
	)
	if err != nil {
		return err
	}

	_, err = statement.Exec(postID)
	if err != nil {
		return err
	}

	return nil
}
