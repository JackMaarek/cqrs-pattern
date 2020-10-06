package services

import (
	"fmt"
	"github.com/JackMaarek/cqrsPattern/application/events"
	"github.com/JackMaarek/cqrsPattern/application/util"
	"github.com/JackMaarek/cqrsPattern/chore/cqrs"
	"github.com/JackMaarek/cqrsPattern/domain"
	"github.com/JackMaarek/cqrsPattern/domain/event"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TopupAccount(c *gin.Context) {
	userId := util.ParseStringToUint64(c.Param("id"))
	var amount uint64
	amount = 300
	topupEvent := events.TopUp{
		UserID: userId,
		Amount: amount,
	}
	command := cqrs.NewCommandMessage(&event.CreateUserTopupAccountCommand{TopupEvent: &topupEvent})
	_, err := domain.Cb.Dispatch(command)
	if err != nil {
		fmt.Println(err)
	}
}

func GetAccountHistory(c *gin.Context) {
	var topup events.TopUp
	query := cqrs.NewQueryMessage(&event.FindTopupEventQuery{Topup: topup})
	_, _ = domain.Qb.Dispatch(query)
	fmt.Println(&topup)
	c.JSON(http.StatusOK, "")
	return
}
