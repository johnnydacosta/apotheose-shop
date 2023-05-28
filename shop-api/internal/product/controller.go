package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

var IN_MEMORY_DB = map[string]Product{
	"0": {
		ID:       "0",
		Title:    "Oscar",
		Price:    23.34,
		Quantity: 2,
	},
	"1": {
		ID:       "1",
		Title:    "Porte Savon",
		Price:    20.34,
		Quantity: 23,
	},
}

type ProductController struct{}

func (p *ProductController) Find(c *gin.Context) {
	c.JSON(http.StatusOK, maps.Values(IN_MEMORY_DB))
}

func (p *ProductController) FindById(c *gin.Context) {
	id, exist := c.Params.Get("id")

	if exist == false {
		c.JSON(http.StatusBadRequest, "No id provided")
		return
	}

	if p, ok := IN_MEMORY_DB[id]; ok {
		c.JSON(http.StatusOK, p)
		return
	}

	c.JSON(http.StatusNotFound, "Product not found")
}

func (p *ProductController) Create(c *gin.Context) {
	id := uuid.New().String()
	product := Product{
		ID:       id,
		Title:    "New Product",
		Price:    20,
		Quantity: 5,
	}

	IN_MEMORY_DB[id] = product

	c.JSON(http.StatusOK, product)
}

func (p *ProductController) Update(c *gin.Context) {
	// TODO update the name
	c.JSON(http.StatusOK, "product updated")
}

func (p *ProductController) DeleteById(c *gin.Context) {
	id, exist := c.Params.Get("id")

	if exist == false {
		c.JSON(http.StatusBadRequest, "No id provided")
		return
	}

	_, found := IN_MEMORY_DB[id]

	if found == false {
		c.JSON(http.StatusNotFound, "Product not found")
		return
	}

	delete(IN_MEMORY_DB, id)

	c.JSON(http.StatusOK, "Product deleted !")
}
