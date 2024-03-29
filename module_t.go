package browsers

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theMainModuleName     = "github.com/starter-go/browsers"
	theMainModuleVersion  = "v0.0.3"
	theMainModuleRevision = 3
	theMainModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

// ModuleT 创建模块 [github.com/starter-go/browsers]
func ModuleT() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theMainModuleName)
	mb.Version(theMainModuleVersion)
	mb.Revision(theMainModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	mb.Depend(starter.Module())
	return mb
}
