package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./dal/query",
		ModelPkgPath: "./dal/model",

		Mode: gen.WithoutContext,

		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	dsn := "postgres://postgres:pg_pass@127.0.0.1:5432/pcr_bot"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModelAs("charactersparties", "CharactersParties"),
		g.GenerateModelAs("arenasolutions", "ArenaSolutions"),
		g.GenerateModelAs("princessarenalineups", "PrincessArenaLineups"),
	)

	g.Execute()
}
