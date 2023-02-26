package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() // main context
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	cancel()
	// }()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done(): // caso cancelado antes de 5 segundos, a ação não será executada
		fmt.Println("Tempo exedido para bookar o quarto")
	case <-time.After(time.Second * 5): // caso passe 5 segundos, o quarto será reservado
		fmt.Println("Quarto reservado com sucesso")
	}
}
