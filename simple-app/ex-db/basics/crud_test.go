package basics

import (
	testutil "lesson02examples/testuitl"
	"testing"
	"time"

	"gorm.io/gorm"
)

func TestCrudDemo(t *testing.T) {
	// create a new test database
	db := testutil.NewTestDB(t, "crud.db")

	type User struct {
		ID        uint      `gorm:"primaryKey"`
		Name      string    `gorm:"size:64;not null"`
		Email     string    `gorm:"size:128;uniqueIndex;not null"`
		Age       uint8     `gorm:"not null"`
		Status    string    `gorm:"size:16;default:active;index"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		t.Fatalf("auto migrate: %v", err)
	}

	// create one record
	// user := User{Name: "Alice", Email: "alice@qq.com", Age: 18, Status: "active"}
	// if err := db.Create(&user).Error; err != nil {
	// 	t.Fatalf("create user: %v", err)
	// }

	// create multiple records
	// users := []User{
	// 	{Name: "alice1", Email: "alice1@qq.com", Age: 19, Status: "active"},
	// 	{Name: "alice2", Email: "alice2@qq.com", Age: 20, Status: "inactive"},
	// 	{Name: "alice3", Email: "alice3@qq.com", Age: 21, Status: "inactive"},
	// 	{Name: "bob", Email: "bob@qq.com", Age: 20, Status: "active"},
	// 	{Name: "bob1", Email: "bob1@qq.com", Age: 21, Status: "active"},
	// 	{Name: "bob2", Email: "bob2@qq.com", Age: 22, Status: "inactive"},
	// 	{Name: "bob3", Email: "bob3@qq.com", Age: 23, Status: "inactive"},
	// }

	// if err := db.Create(&users).Error; err != nil {
	// 	t.Fatalf("create users: %v", err)
	// }

	//find first record
	// var user User
	// if err := db.First(&user).Error; err != nil {
	// 	t.Fatalf("find first user: %v", err)
	// }
	// t.Logf("first active user: %+v", user)

	//find user by primary key
	// if err := db.First(&user, 3).Error; err != nil {
	// 	t.Fatalf("find user by primary key: %v", err)
	// }
	// t.Logf("user with primary key 1: %+v", user)

	// find records by condition
	// var users []User
	// if err := db.Where("status = ? and age>?", "active", 20).Find(&users).Error; err != nil {
	// 	t.Fatalf("find users by condition: %v", err)
	// }
	// t.Logf("users with status active: %+v", users)

	//scan records to struct
	// type UserSummary struct {
	// 	Name  string
	// 	Email string
	// }
	// var summaries []UserSummary
	// if err := db.Model(&user).Select("name", "email").Scan(&summaries).Error; err != nil {
	// 	t.Fatalf("scan users to summary: %v", err)
	// }
	// t.Logf("user summaries: %+v", summaries)

	//find records by complax condition
	// var users []User
	// if err := db.Where("status in (?) and age between ? and ? and email like ?", []string{"VIP", "active"}, 18, 40, "b%").Order("id desc").Find(&users).Error; err != nil {
	// 	t.Fatalf("find users by complex condition: %v", err)
	// }
	// t.Logf("users with complex condition: %+v", users)

	//find page records
	var users []User
	page := 2
	pageSize := 3
	offset := (page - 1) * pageSize
	if err := db.Limit(pageSize).Offset(offset).Order("id desc").Find(&users).Error; err != nil {
		t.Fatalf("find page users: %v", err)
	}
	t.Logf("page %d users: %+v", page, users)

	//update records with condition
	// if err := db.Model(&user).Where(" status = ?", "pending").Updates(map[string]any{"status": "active", "age": 17}).Error; err != nil {
	// 	t.Fatalf("update users: %v", err)
	// }

	//update all records
	// if err := db.Model(&user).Where("1=1").Updates(User{Age: 33, Status: "VIP"}).Error; err != nil {
	// 	t.Fatalf("update users: %v", err)
	// }

	//find user by email and update age
	// if err := db.First(&user, "email = ?", "bob2@qq.com").Error; err != nil {
	// 	t.Fatalf("find user by email: %v", err)
	// }
	// user.Age = 58
	// user.Status = "active"
	// if err := db.Model(&user).Updates(&user).Error; err != nil {
	// 	t.Fatalf("update user age: %v", err)
	// }

	//find user by email and delete user
	// if err := db.First(&user, "email = ?", "bob1@qq.com").Error; err != nil {
	// 	t.Fatalf("find user by email: %v", err)
	// }
	// if err := db.Delete(&user).Error; err != nil {
	// 	t.Fatalf("delete user: %v", err)
	// }

	//delete users with condition
	// if err := db.Where(" status = ?", "active").Delete(&user).Error; err != nil {
	// 	t.Fatalf("delete users: %v", err)
	// }
}

func TestCrudExerciseDemo(t *testing.T) {
	// create a new test database
	db := testutil.NewTestDB(t, "crud_exercise.db")

	type User struct {
		ID          uint      `gorm:"primaryKey"`
		Name        string    `gorm:"size:64;not null"`
		Email       string    `gorm:"size:128;uniqueIndex;not null"`
		Age         uint8     `gorm:"not null"`
		Status      string    `gorm:"size:16;default:active;index"`
		Phone       string    `gorm:"size:32;uniqueIndex;not null"`
		LastLoginAt time.Time `gorm:"type:datetime;default:null"`
		CreatedAt   time.Time `gorm:"autoCreateTime"`
		UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		t.Fatalf("auto migrate: %v", err)
	}

	// create multiple records
	// users := []User{
	// 	{Name: "hehe", Email: "hehe@qq.com", Age: 10, Status: "active", Phone: "1234567890", LastLoginAt: time.Now()},
	// 	{Name: "hehe1", Email: "hehe1@qq.com", Age: 11, Status: "active", Phone: "1234567891", LastLoginAt: time.Now()},
	// 	{Name: "hehe2", Email: "hehe2@qq.com", Age: 12, Status: "active", Phone: "1234567892", LastLoginAt: time.Now()},
	// 	{Name: "hehe3", Email: "hehe3@qq.com", Age: 13, Status: "active", Phone: "1234567893", LastLoginAt: time.Now()},
	// 	{Name: "hehe4", Email: "hehe4@qq.com", Age: 14, Status: "inactive", Phone: "1234567894", LastLoginAt: time.Now()},
	// 	{Name: "hehe5", Email: "hehe5@qq.com", Age: 15, Status: "inactive", Phone: "1234567895", LastLoginAt: time.Now()},
	// 	{Name: "hehe6", Email: "hehe6@qq.com", Age: 16, Status: "inactive", Phone: "1234567896", LastLoginAt: time.Now()},
	// 	{Name: "hehe7", Email: "hehe7@qq.com", Age: 17, Status: "active", Phone: "1234567897", LastLoginAt: time.Now()},
	// 	{Name: "hehe8", Email: "hehe8@qq.com", Age: 18, Status: "active", Phone: "1234567898", LastLoginAt: time.Now()},
	// 	{Name: "hehe9", Email: "hehe9@qq.com", Age: 19, Status: "active", Phone: "1234567899", LastLoginAt: time.Now()},
	// 	{Name: "hehe10", Email: "hehe10@qq.com", Age: 10, Status: "active", Phone: "1234567880", LastLoginAt: time.Now()},
	// 	{Name: "hehe11", Email: "hehe11@qq.com", Age: 11, Status: "active", Phone: "1234567881", LastLoginAt: time.Now()},
	// }
	// if err := db.Create(&users).Error; err != nil {
	// 	t.Fatalf("create users: %v", err)
	// }

	// find users with email pattern and paginate results
	// myusers := []User{}
	// emailPattern := "hehe1%"
	// page := 1
	// size := 2
	// if page < 1 {
	// 	page = 1
	// }
	// if size < 1 {
	// 	size = 10
	// }
	// offset := (page - 1) * size
	// if err := db.Where("email LIKE ?", emailPattern).Limit(size).Offset(offset).Find(&myusers).Error; err != nil {
	// 	t.Fatalf("search users by email: %v", err)
	// }
	// t.Logf("search users by email pattern %s: %+v", emailPattern, myusers)

	//update multiple users with condition
	// ids := []uint{7, 8, 9}
	// if err := db.Model(&User{}).Where("id in ?", ids).Updates(map[string]any{"status": "pending"}).Error; err != nil {
	// 	t.Fatalf("update users status: %v", err)
	// }

	// ids := []uint{7, 8, 9}
	// if err := db.Model(&User{}).Where("id in ?", ids).Updates(map[string]any{"last_login_at": time.Now().AddDate(0, 0, -30)}).Error; err != nil {
	// 	t.Fatalf("update users last login: %v", err)
	// }

	// if err := db.Model(&User{}).Where("last_login_at < ?", time.Now().AddDate(0, 0, -30)).Delete(&User{}).Error; err != nil {
	// 	t.Fatalf("delete inactive users: %v", err)
	// }

	myusers := []User{}
	if err := db.Scopes(pagenite(1, 3)).Where("age between ? and ?", 10, 12).Find(&myusers).Error; err != nil {
		t.Fatalf("search users by age: %v", err)
	}
	t.Logf("search users by age: %+v", myusers)

}

func pagenite(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		if size < 1 {
			size = 5
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
