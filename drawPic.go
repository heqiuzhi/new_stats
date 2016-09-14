package main

import (
	"log"
	"os"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/jinzhu/gorm"
)

func draw() {

	heroData := getHeroData()

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "英雄散列图"
	p.X.Label.Text = "Damage"
	p.Y.Label.Text = "Control"

	// bs, err := plotter.NewHeroes(heroData)
	// if err != nil {
	// 	panic(err)
	// }

	p.Add(plotter.NewHeroShowInfo(heroData))

	os.Remove("heroShow.png")
	if err := p.Save(20*vg.Inch, 12*vg.Inch, "heroShow1.png"); err != nil {
		panic(err)
	}
}

func getHeroData() []plotter.HeroShowInfo {
	db, err := gorm.Open("mysql", "root:123456@/dota2_new_stats?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("failed to connect database\n")
	}
	defer db.Close()

	type Result struct {
		Hero_name string
		Damage    float64
		Control   float64
	}
	var results []Result
	db.Table("stats").Select("hero_name,avg(create_deadly_damages_per_death) as damage,avg(create_deadly_stiff_control_per_death) as control").Group("hero_name").Scan(&results)

	data := make([]plotter.HeroShowInfo, len(results))
	for i, aResult := range results {
		data[i].Hero_name = aResult.Hero_name
		data[i].Damage = aResult.Damage
		data[i].Control = aResult.Control
	}

	// rows, err := db.Table("stats").Select("hero_name,avg(create_deadly_damages_per_death) as damage,avg(create_deadly_stiff_control_per_death) as control").Group("hero_name").Rows()
	// // var results []Result
	// // db.Table("stats").Select("hero_name,avg(create_deadly_damages_per_death) as damage,avg(create_deadly_stiff_control_per_death) as control").Group("hero_name").Scan(&results)

	// n := 105
	// // db.Table("stats").Select("hero_name,avg(create_deadly_damages_per_death) as damage,avg(create_deadly_stiff_control_per_death) as control").Group("hero_name").Count(&n)
	// data := make(plotter.XYZs, n)
	// i := 0
	// for rows.Next() {
	// 	var result Result
	// 	rows.Scan(&result)
	// 	data[i].X = result.Damage
	// 	data[i].Y = result.Control
	// 	i++
	// }

	return data
}
