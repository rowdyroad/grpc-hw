package client

import (
	"context"
	_ "github.com/rowdyroad/grpc-hw/docs/client"
	"github.com/rowdyroad/grpc-hw/internal/client/schema"
	"github.com/rowdyroad/grpc-hw/internal/storage"
	"github.com/rowdyroad/grpc-hw/pkg/client"
	"math"
	"net/http"
	"strconv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/records", client.handlerRecords)
	router.GET("/records/:time", client.handleRecordValue)
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

// GetRecords godoc
// @Description Get records
// @Success 200 {object} []storage.Record
// @Tags Record
// @Produce json
// @Router /records [get]
// @Param from query string true "From"
// @Param to query string true "To"
// @Param low query number true "Low"
// @Param high query number true "High"
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
func (cl *Client) handlerRecords(c *gin.Context) {
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

// GetRecordValue godoc
// @Description Get record by time
// @Success 200 {object} number
// @Tags Record
// @Produce json
// @Router /records/{time} [get]
// @Param time path string true "Time"
func (cl *Client) handleRecordValue(c *gin.Context) {
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

// GetStats godoc
// @Description Get daily stats
// @Success 200 {object} map[time.Time]storage.Stat
// @Tags Stats
// @Produce json
// @Router /stats/daily [get]
// @Param from query string true "From"
// @Param to query string true "To"
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