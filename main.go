package main

import (
	"GunungKuy/config"

	ad "GunungKuy/features/admin/delivery"
	ar "GunungKuy/features/admin/repository"
	as "GunungKuy/features/admin/services"
	bd "GunungKuy/features/booking/delivery"
	br "GunungKuy/features/booking/repository"
	bs "GunungKuy/features/booking/services"
	cd "GunungKuy/features/climber/delivery"
	cr "GunungKuy/features/climber/repository"
	cs "GunungKuy/features/climber/services"
	gd "GunungKuy/features/google/delivery"
	gr "GunungKuy/features/google/repository"
	gs "GunungKuy/features/google/services"
	pd "GunungKuy/features/product/delivery"
	pr "GunungKuy/features/product/repository"
	ps "GunungKuy/features/product/services"
	rd "GunungKuy/features/ranger/delivery"
	rr "GunungKuy/features/ranger/repository"
	rs "GunungKuy/features/ranger/services"
	ud "GunungKuy/features/users/delivery"
	ur "GunungKuy/features/users/repository"
	us "GunungKuy/features/users/services"
	"GunungKuy/utils/database"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	uRepo := ur.New(db)
	bRepo := br.New(db)
	aRepo := ar.New(db)
	gRepo := gr.New(db)
	pRepo := pr.New(db)
	cRepo := cr.New(db)
	rRepo := rr.New(db)
	uService := us.New(uRepo)
	bService := bs.New(bRepo)
	aService := as.New(aRepo)
	pService := ps.New(pRepo)
	gService := gs.New(gRepo)
	cService := cs.New(cRepo)
	rService := rs.New(rRepo)
	ud.New(e, uService)
	bd.New(e, bService)
	ad.New(e, aService)
	pd.New(e, pService)
	gd.New(e, gService)
	cd.New(e, cService)
	rd.New(e, rService)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)))
}
