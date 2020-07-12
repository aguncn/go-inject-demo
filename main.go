package main

import (
	"fmt"

	"github.com/facebookgo/inject"
)

type DBEngine struct {
	Name string
}

type NezhaDB struct {
	Db *DBEngine `inject:""`
}

type NezhaService struct {
	Db *NezhaDB `inject:""`
}

type App struct {
	Name     string
	NezhaSVC *NezhaService `inject:""`
}

func (a *App) Create() string {
	return "create app, in db name:" + a.NezhaSVC.Db.Db.Name + " app name :" + a.Name
}

type Object struct {
	App *App
}

func Init() *Object {
	db := DBEngine{Name: "nezha-db"}
	var g inject.Graph
	app := App{Name: "nezha-app"}

	_ = g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: &db},
	)
	_ = g.Populate()
	return &Object{
		App: &app,
	}

}

func main() {
	obj := Init()
	fmt.Println(obj.App.Create())
}
