package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/picklesdog70/rbe-api/pkg"
	"github.com/picklesdog70/rbe-api/pkg/config"
)

var dependencyTree *pkg.Dependency

func init() {
	dependency, err := pkg.NewDependency()

	if err != nil {
		panic(fmt.Errorf("erro ao configurar a árvore de dependências: %w", err))
	}

	dependencyTree = dependency
}

func main() {

	port := config.GetEnv("APP_PORT", "8080")

	router := gin.Default()
	router.GET("/clientes/:id/extrato", dependencyTree.ClienteExtratoHandler.GetExtratoByClientId)
	router.POST("/clientes/:id/transacoes", dependencyTree.TransacaoPostHandler.PostTransacao)
	router.Run(":" + port)
}
