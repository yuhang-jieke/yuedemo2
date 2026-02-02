package request

type Register struct {
	Name    string `form:"name"  binding:"required"`
	Age     int    `form:"age"  binding:"required"`
	Address string `form:"address"  binding:"required"`
}
type Login struct {
	Name string `form:"name"  binding:"required"`
	Age  int    `form:"age"  binding:"required"`
}
