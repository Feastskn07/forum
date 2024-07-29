package connection

import (
	"database/sql"
	"fmt"
	"forum/helpers"
	auth "forum/session"
	"net/http"
	"strings"
)

func GetThreePosts(db *sql.DB) ([]helpers.Posts, error) {
	query := `select p.id, p.title, p.content, p.author_name, p.categories, p.created_at, 
			  p.img_url, p.blog_url, u.avatar_url
			  from posts p
			  inner join users u on p.author_name = u.username
			  order by datetime(p.created_at) desc
			  limit 4`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []helpers.Posts
	for rows.Next() {
		var post helpers.Posts
		var categories string
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &categories,
			&post.CreatedAt, &post.ImgUrl, &post.BlogUrl, &post.AvatarUrl)
		if err != nil {
			return nil, err
		}
		post.Categories = strings.Split(categories, ", ")
		var categoryUrls []string
		for _, name := range post.Categories {
			categoryUrls = append(categoryUrls, helpers.CreateUrl(name))
		}
		post.CategoryUrl = categoryUrls
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPosts(db *sql.DB) ([]helpers.Posts, error) {
	query := `select p.id, p.title, p.content, p.author_name, p.categories, p.created_at, 
			  p.img_url, p.blog_url, u.avatar_url
			  from posts p
			  inner join users u on p.author_name = u.username`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []helpers.Posts
	for rows.Next() {
		var post helpers.Posts
		var categories string
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &categories,
			&post.CreatedAt, &post.ImgUrl, &post.BlogUrl, &post.AvatarUrl)
		if err != nil {
			return nil, err
		}

		post.Categories = strings.Split(categories, ", ")
		var categoryUrls []string
		for _, name := range post.Categories {
			categoryUrls = append(categoryUrls, helpers.CreateUrl(name))
		}
		post.CategoryUrl = categoryUrls
		posts = append(posts, post)
	}
	return posts, nil
}

func AddBlog(db *sql.DB, title, content string, categories []string, img_url, blog_url string, r *http.Request) error {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return err
	}

	category := strings.Join(categories, ", ")
	sessionToken := cookie.Value
	username := auth.Sessions[sessionToken].Username
	query := `insert into posts (title, content, author_name, categories, img_url, blog_url) 
			  values (?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(query, strings.TrimSpace(title), strings.TrimSpace(content),
		strings.TrimSpace(username), strings.TrimSpace(category),
		strings.TrimSpace(img_url), strings.TrimSpace(blog_url))

	if err != nil {
		return err
	}
	return nil
}

func GetPost(db *sql.DB, id int) (helpers.Posts, error) {
	query := `select id, title, content, author_name, categories, created_at, img_url, 
			  blog_url
			  from posts where id = ?`
	row := db.QueryRow(query, id)

	var post helpers.Posts
	var categories string
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Author, &categories,
		&post.CreatedAt, &post.ImgUrl, &post.BlogUrl)

	if err != nil {
		if err == sql.ErrNoRows {
			return helpers.Posts{}, nil
		}
		return helpers.Posts{}, err
	}
	post.Categories = strings.Split(categories, ", ")
	return post, nil
}

func DeletePost(db *sql.DB, post helpers.Posts) error {
	query := `delete from posts where id = ?`
	_, err := db.Exec(query, post.ID)
	return err
}

func UpdatePost(db *sql.DB, id int, post helpers.Posts) error {
	categories := strings.Join(post.Categories, ", ")
	query := `update posts set title = ?, blog_url = ?, content = ?, img_url = ?, 
			  categories = ? where id = ?`

	_, err := db.Exec(query, strings.TrimSpace(post.Title), strings.TrimSpace(post.BlogUrl),
		strings.TrimSpace(post.Content), strings.TrimSpace(post.ImgUrl),
		strings.TrimSpace(categories), id)

	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}
	return nil
}

func GetCategoryNamesByID(db *sql.DB, ids []string) ([]string, error) {
	var categoryNames []string
	for _, id := range ids {
		var name string
		err := db.QueryRow("select category from categories where id = ?",
			strings.TrimSpace(id)).Scan(&name)

		if err != nil {
			return nil, err
		}

		categoryNames = append(categoryNames, name)
	}
	return categoryNames, nil
}
