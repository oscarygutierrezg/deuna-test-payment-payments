package testing

import (
	"github.com/stretchr/testify/assert"
	"payment-payments-api/internal/models"
	"testing"
)

//func TestOrderStateCancelled(t *testing.T) {
//	assert := assert.New(t)
//
//	var order, _ = models.NewOrder(models.Order{})
//
//	assert.Equal(models.OrderAvailable, order.State)
//	assert.Len(order.StateHistory, 1)
//	assert.Equal(models.OrderAvailable, order.StateHistory[0].State)
//
//	_, orderCancelledErr := order.SetNextState(models.OrderCancelled)
//	assert.Nil(orderCancelledErr)
//	assert.Equal(models.OrderCancelled, order.State)
//	assert.Len(order.StateHistory, 2)
//	assert.Equal(models.OrderAvailable, order.StateHistory[0].State)
//	assert.Equal(models.OrderCancelled, order.StateHistory[1].State)
//}

func TestOrderStateHappyWay(t *testing.T) {
	assert := assert.New(t)

	var order, _ = models.NewOrder(models.Order{})

	assert.Equal(models.OrderAvailable, order.State)
	assert.Len(order.StateHistory, 1)
	assert.Equal(models.OrderAvailable, order.StateHistory[0].State)

	order.State = models.OrderAssigned

	_, orderPickedUpErr := order.SetNextState(models.OrderTypeOrigin, models.OrderReqSuccess)
	assert.Nil(orderPickedUpErr)
	assert.Equal(models.OrderPickedUp, order.State)
	assert.Len(order.StateHistory, 2)
	assert.Equal(models.OrderAvailable, order.StateHistory[0].State)
	assert.Equal(models.OrderPickedUp, order.StateHistory[1].State)

	_, orderDeliveredErr := order.SetNextState(models.OrderTypeDestination, models.OrderReqSuccess)
	assert.Nil(orderDeliveredErr)
	assert.Equal(models.OrderDelivered, order.State)
	assert.Len(order.StateHistory, 3)
	assert.Equal(models.OrderAvailable, order.StateHistory[0].State)
	assert.Equal(models.OrderPickedUp, order.StateHistory[1].State)
	assert.Equal(models.OrderDelivered, order.StateHistory[2].State)
}
