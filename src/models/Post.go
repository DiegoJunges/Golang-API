package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID             uint64    `json:"id"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorID       uint64    `json:"author_id,omitempty"`
	AuthorNickname string    `json:"author_nickname,omitempty"`
	Likes          uint64    `json:"likes,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("Title is needed")
	}

	if post.Content == "" {
		return errors.New("Content is needed")
	}
	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
