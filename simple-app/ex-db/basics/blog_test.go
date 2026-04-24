package basics

import (
	testutil "lesson02examples/testuitl"
	"testing"
)

func TestBlogDemo(t *testing.T) {
	// create a new test database
	db := testutil.NewTestDB(t, "blog.db")

	if err := db.AutoMigrate(&User{}, &Post{}, &Comment{}, &Tag{}); err != nil {
		t.Fatalf("auto migrate: %v", err)
	}

	// insert datas
	// tags := []Tag{
	// 	{Name: "golang"},
	// 	{Name: "gorm"},
	// 	{Name: "solidity"},
	// }

	// if err := db.Create(&tags).Error; err != nil {
	// 	panic(err)
	// }

	// posts := []Post{
	// 	{
	// 		Title:   "alice's first post",
	// 		Content: "this is alice's first post",
	// 		Tags: []Tag{
	// 			tags[0],
	// 			tags[1],
	// 		},
	// 	},
	// 	{
	// 		Title:   "alice's second post",
	// 		Content: "this is alice's second post",
	// 		Tags: []Tag{
	// 			tags[2],
	// 		},
	// 	},
	// 	{
	// 		Title:   "bob's first post",
	// 		Content: "this is bob's first post",
	// 		Tags: []Tag{
	// 			tags[1], // ✅ 不会重复插入
	// 		},
	// 	},
	// }

	// // 3. 构造 users
	// users := []User{
	// 	{
	// 		Name:   "alice",
	// 		Email:  "alice@qq.com",
	// 		Age:    18,
	// 		Status: "active",
	// 		Posts:  []Post{posts[0], posts[1]},
	// 	},
	// 	{
	// 		Name:   "bob",
	// 		Email:  "bob@qq.com",
	// 		Age:    20,
	// 		Status: "active",
	// 		Posts:  []Post{posts[2]},
	// 	},
	// }

	// // 4. 直接创建（❗不要用 FullSaveAssociations）
	// if err := db.Create(&users).Error; err != nil {
	// 	panic(err)
	// }

	// comments := []Comment{
	// 	{Content: "content of alice's first post by alice", PostID: users[0].Posts[0].ID, UserID: users[0].ID},
	// 	{Content: "content of alice's first post by bob", PostID: users[0].Posts[0].ID, UserID: users[1].ID},
	// 	{Content: "content of bob's first post by alice", PostID: users[1].Posts[0].ID, UserID: users[0].ID},
	// }
	// if err := db.Create(&comments).Error; err != nil {
	// 	panic(err)
	// }

	// find posts of user
	// var postOfUser []Post
	// if err := db.Where("user_id = ?", 1).Order("created_at desc").Limit(10).Find(&postOfUser).Error; err != nil {
	// 	panic(err)
	// }
	// t.Logf("posts of alice: %+v", postOfUser)

	// preload comments of posts
	// var posts []Post
	// if err := db.Preload("Comments").Find(&posts).Error; err != nil {
	// 	panic(err)
	// }
	// t.Logf("posts with comments: %+v", posts)

	// find post's comments count
	// var postWithCommentCount struct {
	// 	Post         []Post
	// 	CommentCount int64
	// }

	// if err := db.Table("posts").Select("posts.*, count(comments.id) as comment_count").Joins("left join comments on comments.post_id = posts.id").Group("posts.id").Scan(&postWithCommentCount).Error; err != nil {
	// 	panic(err)
	// }
	// t.Logf("post with comment count: %+v", postWithCommentCount)

	//transaction
	// err := db.Transaction(func(tx *gorm.DB) error {
	// 	var us User
	// 	if err := tx.Model(&User{}).Where("id = ?", 1).First(&us).Error; err != nil {
	// 		return err
	// 	}
	// 	po := Post{
	// 		Title:   "alice's third post",
	// 		Content: "this is alice's third post",
	// 	}

	// 	po.UserID = us.ID
	// 	if err := tx.Model(&Post{}).Create(&po).Error; err != nil {
	// 		return err
	// 	}

	// 	var tags []Tag
	// 	if err := tx.Model(&Tag{}).Find(&tags).Error; err != nil {
	// 		return err
	// 	}
	// 	po.Tags = []Tag{tags[0], tags[2]}

	// 	if err := tx.Model(&po).Updates(&po).Error; err != nil {
	// 		return err
	// 	}
	// 	t.Logf(" alice's post count:%d", len(us.Posts))

	// 	return nil
	// })

	// if err != nil {
	// 	t.Fatalf("transaction failed: %v", err)
	// }

	//soft delete
	// if err := db.Model(&Comment{}).Where("id = ?", 2).Delete(&Comment{}).Error; err != nil {
	// 	t.Fatalf("soft delete comment: %v", err)
	// }

	// var comment Comment
	// var comments []Comment
	// if err := db.Unscoped().Model(&Comment{}).Where("id = ?", 2).First(&comment).Error; err != nil {
	// 	t.Fatalf("find soft deleted comment: %v", err)
	// }
	// t.Logf("soft deleted comment: %+v", comment)

	// if err := db.Find(&comments).Error; err == nil {
	// 	t.Fatalf("find soft deleted comment without unscoped: %v", err)
	// }

	//hard delete
	// if err := db.Unscoped().Model(&Comment{}).Where("id = ?", 1).Delete(&Comment{}).Error; err != nil {
	// 	t.Fatalf("hard delete comment: %v", err)
	// }

	count := int64(0)
	var comments []Comment
	if err := db.Model(&Comment{}).Where("post_id = ?", 3).Count(&count).Error; err != nil {
		t.Fatalf("count comments of post 1: %v", err)
	}
	t.Logf(" post's comment count %d", count)

	if err := db.Model(&Comment{}).Where("post_id = ?", 3).Find(&comments).Error; err != nil {
		t.Fatalf("find comments of post 1: %v", err)
	}
	t.Logf(" post's comments: %+v", comments)
}
