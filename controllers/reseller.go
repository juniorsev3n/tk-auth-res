package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juniorsev3n/td-auth-res/structs"
)

func (idb *InDB) GetReseller(c *gin.Context) {
	var (
		reseller structs.Reseller
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&reseller).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": reseller,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetResellers(c *gin.Context) {
	var (
		reseller []structs.Reseller
		result   gin.H
	)

	idb.DB.Find(&reseller)
	if len(reseller) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": reseller,
			"count":  len(reseller),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreateReseller(c *gin.Context) {
	var (
		reseller structs.Reseller
		result   gin.H
	)
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	reseller.First_Name = first_name
	reseller.Last_Name = last_name
	idb.DB.Create(&reseller)
	result = gin.H{
		"result": reseller,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateReseller(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	var (
		reseller    structs.Reseller
		newReseller structs.Reseller
		result      gin.H
	)

	err := idb.DB.First(&reseller, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newReseller.First_Name = first_name
	newReseller.Last_Name = last_name
	err = idb.DB.Model(&reseller).Updates(newReseller).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteReseller(c *gin.Context) {
	var (
		reseller structs.Reseller
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&reseller, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&reseller).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
