package main

import (
	"flag"
	"log"
	"net/http"
	"projeto-magalu-api-go/pkg/api/controller"
	"projeto-magalu-api-go/pkg/api/ingestor"
	"projeto-magalu-api-go/pkg/api/provider/mongo/dao"
	"projeto-magalu-api-go/pkg/api/router"
	"projeto-magalu-api-go/pkg/api/service"
	"projeto-magalu-api-go/pkg/utl/config"
	"projeto-magalu-api-go/pkg/utl/mg"
)

// @title magalu Challenge API
// @version 1.0
// @description API para desafio técnico da Magalu Cloud
// @termsOfService http://swagger.io/terms/
// @contact.name Thiago Menezes
// @contact.email thg.mnzs@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {

	cfgPath := flag.String("p", "./cmd/conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		log.Fatal("error get config file")
	}

	client, database, err := mg.New(cfg.DB.PSN, cfg.DB.DB)
	if err != nil {
		log.Fatal("error get database")
	}

	transactionDAO := dao.NewMongoTransaction(client, database)
	tenantDAO := dao.NewMongoTenant(client, database)
	productDAO := dao.NewMongoProduct(client, database)
	contractDAO := dao.NewMongoContract(client, database)
	pulseDAO := dao.NewMongoPulse(client, database)

	transactionService := service.NewTransaction(transactionDAO)
	tenantService := service.NewTenant(tenantDAO)
	productService := service.NewProduct(productDAO)
	contractService := service.NewContract(contractDAO)
	pulseService := service.NewPulse(pulseDAO)

	transactionHandler := controller.NewTransactionController(transactionService)
	tenantHandler := controller.NewTenantController(tenantService)
	productHandler := controller.NewProductController(productService)
	contractHandler := controller.NewContractController(contractService)
	pulseHandler := controller.NewPulseController(pulseService)
	healthHandler := controller.NewHealthHandler(nil)

	ingestor.StartPulseIngestor(pulseService, cfg)
	routers := router.NewRARouter(transactionHandler, pulseHandler, healthHandler, tenantHandler, productHandler, contractHandler)

	log.Println("Servidor esta rodando na porta " + cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, routers))
}
