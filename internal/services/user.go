package services

import "arriba/internal/domain"

func GetUser(ctx *domain.ArribaContext, userID int64) domain.User {
	ctx.Mutex.RLock()
	user := ctx.Users[userID]
	ctx.Mutex.RUnlock()
	return user
}

func UpdateUserAccount(ctx *domain.ArribaContext, userID int64, account domain.Account) {
	if user, ok := ctx.Users[userID]; ok {
		user.Account = account
		ctx.Mutex.Lock()
		ctx.Users[userID] = user
		ctx.Mutex.Unlock()
	}
}
