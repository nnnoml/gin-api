package product

import (
	"fmt"
	"go-gin-api/app/common"
	"go-gin-api/app/controller/product/param_bind"
	"go-gin-api/app/controller/product/param_verify"
	"go-gin-api/app/database/mysql"
	"go-gin-api/app/util/bind"
	"go-gin-api/app/util/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xinliangnote/go-util/md5"
)

//数据结构体
type userInfo struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Pwd       string `db:"pwd"`
	Salt      string `db:"salt"`
	CreatedAt string `db:"created_at"`
}

//新增
func Add(c *gin.Context) {
	//gin 对象
	utilGin := response.Gin{Ctx: c}
	// 参数绑定
	bindRes, err := bind.Bind(&param_bind.ProductAdd{}, c)
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	//参数校验器
	err = param_verify.ValidHandle(bindRes)
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}

	// 业务处理...

	//获取参数
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	// var salt string
	//获取随机盐
	salt, err := common.RandomStr(6)
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	//md5加密
	md5pwd := md5.MD5(pwd + salt)
	//获取时间
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	//插入库
	res, err := mysql.Db.Exec("insert into user(name, pwd, salt,created_at)values(?, ?, ?,?)", name, md5pwd, salt, createdAt)
	if err != nil {
		utilGin.Response(-1, "exec failed", nil)
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		utilGin.Response(-1, "exec failed", nil)
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
	//返回值
	ages := make(map[string]int64)
	ages["id"] = id

	utilGin.Response(1, "success", ages)
}

//编辑
func Edit(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	//打印url
	fmt.Println(c.Request.RequestURI)
	//获取路由参数
	id := c.Param("id")

	//获取put参数
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")

	// var salt string
	//获取随机盐
	salt, err := common.RandomStr(6)
	if err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	//md5加密
	md5pwd := md5.MD5(pwd + salt)
	//获取时间
	createdAt := time.Now().Format("2006-01-02 15:04:05")
	//更新库
	res, err := mysql.Db.Exec("update user set name=? , pwd=? , salt=? , created_at=? where id=?", name, md5pwd, salt, createdAt, id)

	if err != nil {
		utilGin.Response(-1, "exec failed", nil)
		fmt.Println("exec failed, ", err)
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		utilGin.Response(-1, "exec failed", nil)
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("update succ , Rows Affected:", affected)
	//返回值
	ages := make(map[string]string)
	ages["id"] = id

	utilGin.Response(1, "success", ages)
}

//删除
func Delete(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	fmt.Println(c.Request.RequestURI)
	//获取路由参数
	id := c.Param("id")

	//更新库
	res, err := mysql.Db.Exec("delete from user  where id=?", id)
	if err != nil {
		utilGin.Response(-1, "exec failed", nil)
		fmt.Println("exec failed, ", err)
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		utilGin.Response(-1, "exec failed", nil)
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("delete succ , Rows Affected:", affected)
	//返回值
	ages := make(map[string]string)
	ages["id"] = id

	utilGin.Response(1, "success", ages)
}

//详情
func Detail(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	//打印url
	fmt.Println(c.Request.RequestURI)
	//获取路由参数
	id := c.Param("id")
	//结构体数组
	var userInfoSlice []userInfo

	err := mysql.Db.Select(&userInfoSlice, "select * from user where id="+id)
	if err != nil {
		utilGin.Response(-1, "exec failed", nil)
		fmt.Println("exec failed, ", err)
		return
	}

	for _, v := range userInfoSlice {
		fmt.Println(v.Name)
	}

	utilGin.Response(1, "detail", userInfoSlice)
}
