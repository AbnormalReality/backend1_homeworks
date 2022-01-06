package main

import (
	"context"
	"backend1_homework/hw5/reguser/api/handler"
	"backend1_homework/hw5/reguser/api/routergin"
	"backend1_homework/hw5/reguser/api/server"
	"backend1_homework/hw5/reguser/app/user"
	"backend1_homework/hw5/reguser/app/starter"
	"backend1_homework/hw5/reguser/usermemstore"
	"os"
	"os/signal"
	"sync"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	ust := usermemstore.NewUsers()
	a := starter.NewApp(ust)
	us := user.NewUsers(ust)
	h := handler.NewHandlers(us)

	rh := routergin.NewRouterGin(h)

	srv := server.NewServer(":8000", rh)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
