package client

import (
	"context"
	"github.com/rowdyroad/grpc-hw/internal/client/schema"
	"github.com/rowdyroad/grpc-hw/internal/storage"
	"github.com/rowdyroad/grpc-hw/pkg/client"
	"math"
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"
import "github.com/gin-gonic/contrib/cors"

type Config struct {
	Listen string
	Addr string
}

type Client struct {
	config Config
	client *client.Client
	server *http.Server
}

func NewClient(config Config) (*Client, error) {
	cl, err := client.NewClient(config.Addr)
	if err != nil {
		return nil, err
	}
	client := &Client{config: config, client: cl}
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/events", client.handlerEvents)
	router.GET("/events/:time", client.handlerValue)
	router.GET("/stats/daily", client.handlerStatsDaily)

	client.server = &http.Server{
		Addr: config.Listen,
		Handler: router,
	}
	return client, nil
}

func (cl *Client) Run() error {
	if err := cl.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (cl *Client) Close() error {
	return cl.server.Shutdown(context.Background())
}

func (cl *Client) handlerEvents(c *gin.Context) {
	var req schema.ListRequest
	if err := c.BindQuery(&req); err != nil {
		return
	}
	low := math.Inf(-1)
	if req.Low != nil {
		low = *req.Low
	}
	high := math.Inf(1)
	if req.High != nil {
		high = *req.High
	}
	count, err := cl.client.GetTotalCount(req.From,req.To, low, high)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Header("X-Total-Count", strconv.Itoa(count))
	records, err := cl.client.GetList(req.From, req.To, low, high, req.Offset, req.Limit)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, records)
}

func (cl *Client) handlerValue(c *gin.Context) {
	var req schema.ValueRequest
	if err := c.BindUri(&req); err != nil {
		return
	}
	value, err := cl.client.GetValue(req.Time)
	if err == storage.ErrValueNotFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, value)
}

func (cl *Client) handlerStatsDaily(c *gin.Context) {
	var req schema.StatsRequest
	if err := c.BindQuery(&req); err != nil {
		return
	}
	stats, err := cl.client.GetDailyStats(req.From, req.To)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, stats)
}