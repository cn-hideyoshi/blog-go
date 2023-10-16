package dao

import (
	"blog-go/models"
	"log"
)

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPage 查询出错:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostSearch(condition string) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post where title like ?", "%"+condition+"%")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPage 查询出错:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where slug = ? limit ?,?", slug, page, pageSize)
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPage 查询出错:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CountGetAllPost() (count int) {
	row := DB.QueryRow("select count(1) from blog_post")
	_ = row.Scan(&count)
	return
}

func CountGetAllPostBySlug(slug string) (count int) {
	row := DB.QueryRow("select count(1) from blog_post where slug = ?", slug)
	_ = row.Scan(&count)
	return
}

func GetPostPageByCId(cId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cId, page, pageSize)
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPage 查询出错:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetAllPost() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPage 查询出错:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CountGetPostByCId(cId int) (count int) {
	row := DB.QueryRow("select count(1) from blog_post where category_id = ?", cId)
	_ = row.Scan(&count)
	return
}

func GetPostById(pId int) (*models.Post, error) {
	p := &models.Post{}
	err := DB.QueryOne(p, "select * from blog_post where pid=?", pId)
	return p, err
	//row := DB.QueryRow("select * from blog_post where pid = ? limit 1", pId)
	//var post = &models.Post{}
	//if row.Err() != nil {
	//	return nil, row.Err()
	//}
	//err := row.Scan(&post.Pid,
	//	&post.Title,
	//	&post.Content,
	//	&post.Markdown,
	//	&post.CategoryId,
	//	&post.UserId,
	//	&post.ViewCount,
	//	&post.Type,
	//	&post.Slug,
	//	&post.CreateAt,
	//	&post.UpdateAt,
	//)
	//if err != nil {
	//	log.Println("GetPostPage 查询出错:", err)
	//	return post, err
	//}
	//return post, nil
}

func SavePost(post *models.Post) {
	exec, err := DB.Exec("insert into blog_post (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at)"+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
		return
	}

	lastInsertId, err := exec.LastInsertId()
	post.Pid = int(lastInsertId)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title =?,content=?,markdown=?,category_id=?,user_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.Type, post.Slug, post.UpdateAt, post.Pid)
	if err != nil {
		log.Println(err)
		return
	}
}
