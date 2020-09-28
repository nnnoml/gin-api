package param_bind

//添加 绑定 结构体
type ProductAdd struct {
	Name string `form:"name" json:"name" validate:"required,NameValid"`
	Pwd  string `form:"pwd" json:"pwd" validate:"required,PwdValid"`
}
