package initialize

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

const (
	gormSpanKey        = "__gorm_span"
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
)

type OpentracingPlugin struct{}

var _ gorm.Plugin = &OpentracingPlugin{}

func (op *OpentracingPlugin) Name() string {
	return "opentracingPlugin"
}

func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error) {
	// 开始前
	db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后
	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

func before(db *gorm.DB) {

	if !opentracing.IsGlobalTracerRegistered() {
		return
	}

	if db.Statement == nil || db.Statement.Schema == nil {
		return
	}

	operationName := fmt.Sprintf("Mysql - %s", db.Statement.Schema.Table)

	span, _ := opentracing.StartSpanFromContext(db.Statement.Context, operationName)

	span.SetTag(string(ext.DBType), "sql")
	span.SetTag("db.table", db.Statement.Schema.Table)

	a, ok := db.Statement.Config.Dialector.(*mysql.Dialector)
	if ok {
		index := strings.Index(a.DSN, "tcp(")
		span.SetTag(string(ext.DBInstance), a.DSN[index:])
	}

	// 记录当前span
	db.InstanceSet(gormSpanKey, span)
	return
}

func after(db *gorm.DB) {
	// 从GORM的DB实例中取出span
	_span, isExist := db.InstanceGet(gormSpanKey)
	if !isExist {
		return
	}

	// 断言进行类型转换
	span, ok := _span.(opentracing.Span)
	if !ok {
		return
	}
	defer span.Finish()

	// Error
	if db.Error != nil {
		span.LogFields(tracerLog.Error(db.Error))
	}

	// sql
	span.SetTag(string(ext.DBStatement), db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...))
	span.LogFields(tracerLog.String("sql", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
	span.SetTag("db.count", db.RowsAffected)
	span.SetTag("db.method", strings.ToUpper(strings.Split(db.Statement.SQL.String(), " ")[0]))
	return
}
