package admin

import (
	"fmt"
	"gindemo/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	// c.String(200, "admin/user接口==")
	// con.success(c)

	//查询数据库
	// userList := []models.User{}

	// models.DB.Find(&userList)
	// c.JSON(200, gin.H{
	// 	"result": userList,
	// })

	//查询age大于20的用户
	userList := []models.User{}
	models.DB.Where("age >?", 20).Find(&userList)
	c.JSON(200, gin.H{
		"result": userList,
	})
}

func (con UserController) Add(c *gin.Context) {

	user := models.User{
		Username: "itying",
		Age:      22,
		Email:    "itying@163.com",
		// AddTime:  1000011111,
		AddTime: int(models.GetUnix()),
	}
	models.DB.Create(&user)
	fmt.Println("user:", user)
	// c.String(200, "admin/user/add接口==")
	c.String(200, "增加数据成功")

}

func (con UserController) Edit(c *gin.Context) {

	//方法1：
	//保存所有字段

	// user := models.User{Id: 4}
	// models.DB.Find(&user)
	// fmt.Println(user)
	// user.Username = "哈哈"
	// user.Email = "haha@qq.com"
	// models.DB.Save(&user)

	//方法2：
	//更新单个列
	// user := models.User{}
	// models.DB.Model(&user).Where("id = ?", 4).Update("username", "嘿嘿")
	// c.String(200, "嘿嘿修改成功")

	//方法3
	user := models.User{}
	models.DB.Where("id=?", 4).Find(&user)

	user.Email = "heihei@qq.com"
	models.DB.Save(&user)

	c.String(200, "admin/user/edit接口==")
}

func (con UserController) Delete(c *gin.Context) {

	// user := models.User{Id: 4}
	// models.DB.Delete(&user)

	user := models.User{}
	models.DB.Where("username=?", "嘿嘿").Delete(&user)
	c.String(200, "admin/user/delete接口")
}
