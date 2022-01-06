package starter

import (
	"context"
	"backend1_homework/hw5/reguser/app/user"
	"sync"
)

type App struct {
	us *user.Users
}

func NewApp(ust user.UserStore) *App {
	a := &App{
		us: user.NewUsers(ust),
	}
	return a
}

type HTTPServer interface {
	Start(us *user.Users)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs HTTPServer) {
	defer wg.Done()
	hs.Start(a.us)
	<-ctx.Done()
	hs.Stop()
}