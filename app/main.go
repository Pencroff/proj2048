// build js, window

package main

import (
	"fmt"
	log "github.com/go-pkgz/lgr"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pencroff/proj2048/app/resources"
	"github.com/pencroff/proj2048/app/scene"
	"github.com/sedyh/mizu/pkg/engine"
	"os"
)

func main() {
	//db := doMigrations()
	//log.Printf("db: %+v", db)
	initResources()
	runGame()
}

func initResources() {
	resources.InitColors()
	resources.InitFonts()
	resources.InitTiles()
}

func runGame() {
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOn)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(500, 800)
	if err := ebiten.RunGame(engine.NewGame(&scene.Game{})); err != nil {
		log.Printf("[Fatal] Can't run game: %v", err)
	}
}

func doMigrations() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Printf("[Info] Connecting to database")
	pgConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"))
	db, err := sqlx.Connect("postgres", pgConnStr)
	if err != nil {
		log.Fatalf("[ERROR] Can't connect to db: %v", err)
	}
	log.Printf("[Info] Migrate Db")
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	m.Down()
	return db
}

//r := rng.NewSmallPrng()
//r.Seed(time.Now().UnixNano())
//fmt.Printf("seed: %x%x\n", r.Uint64(), r.Uint64())
