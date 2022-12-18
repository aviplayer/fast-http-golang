package main

import (
	"fast-http-golang/db"
	"fast-http-golang/dto"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var conn = db.CreateConnection()

func Handle(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/api":
		handleApi(ctx)
	}
}

func handleApi(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)

	b := ctx.Request.Body()
	user := dto.User{}
	if err := json.Unmarshal(b, &user); err != nil {
		fmt.Printf("can,t parse %s\n", string(b))
		createError(err, ctx)
		return
	}
	id, err := conn.InsertUser(user.Username, user.Email, user.Age)
	if err != nil {
		createError(err, ctx)
		return
	}
	usr, err := conn.GetById(id)
	if err != nil {
		createError(err, ctx)
		return
	}

	usr, err = conn.DeleteUser(usr)
	if err != nil {
		createError(err, ctx)
		return
	}
	b, err = json.Marshal(&usr)
	if err != nil {
		fmt.Printf("can,t deserialize %s\n", string(b))
		createError(err, ctx)
		return
	}
	fmt.Println(usr)
	ctx.SetBody(b)
}

func createError(err error, ctx *fasthttp.RequestCtx) {
	fmt.Println(err)
	ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
	errMsg := fmt.Sprintf("%s", err.Error())
	ctx.Response.AppendBodyString(errMsg)
}
