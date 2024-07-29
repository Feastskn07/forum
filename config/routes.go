package config

import "net/http"

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Index)

	mux.HandleFunc("/about-us", handlers.About)

	mux.HandleFunc("/blog", handlers.Blog)
	mux.HandleFunc("/blogs/", handlers.Blog)
	mux.HandleFunc("/blog-content/", handlers.BlogContent)

	mux.HandleFunc("/forum", handlers.Forum)
	mux.HandleFunc("/new-entry", handlers.NewEntry)
	mux.HandleFunc("/entry-comment", handlers.Comment)
	mux.HandleFunc("/forum-entry/", handlers.ForumEntry)
	mux.HandleFunc("/posts", handlers.Posts)
	mux.HandleFunc("/addcomment", handlers.AddComment)
	mux.HandleFunc("/category/", handlers.Category)

	mux.HandleFunc("/profile/", handlers.Profile)
	mux.HandleFunc("/profile-settings", handlers.ProfileSettings)
	mux.HandleFunc("/upload-photo", handlers.UploadPhoto)
	mux.HandleFunc("/update-user-info", handlers.UpdateUserInfo)
	mux.HandleFunc("/user/posts", handlers.UserEntries)
	mux.HandleFunc("/user/comments", handlers.UserComments)
	mux.HandleFunc("/user/likes", handlers.UserLikes)
	mux.HandleFunc("/user/dislikes", handlers.UserDislikes)

	mux.HandleFunc("/login", handlers.Userops{}.Login)
	mux.HandleFunc("/register", handlers.Userops{}.Register)
	mux.HandleFunc("/logout", handlers.Userops{}.Logout)

	mux.HandleFunc("/admin", admin.BlogDashboard{}.Index)
	mux.HandleFunc("/admin/blog-new-item", admin.BlogDashboard{}.NewItem)
	mux.HandleFunc("/admin/blog-add", admin.BlogDashboard{}.Add)
	mux.HandleFunc("/admin/blog-delete/", admin.BlogDashboard{}.Delete)
	mux.HandleFunc("/admin/blog-edit/", admin.BlogDashboard{}.Edit)
	mux.HandleFunc("/admin/blog-update/", admin.BlogDashboard{}.Update)

	mux.HandleFunc("/admin/blog-categories", admin.BlogCategories{}.Index)
	mux.HandleFunc("/admin/blog-categories/blog-add", admin.BlogCategories{}.Add)
	mux.HandleFunc("/admin/blog-categories/blog-delete/", admin.BlogCategories{}.Delete)

	mux.HandleFunc("/admin/form", admin.FormDashboard{}.Index)
	mux.HandleFunc("/admin/form-new-item", admin.FormDashboard{}.NewItem)
	mux.HandleFunc("/admin/form-add", admin.FormDashboard{}.Add)
	mux.HandleFunc("/admin/form-delete/", admin.FormDashboard{}.Delete)
	mux.HandleFunc("/admin/form-edit/", admin.FormDashboard{}.Edit)
	mux.HandleFunc("/admin/form-update/", admin.FormDashboard{}.Update)

	mux.HandleFunc("/admin/form-categories", admin.FormCategories{}.Index)
	mux.HandleFunc("/admin/form-categories/form-add", admin.FormCategories{}.Add)
	mux.HandleFunc("/admin/form-categories/form-delete/", admin.FormCategories{}.Delete)

	mux.HandleFunc("/entry-img/", handlers.EntryImg)
	mux.HandleFunc("/avatars/", handlers.Avatars)
	mux.HandleFunc("/static/", handlers.Static)
	mux.HandleFunc("/uploads/", handlers.Uploads)
	mux.HandleFunc("/profile-photos/", handlers.ProfilePhotos)
	mux.HandleFunc("/internal/admin/assets/", handlers.Internal)

	return mux
}
