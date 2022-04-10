package controllers

import (
	"net/http"
	"strconv"

	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
	"hodeinaiz/go_mvc/models"
)

var decoder = schema.NewDecoder()

type GO_mvc struct {
	controller.BaseController
	Routes []string
}

func (t *GO_mvc) Home() {
	messages := []*models.Message{}
	t.Ctx.DB.Order("created_at desc").Find(&messages)
	t.Ctx.Data["List"] = messages
	t.Ctx.Template = "index"
	t.HTML(http.StatusOK)
}

func (t *GO_mvc) Create() {
	message := &models.Message{}
	req := t.Ctx.Request()
	_ = req.ParseForm()
	if err := decoder.Decode(message, req.PostForm); err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}

	t.Ctx.DB.Create(message)
	t.Ctx.Redirect("/", http.StatusFound)
}

func (t *GO_mvc) Delete() {
	messageID := t.Ctx.Params["id"]
	ID, err := strconv.Atoi(messageID)
	if err != nil {
		t.Ctx.Data["Message"] = err.Error()
		t.Ctx.Template = "error"
		t.HTML(http.StatusInternalServerError)
		return
	}
	t.Ctx.DB.Delete(&models.Message{ID: ID})
	t.Ctx.Redirect("/", http.StatusFound)
}

func NewMessage() controller.Controller {
	return &GO_mvc{
		Routes: []string{
			"get;/;Home",
			"post;/create;Create",
			"get;/delete/{id};Delete",
		},
	}
}
