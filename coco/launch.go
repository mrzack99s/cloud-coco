package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mrzack99s/cloud-coco/docs"
	"github.com/mrzack99s/cloud-coco/src/apis"
	"github.com/mrzack99s/cloud-coco/src/configures"
	"github.com/mrzack99s/cloud-coco/src/utils"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Cloud COCO
// @version 0.1.0
// @description This is a Cloud COCO API

// @license.name Apache License Version 2.0
// @license.url https://github.com/mrzack99s/cloud-coco

// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-access-token
func main() {

	config_filename := "./config.yaml"
	// Init config
	if err := utils.FindFileInDirectory("./", "config.dev.yaml"); err == nil {
		config_filename = "./config.dev.yaml"
	}

	confFile, _ := ioutil.ReadFile(config_filename)
	fmt.Println(string(confFile))
	configures.Parse(confFile)

	fmt.Println(configures.Sys().COCO.DB)
	// Setup DB
	configures.Sys().COCO.DB.SetupDatabase()

	// Setup Redis Cache
	configures.Sys().COCO.RedisCache.SetupCache()

	// Cleanup cache
	utils.RedisDeleteWithPrefix("token:access/*")

	// // Create kube config
	// var kubehostname models.ConfigVariables
	// err := services.GetWithCondition(
	// 	configures.DBInstance().Where("name = ?", "KUBE_API_HOSTNAME"),
	// 	&kubehostname,
	// )
	// if err != nil {
	// 	kubehostname.Name = "KUBE_API_HOSTNAME"
	// 	kubehostname.Value = "https://172.20.0.1/k8s/clusters/c-f756c"
	// 	services.Update(&kubehostname)
	// }

	// var kubepath models.ConfigVariables
	// err = services.GetWithCondition(
	// 	configures.DBInstance().Where("name = ?", "KUBE_API_PATH"),
	// 	&kubepath,
	// )
	// if err != nil {
	// 	kubepath.Name = "KUBE_API_HOSTNAME"
	// 	kubepath.Value = ""
	// 	services.Update(&kubehostname)
	// }

	// var kubetoken models.ConfigVariables
	// err = services.GetWithCondition(
	// 	configures.DBInstance().Where("name = ?", "KUBE_BEARER_TOKEN"),
	// 	&kubetoken,
	// )
	// if err != nil {
	// 	kubetoken.Name = "KUBE_API_HOSTNAME"
	// 	kubetoken.Value = "kubeconfig-u-gz9nckvsh4:mh8bldcbxs6cbbcx9bcmxgqn2f2rjc85xddbkv6bk67zhrq7dpdwpf"
	// 	services.Update(&kubehostname)
	// }

	// config := rest.Config{
	// 	Host:        kubehostname.Value.(string),
	// 	APIPath:     kubepath.Value.(string),
	// 	BearerToken: kubetoken.Value.(string),
	// 	TLSClientConfig: rest.TLSClientConfig{
	// 		Insecure: true,
	// 	},
	// }
	// // creates the clientset
	// clientset, _ := kubernetes.NewForConfig(&config)
	// // access the API to list pods
	// pods, _ := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})

	// fmt.Println(pods)

	r := gin.Default()
	r.Use(utils.CORSMiddleware)
	api := r.Group("/api")
	apis.NewDirectoriesController(api)
	apis.NewSubscriptionsController(api)
	apis.NewUsersController(api)
	apis.NewRolesController(api)
	apis.NewResourceController(api)
	apis.NewResourcePoolsController(api)
	apis.NewServicesController(api)
	apis.NewServiceVersionsController(api)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/docs/index.html")
	})

	r.Run(":8000")
}
