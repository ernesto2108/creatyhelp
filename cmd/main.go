package main

import (
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/golang-migrate/migrate/v4"
	migration "github.com/golang-migrate/migrate/v4/database/postgres"

	app "github.com/ernesto2108/AP_CreatyHelp/pkg/user/application/handlers"

	api "github.com/ernesto2108/AP_CreatyHelp/api/routes"
	server "github.com/ernesto2108/AP_CreatyHelp/api/server"
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = logs.InitLogger()

	client := internal.NewPSqlClient("postgres://postgres:Root@localhost:5432/test?sslmode=disable")
	doMigrate(client, "test")


	handler := app.NewCreateUsersHandler(client)

	r := api.AllRoutes(handler)
	server := server.NewServer(r)
	server.Run()
}

func doMigrate(cl *internal.PostSqlClient, dbName string)  {
	driver, _ := migration.WithInstance(cl.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version in %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}