package configures

import (
	"fmt"

	"github.com/mrzack99s/cloud-coco/src/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseConfig struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
}

var db *gorm.DB

func DBInstance() *gorm.DB {
	return db
}

func (conf *DatabaseConfig) SetupDatabase() {

	var err error

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", conf.Username, conf.Password, conf.Hostname, conf.Port, conf.DBName)
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "ccoco_",
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(dsn)

	// Migrate the schema
	migrateDB()

	// Initail offload data
	initOffloadData()
}

func migrateDB() {

	dbInterfaces := []interface{}{
		models.ServiceVersions{},
		models.DirectoriesUsers{},
		models.RolesPermissions{},
		models.RBACSubscriptions{},
		models.RBACResourcePools{},
		models.Directories{},
		models.Subscriptions{},
		models.Users{},
		models.Roles{},
		models.Permissions{},
		models.ResourcePools{},
		models.Resources{},
		models.Services{},
		models.ConfigVariables{},
	}

	for _, element := range dbInterfaces {
		if !db.Migrator().HasTable(&element) {
			db.AutoMigrate(&element)
		}
	}

}

func initOffloadData() {

	if result := db.Find(&models.ConfigVariables{}); result.RowsAffected == 0 {
		data := []models.ConfigVariables{
			{
				Name:  "KUBE_API_HOSTNAME",
				Value: "",
			},
			{
				Name:  "KUBE_API_PATH",
				Value: "",
			},
			{
				Name:  "KUBE_BEARER_TOKEN",
				Value: "",
			},
		}

		for _, ele := range data {
			db.Create(&ele)
		}
	}

	if result := db.Find(&models.Permissions{}); result.RowsAffected == 0 {
		data := []models.Permissions{
			{
				Name: "Create",
			},
			{
				Name: "Read",
			},
			{
				Name: "Update",
			},
			{
				Name: "Delete",
			},
		}

		for _, ele := range data {
			db.Create(&ele)
		}
	}

	if result := db.Find(&models.Roles{}); result.RowsAffected == 0 {
		var create_permission models.Permissions
		var read_permission models.Permissions
		var update_permission models.Permissions
		var delete_permission models.Permissions

		db.Where("name = ?", "Create").Find(&create_permission)
		db.Where("name = ?", "Read").Find(&read_permission)
		db.Where("name = ?", "Update").Find(&update_permission)
		db.Where("name = ?", "Delete").Find(&delete_permission)

		data := []models.Roles{
			{
				Name: "SuperOwner",
				Permissions: []models.RolesPermissions{
					{
						Permission: create_permission,
					},
					{
						Permission: read_permission,
					},
					{
						Permission: update_permission,
					},
					{
						Permission: delete_permission,
					},
				},
			},
			{
				Name: "Owner",
				Permissions: []models.RolesPermissions{
					{
						Permission: create_permission,
					},
					{
						Permission: read_permission,
					},
					{
						Permission: update_permission,
					},
					{
						Permission: delete_permission,
					},
				},
			},
			{
				Name: "Developer",
				Permissions: []models.RolesPermissions{
					{
						Permission: read_permission,
					},
					{
						Permission: update_permission,
					},
				},
			},
			{
				Name: "Viewer",
				Permissions: []models.RolesPermissions{
					{
						Permission: read_permission,
					},
				},
			},
		}

		for _, ele := range data {
			db.Create(&ele)
		}
	}

	if result := db.Find(&models.Directories{}); result.RowsAffected == 0 {
		data := []models.Directories{
			{
				Name: "Administrator",
			},
		}

		for _, ele := range data {
			db.Create(&ele)
		}
	}

	if result := db.Find(&models.ResourcesStatus{}); result.RowsAffected == 0 {
		data := []models.ResourcesStatus{
			{
				Name: "Pending",
			},
			{
				Name: "Running",
			},
			{
				Name: "Stopped",
			},
		}

		for _, ele := range data {
			db.Create(&ele)
		}
	}
}
