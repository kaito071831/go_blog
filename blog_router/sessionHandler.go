package blog_router

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kaito071831/go_blog/utility"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ユーザーの型
type User struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
	Articles []Article
}

// セッションでユーザー名を保存しているキー
const userKey string = "UserID"

func init() {
	utility.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})
}

// パスワードをハッシュ化
func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 平文のパスワードとハッシュ化したパスワードを比較する
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// ユーザーを作成する
func createUser(username string, password string) error {
	passwordEncrypt, _ := passwordEncrypt(password)
	if err := utility.Db.Create(&User{Username: username, Password: passwordEncrypt}).Error; err != nil {
		return err
	}
	return nil
}

// ユーザーを取得する
func getUser(username string) User {
	var user User
	utility.Db.First(&user, "username = ?", username)
	return user
}

// ログインしているか確認する
func isLogin(c *gin.Context) bool {
	username := sessions.Default(c).Get(userKey)
	if username == nil {
		c.Redirect(http.StatusSeeOther, "/login")
		return false
	}
	return true
}

// ユーザー新規登録
func Signup(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "session/signup.html", nil)
	case "POST":
		var form User
		if err := c.Bind(&form); err != nil {
			c.HTML(http.StatusBadRequest, "session/signup.html", err)
			c.Abort()
		} else {
			username := c.PostForm("username")
			password := c.PostForm("password")

			// ユーザーが重複したときのエラー処理
			if err := createUser(username, password); err != nil {
				c.HTML(http.StatusBadRequest, "session/signup.html", err)
				c.Abort()
			}
			c.Redirect(http.StatusFound, "/")
		}
	default:
		c.HTML(http.StatusOK, "session/signup.html", nil)
	}
}

// ユーザーログイン
func Login(c *gin.Context) {
	switch c.Request.Method {
	case "GET":
		c.HTML(http.StatusOK, "session/login.html", nil)
	case "POST":

		user := getUser(c.PostForm("username"))

		// DBに登録されているハッシュ化されたパスワードを取得する
		dbPassword := user.Password

		// フォームから取得したパスワード
		formPassword := c.PostForm("password")

		session := sessions.Default(c)

		// ユーザーパスワードの比較
		if err := compareHashAndPassword(dbPassword, formPassword); err != nil {
			log.Println("ログインできませんでした")
			c.HTML(http.StatusOK, "session/login.html", err)
			c.Abort()
		} else {
			session.Set(userKey, user.Username)
			if err := session.Save(); err != nil {
				log.Println("セッションの保存に失敗しました")
				c.HTML(http.StatusOK, "session/login.html", err)
				c.Abort()
			}
			log.Println("ログインできました")
			c.Redirect(http.StatusSeeOther, "/article")
		}
	default:
		c.HTML(http.StatusOK, "session/login.html", nil)
	}
}

// ユーザーログアウト
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(userKey)
	if err := session.Save(); err != nil {
		log.Printf("ログアウトに失敗しました: %v", err)
		return
	}
	c.Redirect(http.StatusSeeOther, "/")
}
