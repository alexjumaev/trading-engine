package engine

import (
	"fmt"
	"strconv"
	"time"
)

type event interface {
	getTime() time.Time
	getName() string
	getSymbol() string
	String() string
}

type BaseEvent struct {
	Time   time.Time
	Ticker *Instrument
}

func (c *BaseEvent) getSymbol() string {
	return c.Ticker.Symbol
}

func (c *BaseEvent) getTime() time.Time {
	return c.Time
}

func (c *BaseEvent) getStringTime() string {
	return c.Time.Format("2006-01-02 15:04:05")
}

func (c *BaseEvent) String() string {
	return fmt.Sprintf("%v %v", c.getTime(), "Base Event")
}

func be(datetime time.Time, symbol *Instrument) BaseEvent {
	b := BaseEvent{Time: datetime, Ticker: symbol}
	return b
}

type CandleOpenEvent struct {
	BaseEvent
	CandleTime time.Time
	Price      float64
	TimeFrame  string
}

func (c *CandleOpenEvent) getName() string {
	return "CandleOpenEvent"
}

func (c *CandleOpenEvent) String() string {
	return fmt.Sprintf("%v **%v** Price: %v", c.getStringTime(), c.getName(), c.Price)
}

type CandleCloseEvent struct {
	BaseEvent
	Candle    *Candle
	TimeFrame string
}

func (e *CandleCloseEvent) setEventTimeFromCandle() {
	if e.Candle == nil {
		panic("Can't get Event time from Nil candle")
	}
	switch e.TimeFrame {
	case "D":
		closeTime := e.Candle.Datetime
		closeTime = time.Date(closeTime.Year(), closeTime.Month(), closeTime.Day(), 23, 59, 59, 0,
			e.Candle.Datetime.Location())
		e.BaseEvent.Time = closeTime
	case "W":
		closeTime := e.Candle.Datetime.AddDate(0, 0, 7)
		closeTime = time.Date(closeTime.Year(), closeTime.Month(), closeTime.Day(), 23, 59, 59, 0,
			e.Candle.Datetime.Location())
		e.BaseEvent.Time = closeTime
	default:
		minutes, err := strconv.ParseInt(e.TimeFrame, 10, 8)
		if err != nil {
			panic("Unknown timeframe: " + e.TimeFrame)
		}
		closeTime := e.Candle.Datetime.Add(time.Duration(minutes) * time.Minute)
		e.BaseEvent.Time = closeTime
	}
}

func (c *CandleCloseEvent) getName() string {
	return "CandleCloseEvent"
}

func (c *CandleCloseEvent) String() string {
	return fmt.Sprintf("%v **%v** Candle: %+v", c.getStringTime(), c.getName(), c.Candle)
}

type CandlesHistoryEvent struct {
	BaseEvent
	Candles CandleArray
}

func (c *CandlesHistoryEvent) getName() string {
	return "CandlesHistoryEvent"
}

func (c *CandlesHistoryEvent) String() string {
	return fmt.Sprintf("%v **%v** Total candles: %v", c.getStringTime(), c.getName(), len(c.Candles))
}

type NewTickEvent struct {
	BaseEvent
	Tick *Tick
}

func (c *NewTickEvent) getName() string {
	return "NewTickEvent"
}

func (c *NewTickEvent) String() string {
	return fmt.Sprintf("%v **%v** Tick: %+v", c.getStringTime(), c.getName(), c.Tick)
}

type TickHistoryEvent struct {
	BaseEvent
	Ticks TickArray
}

func (c *TickHistoryEvent) getName() string {
	return "TickHistoryEvent"
}

func (c *TickHistoryEvent) String() string {
	return fmt.Sprintf("%v **%v** Total ticks: %v", c.getStringTime(), c.getName(), len(c.Ticks))
}

type NewOrderEvent struct {
	BaseEvent
	LinkedOrder *Order
}

func (c *NewOrderEvent) getName() string {
	return "NewOrderEvent"
}

func (c *NewOrderEvent) String() string {
	return fmt.Sprintf("%v **%v** Order: %+v", c.getStringTime(), c.getName(), c.LinkedOrder)
}

type OrderConfirmationEvent struct {
	BaseEvent
	OrdId string
}

func (c *OrderConfirmationEvent) getName() string {
	return "OrderConfirmationEvent"
}

func (c *OrderConfirmationEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderID: %v", c.getStringTime(), c.getName(), c.OrdId)
}

type OrderFillEvent struct {
	BaseEvent
	OrdId string
	Price float64
	Qty   int64
}

func (c *OrderFillEvent) getName() string {
	return "OrderFillEvent"
}

func (c *OrderFillEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderID: %v Price: %v Qty: %v", c.getStringTime(), c.getName(), c.OrdId,
		c.Price, c.Qty)
}

type OrderCancelEvent struct {
	BaseEvent
	OrdId string
}

func (c *OrderCancelEvent) getName() string {
	return "OrderCancelEvent"
}

func (c *OrderCancelEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderID: %v", c.getStringTime(), c.getName(), c.OrdId)
}

type OrderCancelRejectEvent struct {
	BaseEvent
	OrdId  string
	Reason string
}

func (c *OrderCancelRejectEvent) getName() string {
	return "OrderCancelRejectEvent"
}

func (c *OrderCancelRejectEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderId: %v Reason: %+v", c.getStringTime(), c.getName(), c.OrdId, c.Reason)
}

type OrderCancelRequestEvent struct {
	BaseEvent
	OrdId string
}

func (c *OrderCancelRequestEvent) getName() string {
	return "OrderCancelRequestEvent"
}

func (c *OrderCancelRequestEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderID: %v", c.getStringTime(), c.getName(), c.OrdId)
}

type OrderReplaceRequestEvent struct {
	BaseEvent
	OrdId    string
	NewPrice float64
}

func (c *OrderReplaceRequestEvent) getName() string {
	return "OrderReplaceRequestEvent"
}

func (c *OrderReplaceRequestEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderId:%v New Price: %v", c.getStringTime(), c.getName(), c.OrdId, c.NewPrice)
}

type OrderReplaceRejectEvent struct {
	BaseEvent
	OrdId  string
	Reason string
}

func (c *OrderReplaceRejectEvent) getName() string {
	return "OrderReplaceRejectEvent"
}

func (c *OrderReplaceRejectEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderId: %v Reason: %v", c.getStringTime(), c.getName(), c.OrdId, c.Reason)
}

type OrderReplacedEvent struct {
	BaseEvent
	OrdId    string
	NewPrice float64
}

func (c *OrderReplacedEvent) getName() string {
	return "OrderReplacedEvent"
}

func (c *OrderReplacedEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderId: %v New Price: %v", c.getStringTime(), c.getName(), c.OrdId, c.NewPrice)
}

type OrderRejectedEvent struct {
	BaseEvent
	OrdId  string
	Reason string
}

func (c *OrderRejectedEvent) getName() string {
	return fmt.Sprintf("OrderRejectedEvent: %v Reason: %v", c.OrdId, c.Reason)
}

func (c *OrderRejectedEvent) String() string {
	return fmt.Sprintf("%v **%v** OrderId: %v Reason: %v", c.getStringTime(), c.getName(), c.OrdId, c.Reason)
}

type StrategyRequestNotDeliveredEvent struct {
	BaseEvent
	Request event
}

func (c *StrategyRequestNotDeliveredEvent) getName() string {
	return fmt.Sprintf("StrategyRequestNotDeliveredEvent: Reason: %+v", c.Request)
}

func (c *StrategyRequestNotDeliveredEvent) String() string {
	return fmt.Sprintf("%v **%v** Request: %+v", c.getStringTime(), c.getName(), c.Request)
}

type TimerTickEvent struct {
	BaseEvent
}

func (c *TimerTickEvent) getName() string {
	return "TimerTickEvent"
}

type EndOfDataEvent struct {
	BaseEvent
}

func (c *EndOfDataEvent) getName() string {
	return "EndOfDataEvent"
}

func (c *EndOfDataEvent) String() string {
	return fmt.Sprintf("%v %v", c.getTime(), c.getName())
}


type PortfolioNewPositionEvent struct {
	BaseEvent
	trade *Trade
}

func (c *PortfolioNewPositionEvent) getName() string {
	return "PortfolioNewPositionEvent"
}

func (c *PortfolioNewPositionEvent) String() string {
	return fmt.Sprintf("%v **%v** Trade: %+v", c.getStringTime(), c.getName(), c.trade.Id)
}

type StrategyFinishedEvent struct {
	BaseEvent
	strategy string
}

func (c *StrategyFinishedEvent) getName() string {
	return "StrategyFinishedEvent"
}

func (c *StrategyFinishedEvent) String() string {
	return fmt.Sprintf("%v **%v** Strategy: %+v", c.getStringTime(), c.getName(), c.strategy)
}
