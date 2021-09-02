package api

import (
	"net/http"

	"github.com/tonymontanapaffpaff/timescale-postgres-load-testing/pkg/data"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// VisitAPI allows transferring handlers to a router.
type VisitAPI struct {
	data *data.VisitData
}

// NewVisitAPI creates new VisitAPI instance.
func NewVisitAPI(data *data.VisitData) VisitAPI {
	return VisitAPI{data: data}
}

// ServeVisitResource registers new request handles.
func ServeVisitResource(r *gin.Engine, data data.VisitData) {
	api := &VisitAPI{data: &data}
	r.GET("/get", api.Get)
	r.GET("/getCount", api.GetCount)
	r.PUT("/update", api.Update)
	r.DELETE("/delete", api.Delete)
}

//COUNT http://localhost:8080/getCount?field=id
// Update  http://localhost:8080/update?field=id&value=1234&toId=10&fromId=10000
// Delete  http://localhost:8080/delete?id=121321

// Get returns list of visits by passing time interval and spaceId
// or an error message if it failed.
func (a VisitAPI) Get(c *gin.Context) {
	log.Info("Method Get starts work")
	from := c.Query("from")
	to := c.Query("to")
	spaceId := c.Query("spaceId")
	visits, err := a.data.Read(from, to, spaceId)
	log.Info(visits)
	if err != nil {
		log.Errorf("Can't get list of visits, error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "got an error when tried to get list of visits",
		})
		return
	}
	c.JSON(http.StatusOK, visits)
}

func (a VisitAPI) GetCount(c *gin.Context) {
	log.Info("Method GetCount starts work")
	field := c.Query("field")
	visits, err := a.data.GetCount(field)
	if err != nil {
		log.Errorf("Can't get list of visits, error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "got an error when tried to get list of visits",
		})
		return
	}
	c.JSON(http.StatusOK, visits)
}

func (a VisitAPI) Update(c *gin.Context) {
	log.Info("Method Update starts work")
	field := c.Query("field")
	value := c.Query("value")
	toId := c.Query("toId")
	fromId := c.Query("fromId")
	err := a.data.Update(field, value, toId, fromId)
	if err != nil {
		log.Errorf("Can't get list of visits, error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "got an error when tried to update data",
		})
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (a VisitAPI) Delete(c *gin.Context) {
	log.Info("Method Delete starts work")
	id := c.Query("id")
	err := a.data.Delete(id)
	if err != nil {
		log.Errorf("Can't get list of visits, error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "got an error when tried to delete data",
		})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
