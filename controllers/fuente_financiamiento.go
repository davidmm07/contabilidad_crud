package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mongo_crud/managers/logManager"
	"github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// FuenteFinanciamientoController operations for FuenteFinanciamiento
type FuenteFinanciamientoController struct {
	beego.Controller
}

// URLMapping ...
func (j *FuenteFinanciamientoController) URLMapping() {
	j.Mapping("Post", j.Post)
	j.Mapping("VincularFuente", j.VincularFuente)
}

// Post ...
// @Title Create
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {object} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router / [post]
func (j *FuenteFinanciamientoController) Post() {
	var fuente models.FuenteFinanciamiento

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			logManager.LogError(r)
			panic(r)
		}
	}()

	json.Unmarshal(j.Ctx.Input.RequestBody, &fuente)
	log.Println("fuente: ", fuente)

	if err := models.InsertFuenteFinanciamiento(&fuente); err == nil {
		j.Data["json"] = "insert success!"
	} else {
		j.Data["json"] = err.Error()
	}
}

// VincularFuente ...
// @Title Create
// @Description create FuenteFinanciamiento
// @Param	body		body 	models.FuenteFinanciamiento	true		"body for FuenteFinanciamiento content"
// @Success 201 {object} models.FuenteFinanciamiento
// @Failure 403 body is empty
// @router /VincularFuente [post]
func (j *FuenteFinanciamientoController) VincularFuente() {
	var fuente models.FuenteFinanciamiento
	json.Unmarshal(j.Ctx.Input.RequestBody, &fuente)
	fmt.Println("fuente:", fuente)
}

// Put de HTTP
// @Title Update
// @Description update the FuenteFinanciamiento
// @Param	codigo		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :codigo is empty
// @router /:codigo [put]
func (j *FuenteFinanciamientoController) Put() {
	codigo := j.Ctx.Input.Param(":objectId")
	var fuente models.FuenteFinanciamiento

	json.Unmarshal(j.Ctx.Input.RequestBody, &fuente)

	if err := models.UpdateFuenteFinanciamiento(&fuente, codigo); err == nil {
		j.Data["json"] = "update success!"
	} else {
		j.Data["json"] = err.Error()
	}
}
