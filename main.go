package main

import (
	"context"
	"entdemo/ent"
	"fmt"
	"log"

	_ "entdemo/ent/runtime"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ctx := context.Background()

	user, err := client.User.Create().SetName("a8m").Save(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}

	tweet, err := client.Tweet.Create().SetText("good morning").Save(ctx)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if err := client.Tweet.UpdateOne(tweet).AddLikedUsers(user).Exec(ctx); err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println("aaa")
}
