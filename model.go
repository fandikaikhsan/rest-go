package main

// func conn () {

// 	ctx := context.Background()

// 	dsn := "postgres://fandikaikhsan:bandoeng14022@localhost:5432/go?sslmode=disable"
// 	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
// 	db := bun.NewDB(sqldb, pgdialect.New())

// 	album := new(album)
// 	if err := db.NewSelect().Model(album).Where("id = ?", 1).Scan(ctx); err != nil { panic(err) }

// 	fmt.Printf("album: %#v", album)
// 	return
// }