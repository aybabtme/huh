package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/charmbracelet/huh/spinner"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := spinner.New().
		Context(ctx).
		Action(func() error {
			time.Sleep(time.Minute)
			return nil
		}).
		Accessible(rand.Int()%2 == 0).
		Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done!")
}