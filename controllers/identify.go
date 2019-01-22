package controllers

import (
	//"encoding/json"
	//"my/FaceAuthChain/models"

	"github.com/astaxie/beego"
)

// Operations about object
type IdentifyController struct {
	beego.Controller
}

// @Title Identify
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (i *IdentifyController) Post() {
	//var ob models.Object
	//json.Unmarshal(i.Ctx.Input.RequestBody, &ob)
	//objectid := models.AddOne(ob)
	//i.Data["json"] = map[string]string{"ObjectId": objectid}
	//i.ServeJSON()
}
