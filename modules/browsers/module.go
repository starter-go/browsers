package browsers

import (
	"github.com/starter-go/application"
	"github.com/starter-go/browsers"
	"github.com/starter-go/browsers/gen/main4browsers"
	"github.com/starter-go/browsers/gen/test4browsers"
)

// Module 导出模块 [github.com/starter-go/browsers]
func Module() application.Module {
	mb := browsers.ModuleT()
	mb.Components(main4browsers.ExportConfig)
	return mb.Create()
}

// ModuleForTest 导出测试模块
func ModuleForTest() application.Module {

	parent := Module()

	mb := browsers.ModuleT()
	mb.Name(parent.Name() + "#test")
	mb.Components(test4browsers.ExportConfig)
	mb.Depend(parent)
	return mb.Create()
}
