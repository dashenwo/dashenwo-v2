package gorm

import "C"
import (
	"fmt"
	conf "github.com/dashenwo/dashenwo/v2/console/captcha/config"
	"github.com/dashenwo/dashenwo/v2/console/captcha/internal/model"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"sync"

	// gorm驱动注入
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
	"time"
)

var (
	db     *gorm.DB
	once   sync.Once
	dbConf conf.Database
)

// 新建数据库连接
func InitDb() {
	once.Do(func() {
		dbConf = conf.Database{}
		err := config.Get("database").Scan(&dbConf)
		if err != nil {
			log.Fatal(err)
		}
		sqlConnection := dbConf.User + ":" + dbConf.Password + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/" + dbConf.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(dbConf.Engine, sqlConnection)
		if err != nil {
			log.Fatal(err)
		}
		if dbConf.LogMode {
			db = db.Debug()
		}

		err = db.DB().Ping()
		if err != nil {
			log.Fatal(err)
		}

		db.DB().SetMaxIdleConns(dbConf.MaxIdleConns)
		db.DB().SetMaxOpenConns(dbConf.MaxOpenConns)
		db.DB().SetConnMaxLifetime(dbConf.ConnMaxLifetime)
		db.SingularTable(true)
		// 创建时时间钩子
		db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		// 修改时时间的钩子
		db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
		// 删除时
		db.Callback().Delete().Replace("gorm:delete", deleteCallback)
		// 表映射
		if dbConf.AutoMigrate {
			err = AutoMigrate(db)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

// AutoMigrate 自动映射数据表
func AutoMigrate(db *gorm.DB) error {
	if dbType := dbConf.Engine; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}
	return db.AutoMigrate(
		new(model.Captcha), //验证码信息
	).Error
}

// 注册新建钩子在持久化之前
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedTime"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedTime"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 注册更新钩子在持久化之前
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("UpdatedTime", time.Now().Unix())
	}
}

// 注册删除钩子在删除之前
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedTime")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
