package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx, "Del Mar")
}

// por convenção, o context deve vir sempre primeiro
func bookHotel(ctx context.Context, nome string) {
	token := ctx.Value("token")
	fmt.Println(token)

}
