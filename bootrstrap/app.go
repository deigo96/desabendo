package bootrstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	Db  *Databases
}

type ServerConfig struct {
	Host string
	Port string
}

type Databases struct {
	Db       *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Db = &Databases{
		Db: NewDatabase(app.Env),
	}
	return *app
}

func (app *Application) CloseDBConnection() {
	go CloseConnection(app.Db.Db)
}
