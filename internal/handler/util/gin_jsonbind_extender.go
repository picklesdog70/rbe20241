package util

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindJSON(obj interface{}, c *gin.Context) error {
	return BindWith(obj, binding.JSON, c)
}

func BindWith(obj interface{}, b binding.Binding, c *gin.Context) error {
	if err := b.Bind(c.Request, obj); err != nil {
		return err
	}
	return nil
}
