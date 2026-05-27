package main

import (
	"os"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	"wayne/internal/k8s/client"
	"wayne/internal/router"

	_ "wayne/internal/api/auth/oauth2"

	util "wayne/pkg"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/dgrijalva/jwt-go"

	"database/sql"
	"fmt"
	"strings"

	"wayne/pkg/logger"

	"github.com/beego/beego/v2/adapter/orm"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//初始化日志
	initLog()

	//初始化db
	initDb()

	// 定时更新 K8S Client ，
	go wait.Forever(client.BuildApiserverClient, 5*time.Second)

	// 初始化RsaPrivateKey
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(readKey("RsaPrivateKey"))
	if err != nil {
		panic(err)
	}
	util.RsaPrivateKey = privateKey

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(readKey("RsaPublicKey"))
	if err != nil {
		panic(err)
	}
	util.RsaPublicKey = publicKey

	// init kube labels
	util.AppLabelKey = beego.AppConfig.DefaultString("AppLabelKey", "wayne-app")
	util.NamespaceLabelKey = beego.AppConfig.DefaultString("NamespaceLabelKey", "wayne-ns")

	router.Setup()

	beego.Run()
}

func readKey(key string) []byte {
	filename := beego.AppConfig.String(key)
	// get the abs
	// which will try to find the 'filename' from current workind dir too.
	pem, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	// read the raw contents of the file
	data, err := os.ReadFile(pem)
	if err != nil {
		panic(err)
	}

	return data
}

func initDb() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	err := ensureDatabase()
	if err != nil {
		panic(err)
	}
	db, err := orm.GetDB()

	if err != nil {
		panic(err)
	}

	ttl := beego.AppConfig.DefaultInt("DBConnTTL", 30)

	db.SetConnMaxLifetime(time.Duration(ttl) * time.Second)

	orm.Debug = beego.AppConfig.DefaultBool("ShowSql", false)
}

func ensureDatabase() error {

	dbName := beego.AppConfig.String("DBName")

	dbURL := fmt.Sprintf("%s:%s@%s/", beego.AppConfig.String("DBUser"),
		beego.AppConfig.String("DBPasswd"), beego.AppConfig.String("DBTns"))

	db, err := sql.Open("mysql", fmt.Sprintf("%s%s", dbURL, dbName))

	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		return err
	}

	logger.Debugf("Initialize database connection: %s", strings.Replace(dbURL, beego.AppConfig.String("DBPasswd"), "****", 1))

	err = orm.RegisterDataBase("default", "mysql", addLocation(fmt.Sprintf("%s%s", dbURL, dbName)))
	if err != nil {
		return err
	}

	return nil
}

func addLocation(dbURL string) string {
	return fmt.Sprintf("%s?charset=utf8mb4&loc=%s", dbURL, beego.AppConfig.DefaultString("DBLoc", "Asia%2FShanghai"))
}

func initLog() {

	confMap, err := beego.AppConfig.GetSection("log")

	if err != nil {
		panic(err)
	}

	lev := beego.AppConfig.DefaultInt("log::log_level", 4)
	conf := logger.LogOption{}
	conf.Format = confMap["format"]
	conf.Output = confMap["output"]
	conf.TimeFormat = confMap["time_format"]
	conf.Level = logger.LogLevel(lev)
	logger.CreateLogger(&conf)

}
