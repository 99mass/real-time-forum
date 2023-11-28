package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	"forum/controller"
	"forum/helper"
	"forum/middlewares"
	"forum/models"
)

func FilterMyPage(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ok, errorPage := middlewares.CheckRequest(r, "/filtermypage", "post")
		if !ok {
			helper.ErrorPage(w, errorPage)
			return
		}
		err := r.ParseForm()
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}
		var filterPosts []models.HomeDataPost
		var category []models.Category
		var DataMyPage models.DataMypage
		Categorystring := r.Form["category"]

		Datas, err := helper.GetDataTemplate(db, r, true, false, false, false, true)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		DataMyPage.Category = Datas.Category
		DataMyPage.Session = Datas.Session
		DataMyPage.User = Datas.User
		userID := Datas.User.ID

		PostUsers, err := helper.GetPostsForOneUser(db, userID)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		//Set likes and dislikes
		PostUser, err := helper.SetLikesAndDislikes(Datas.User, PostUsers, db)
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
		}
		//Set Routes
		for i := range PostUser {
			PostUser[i].Route = "mypage"
			//fmt.Println(PostUser[i].Route)
			for j := range PostUser[i].Comment {
				PostUser[i].Comment[j].Route = "mypage"
			}
		}

		if Categorystring != nil {

			for _, v := range Categorystring {
				v = strings.TrimSpace(v)
				var cat models.Category
				catuuid, err := uuid.FromString(v)
				if !helper.VerifCategory(db, catuuid) || err != nil || catuuid == uuid.Nil {
					DataMyPage.Datas = PostUser
					DataMyPage.ErrorFilter = "one of the categories is not compliant"
					helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
					return
				}
				cat.ID = catuuid
				category = append(category, cat)
			}
			postcat, err := controller.GetAllPostCategories(db)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			var idpost []uuid.UUID
			for _, pc := range postcat {
				for _, categ := range category {
					if pc.CategoryID == categ.ID {
						idpost = append(idpost, pc.PostID)
					}
				}
			}
			idpost = RemoveDuplicates(idpost)
			for _, post := range PostUser {
				for _, postid := range idpost {
					if postid == post.Posts.ID {
						filterPosts = append(filterPosts, post)
					}
				}
			}
		}
		if filterPosts == nil && Categorystring == nil {
			filterPosts = PostUser
		}

		date1 := r.FormValue("date1")
		date2 := r.FormValue("date2")
		likemi := r.FormValue("likemin")
		likema := r.FormValue("likemax")
		var likemin, likemax int
		if likemi == "" {
			likemin = 0
		} else {
			likemin, err = strconv.Atoi(likemi)
			if err != nil {
				DataMyPage.Datas = PostUser
				Datas.ErrorFilter = "give us an int"
				helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
				return
			}
		}
		if likema == "" {
			likemax = 1000
		} else {
			likemax, err = strconv.Atoi(likema)
			if err != nil {
				DataMyPage.Datas = PostUser
				DataMyPage.ErrorFilter = "give us an int"
				helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
				return
			}
		}

		if likemax < 0 || likemin < 0 {
			DataMyPage.Datas = PostUser
			DataMyPage.ErrorFilter = "give positive int for filtering by the like"
			helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
			return
		}
		if likemin > likemax {
			DataMyPage.Datas = PostUser
			DataMyPage.ErrorFilter = "the min value can't be over than the max value"
			helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
			return
		}
		if date1 == "" {
			date1 = "2023-08-01"
		}
		if date2 == "" {
			date2 = "2025-08-01"
		}
		date, err := CompareDate(date1, date2)
		if err != nil {
			DataMyPage.Datas = PostUser
			DataMyPage.ErrorFilter = "date format is incorrect"
			helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
			return
		}
		if !date {
			DataMyPage.Datas = PostUser
			DataMyPage.ErrorFilter = "the min value can't be over than the max value"
			helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
			return
		}
		filterPosts, err = GetFilteredPostsMyPage(db, filterPosts, date1, date2)
		if err != nil {
			DataMyPage.Datas = PostUser
			DataMyPage.ErrorFilter = "format date given is incorrect"
			helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
			return
		}

		var PostsFiltered []models.HomeDataPost
		for _, v := range filterPosts {
			if v.PostLike >= likemin && v.PostLike <= likemax {
				PostsFiltered = append(PostsFiltered, v)
			}
		}
		//Set likes and dislikes
		Datasliked, err := helper.SetLikesAndDislikes(Datas.User, PostsFiltered, db)
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
		}
		//Set Routes
		for i := range Datasliked {
			Datasliked[i].Route = "mypage"
			//fmt.Println(Datasliked[i].Route)
			for j := range Datasliked[i].Comment {
				Datasliked[i].Comment[j].Route = "mypage"
			}
		}

		DataMyPage.Datas = Datasliked

		helper.RenderTemplate(w, "mypage", "mypages", DataMyPage)
	}

}

func Filter(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ok, errorPage := middlewares.CheckRequest(r, "/filter", "post")
		if !ok {
			helper.ErrorPage(w, errorPage)
			return
		}
		err := r.ParseForm()
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
			return
		}
		var filterPosts []models.Post
		var category []models.Category
		posts, err := controller.GetAllPosts(db)
		Categorystring := r.Form["category"]
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}

		if Categorystring != nil {

			for _, v := range Categorystring {
				v = strings.TrimSpace(v)
				var cat models.Category
				catuuid, err := uuid.FromString(v)
				if !helper.VerifCategory(db, catuuid) || err != nil || catuuid == uuid.Nil {
					Datas, err := helper.GetDataTemplate(db, r, true, false, true, false, true)
					if err != nil {
						helper.ErrorPage(w, http.StatusInternalServerError)
						return
					}
					Datas.ErrorFilter = "one of the categories is not compliant"
					helper.RenderTemplate(w, "index", "index", Datas)
					return
				}
				cat.ID = catuuid
				category = append(category, cat)
			}
			postcat, err := controller.GetAllPostCategories(db)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			var idpost []uuid.UUID
			for _, pc := range postcat {
				for _, categ := range category {
					if pc.CategoryID == categ.ID {
						idpost = append(idpost, pc.PostID)
					}
				}
			}
			idpost = RemoveDuplicates(idpost)
			for _, post := range posts {
				for _, postid := range idpost {
					if postid == post.ID {
						filterPosts = append(filterPosts, post)
					}
				}
			}

		}

		if filterPosts == nil && Categorystring == nil {
			filterPosts = posts
		}

		date1 := r.FormValue("date1")
		date2 := r.FormValue("date2")
		likemi := r.FormValue("likemin")
		likema := r.FormValue("likemax")
		var likemin, likemax int
		if likemi == "" {
			likemin = 0
		} else {
			likemin, err = strconv.Atoi(likemi)
			if err != nil {
				Datas, err := helper.GetDataTemplate(db, r, true, false, true, false, true)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				Datas.ErrorFilter = "give us an int"
				helper.RenderTemplate(w, "index", "index", Datas)
				return
			}
		}
		if likema == "" {
			likemax = 1000
		} else {
			likemax, err = strconv.Atoi(likema)
			if err != nil {
				Datas, err := helper.GetDataTemplate(db, r, true, false, true, false, true)
				if err != nil {
					helper.ErrorPage(w, http.StatusInternalServerError)
					return
				}
				Datas.ErrorFilter = "give us an int"
				helper.RenderTemplate(w, "index", "index", Datas)
				return
			}
		}
		if likemin > likemax {
			Datas, err := helper.GetDataTemplate(db, r, true, false, true, false, true)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			Datas.ErrorFilter = "the min value can't be over than the max value"
			helper.RenderTemplate(w, "index", "index", Datas)
			return
		}
		if likemax < 0 || likemin < 0 {
			Datas, err := helper.GetDataTemplate(db, r, true, false, true, false, true)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			Datas.ErrorFilter = "give positive int for filtering by the like"
			helper.RenderTemplate(w, "index", "index", Datas)
			return
		}
		if date1 == "" {
			date1 = "2023-08-01"
		}
		if date2 == "" {
			date2 = "2025-08-01"
		}
		date, err := CompareDate(date1, date2)
		if err != nil {
			Datas, err := helper.GetDataTemplate(db, r, true, false, true, false, true)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			Datas.ErrorFilter = "date format is incorrect"
			helper.RenderTemplate(w, "index", "index", Datas)
			return
		}
		if !date {
			Datas, err := helper.GetDataTemplate(db, r, true, false, true, false, true)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			Datas.ErrorFilter = "the min value can't be over than the max value"
			helper.RenderTemplate(w, "index", "index", Datas)
			return
		}
		filterPosts, err = GetFilteredPosts(db, filterPosts, date1, date2)
		if err != nil {
			Datas, err := helper.GetDataTemplate(db, r, true, false, true, false, true)
			if err != nil {
				helper.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			Datas.ErrorFilter = "format date given is incorrect"
			helper.RenderTemplate(w, "index", "index", Datas)
			return
		}
		Posts, err := helper.GetPostForFilter(db, filterPosts)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}
		var PostsFiltered []models.HomeDataPost
		for _, v := range Posts {
			if v.PostLike >= likemin && v.PostLike <= likemax {
				PostsFiltered = append(PostsFiltered, v)
			}
		}
		Datas, err := helper.GetDataTemplate(db, r, true, false, false, false, true)
		if err != nil {
			helper.ErrorPage(w, http.StatusInternalServerError)
			return
		}

		//Set likes and dislikes
		Datasliked, err := helper.SetLikesAndDislikes(Datas.User, PostsFiltered, db)
		if err != nil {
			helper.ErrorPage(w, http.StatusBadRequest)
		}

		Datas.Datas = Datasliked
		helper.RenderTemplate(w, "index", "index", Datas)
	}

}

// RemoveDuplicates removes duplicate elements from a slice of uuid.UUID values.
func RemoveDuplicates(input []uuid.UUID) []uuid.UUID {
	unique := make(map[uuid.UUID]bool)
	result := []uuid.UUID{}

	for _, item := range input {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}

	return result
}

func GetFilteredPosts(db *sql.DB, posts []models.Post, minDate, maxDate string) ([]models.Post, error) {
	var filteredPosts []models.Post

	for _, post := range posts {

		createdAt, err := time.Parse("2006-01-02 15:04:05", post.CreatedAt)
		if err != nil {
			return nil, err
		}

		minDateTime, err := time.Parse("2006-01-02", minDate)
		if err != nil {
			return nil, err
		}

		maxDateTime, err := time.Parse("2006-01-02", maxDate)
		if err != nil {
			return nil, err
		}

		// Si aucune heure n'est fournie, ajustez les heures à minuit et 23:59:59
		minDateTime = minDateTime.Add(time.Hour * time.Duration(0))
		maxDateTime = maxDateTime.Add(time.Hour*time.Duration(23) + time.Minute*time.Duration(59) + time.Second*time.Duration(59))

		if createdAt.After(minDateTime) && createdAt.Before(maxDateTime) {
			filteredPosts = append(filteredPosts, post)
		}
	}

	return filteredPosts, nil
}

func GetFilteredPostsMyPage(db *sql.DB, posts []models.HomeDataPost, minDate, maxDate string) ([]models.HomeDataPost, error) {
	var filteredPosts []models.HomeDataPost

	for _, post := range posts {

		createdAt, err := time.Parse("2006-01-02 15:04:05", post.Posts.CreatedAt)
		if err != nil {
			return nil, err
		}

		minDateTime, err := time.Parse("2006-01-02", minDate)
		if err != nil {
			return nil, err
		}

		maxDateTime, err := time.Parse("2006-01-02", maxDate)
		if err != nil {
			return nil, err
		}

		// Si aucune heure n'est fournie, ajustez les heures à minuit et 23:59:59
		minDateTime = minDateTime.Add(time.Hour * time.Duration(0))
		maxDateTime = maxDateTime.Add(time.Hour*time.Duration(23) + time.Minute*time.Duration(59) + time.Second*time.Duration(59))

		if createdAt.After(minDateTime) && createdAt.Before(maxDateTime) {
			filteredPosts = append(filteredPosts, post)
		}
	}

	return filteredPosts, nil
}

func CompareDate(minDate, maxDate string) (bool, error) {
	// Analyser les dates
	minTime, err := time.Parse("2006-01-02", minDate)
	if err != nil {
		return false, err
	}

	maxTime, err := time.Parse("2006-01-02", maxDate)
	if err != nil {
		return false, err
	}

	// Comparer les dates
	return minTime.Before(maxTime), nil
}
