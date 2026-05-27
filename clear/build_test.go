package clear

import (
	// "strings"

	"testing"

	"gorm.io/gen"
)

// k8s 升级版本是，需要清洗 数据库资源。配置yaml 结构
func TestBuild(t *testing.T) {

	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath: "./dal/query",
		// 默认情况下会跟随OutPath参数，在同目录下生成model目录
		ModelPkgPath: "model",

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		// https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: false,

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false,
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true,
	})
	// reuse your gorm db
	g.UseDB(DataBase)

	//指定生成表
	// g.ApplyBasic(
	// 	g.GenerateModelAs("employee_log", "LoginLog"),
	// 	g.GenerateModelAs("Employee", "User"),
	// 	g.GenerateModelAs("t_app", "App"),
	// 	g.GenerateModelAs("t_sys_api", "AppApi"),
	// 	g.GenerateModelAs("t_sys_menu", "SysMenu"),
	// 	g.GenerateModelAs("t_sys_menu_map_api", "MenuMapApi"),
	// 	g.GenerateModelAs("t_sys_user_meun", "UserSysMenu"),
	// )

	//生成全表
	g.ApplyBasic(g.GenerateAllTable()...)

	// Generate the code
	g.Execute()
}
