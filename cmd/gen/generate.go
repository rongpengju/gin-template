package main

import (
	"fmt"
	"regexp"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/rongpengju/gin-template/configs"
	"github.com/rongpengju/gin-template/dal"
)

func main() {
	generatorForDBWithDefault()
}

func generatorForDBWithDefault() {
	// 指定生成代码的具体相对目录(相对当前文件)，默认为：./query
	// 默认生成需要使用WithContext之后才可以查询的代码，但可以通过设置gen.WithoutContext禁用该模式
	g := gen.NewGenerator(gen.Config{
		OutPath: fmt.Sprintf("dal/%s/query", getDatabaseNameFromDSN(configs.Conf.DataSource.MySQL.DsnWithDefault)),

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,

		// gen_model.WithoutContext：禁用WithContext模式
		// gen_model.WithDefaultQuery：生成一个全局Query对象Q
		// gen_model.WithQueryInterface：生成Query接口
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})

	// 自定义字段的数据类型
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"json":      func(columnType gorm.ColumnType) (dataType string) { return "datatypes.JSON" },
		"datetime":  func(columnType gorm.ColumnType) (dataType string) { return "*time.Time" },
	}
	g.WithDataTypeMap(dataMap)

	// 通常复用项目中已有的SQL连接配置db(*gorm.DB)
	// 非必需，但如果需要复用连接时的gorm.Config或需要连接数据库同步表信息则必须设置
	g.UseDB(dal.DBWithDefault)

	// 从连接的数据库为所有表生成Model结构体和CRUD代码
	g.ApplyBasic(g.GenerateAllTable(gen.FieldType("deleted_at", "gorm.DeletedAt"))...)

	// 指定表名生成Model结构体和CRUD代码
	//demoTable := g.GenerateModelAs("demo", "DemoModel", gen.FieldType("deleted_at", "gorm.DeletedAt"))
	//g.ApplyBasic(demoTable)

	// 执行并生成代码
	g.Execute()
}

// 从 DSN 中提取出数据库的名称
func getDatabaseNameFromDSN(dsn string) string {
	// 使用正则表达式匹配数据库名称
	re := regexp.MustCompile(`/(\w+)\?`)
	matches := re.FindStringSubmatch(dsn)

	if len(matches) < 1 {
		return ""
	}
	return matches[1]
}
