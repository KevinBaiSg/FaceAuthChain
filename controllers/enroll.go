package controllers

import (
	//"encoding/json"
	//"my/FaceRecognizerServer/models"

	"github.com/astaxie/beego"
)

// Operations about object
type EnrollController struct {
	beego.Controller
}

// @Title Enroll
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (e *EnrollController) Post() {
	//var ob models.Object
	//json.Unmarshal(e.Ctx.Input.RequestBody, &ob)
	//objectid := models.AddOne(ob)
	//e.Data["json"] = map[string]string{"ObjectId": objectid}
	//e.ServeJSON()
}
