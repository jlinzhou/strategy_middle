//实现策略的增删改查功能，并和前端进行交互


package api

import(
	"net/http"
	"strategy_middle/logging"
	"strategy_middle/models"
	"reflect"
	"gopkg.in/mgo.v2/bson"
	//"github.com/Unknwon/com"
	//"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func Get_Stra_DataAll() (strategy_infos  []models.Strategy_info){
	c := models.ConnecToDB("strategy_info")
	err := c.Find(nil).All(&strategy_infos)
	if err != nil {

		logging.Error(err)
	}
	logging.Info(strategy_infos)
	return
}

func GetALLInst(c *gin.Context) {
	/*
	   c.Query可用于获取?id=1这类URL参数，而c.DefaultQuery则支持设置一个默认值
	   code变量使用了e模块的错误编码，这正是先前规划好的错误码，方便排错和识别记录
	   util.GetPage保证了各接口的page处理是一致的
	   c *gin.Context是Gin很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的JSON和呈现JSON响应
	*/
	//id := c.Query("id")
	
	d := models.ConnecToDB("strategy_info")
	strategy_infos := make([]models.Strategy_info, 20)
	err := d.Find(nil).All(&strategy_infos)
	if err != nil {
		logging.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": strategy_infos,
	})

}
func AddInst(c *gin.Context){
	d := models.ConnecToDB("strategy_info")
	//result := make([]models.Strategy_info,20)
	 
	mydata := d.Find(nil).Sort(bson.M{"id","1"})
	t:=reflect.TypeOf(mydata)
	logging.Info(t)
	 
	// for i:=0;i<t.NumField();i++{
	// 	f:=t.Field(i)
	// 	val :=v.Field(i).Interface()
	// 	logging.Info(f.Name,f.Type,val)
	// }
	//t := reflect.TypeOf(mydata) 
	if mydata != nil {
		logging.Info(mydata)
		logging.Info(reflect.ValueOf(mydata))
	}
		c.JSON(http.StatusOK, gin.H{
		"data": mydata,
	})
	//logging.Info(result)
	// .Sort("-age")
	// stra := models.Strategy_info{
	// 	Instance_name:  "zhangsan",
	// 	Phone: "13480989765",
	// 	Email: "329832984@qq.com",
	// 	Sex:   "F",
	// }

}

// func InsertToMogo() {
// 	c := ConnecToDB("student2")
// 	stu1 := Student{
// 		Name:  "zhangsan",
// 		Phone: "13480989765",
// 		Email: "329832984@qq.com",
// 		Sex:   "F",
// 	}
// 	stu2 := Student{
// 		Name:  "liss",
// 		Phone: "13980989767",
// 		Email: "12832984@qq.com",
// 		Sex:   "M",
// 	}
// 	err := c.Insert(&stu1, &stu2)
// 	if err != nil {

// 		logging.Fatal(err)
// 	}
// }