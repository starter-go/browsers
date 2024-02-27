package main4browsers
import (
    p0d2a11d16 "github.com/starter-go/afs"
    pa1bfd2596 "github.com/starter-go/browsers"
    peec03e62f "github.com/starter-go/browsers/src/main/golang/lib"
     "github.com/starter-go/application"
)

// type peec03e62f.Chrome in package:github.com/starter-go/browsers/src/main/golang/lib
//
// id:com-eec03e62f29f1187-lib-Chrome
// class:class-a1bfd25967b1d55d99511da18bc92d78-Registry
// alias:
// scope:singleton
//
type peec03e62f2_lib_Chrome struct {
}

func (inst* peec03e62f2_lib_Chrome) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-eec03e62f29f1187-lib-Chrome"
	r.Classes = "class-a1bfd25967b1d55d99511da18bc92d78-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* peec03e62f2_lib_Chrome) new() any {
    return &peec03e62f.Chrome{}
}

func (inst* peec03e62f2_lib_Chrome) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*peec03e62f.Chrome)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)
    com.ExePath = inst.getExePath(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Priority = inst.getPriority(ie)


    return nil
}


func (inst*peec03e62f2_lib_Chrome) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*peec03e62f2_lib_Chrome) getExePath(ie application.InjectionExt)string{
    return ie.GetString("${browser.chrome.executable}")
}


func (inst*peec03e62f2_lib_Chrome) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${browser.chrome.enabled}")
}


func (inst*peec03e62f2_lib_Chrome) getPriority(ie application.InjectionExt)int{
    return ie.GetInt("${browser.chrome.priority}")
}



// type peec03e62f.Firefox in package:github.com/starter-go/browsers/src/main/golang/lib
//
// id:com-eec03e62f29f1187-lib-Firefox
// class:class-a1bfd25967b1d55d99511da18bc92d78-Registry
// alias:
// scope:singleton
//
type peec03e62f2_lib_Firefox struct {
}

func (inst* peec03e62f2_lib_Firefox) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-eec03e62f29f1187-lib-Firefox"
	r.Classes = "class-a1bfd25967b1d55d99511da18bc92d78-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* peec03e62f2_lib_Firefox) new() any {
    return &peec03e62f.Firefox{}
}

func (inst* peec03e62f2_lib_Firefox) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*peec03e62f.Firefox)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)
    com.ExePath = inst.getExePath(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Priority = inst.getPriority(ie)


    return nil
}


func (inst*peec03e62f2_lib_Firefox) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*peec03e62f2_lib_Firefox) getExePath(ie application.InjectionExt)string{
    return ie.GetString("${browser.firefox.executable}")
}


func (inst*peec03e62f2_lib_Firefox) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${browser.firefox.enabled}")
}


func (inst*peec03e62f2_lib_Firefox) getPriority(ie application.InjectionExt)int{
    return ie.GetInt("${browser.firefox.priority}")
}



// type peec03e62f.MSEdge in package:github.com/starter-go/browsers/src/main/golang/lib
//
// id:com-eec03e62f29f1187-lib-MSEdge
// class:class-a1bfd25967b1d55d99511da18bc92d78-Registry
// alias:
// scope:singleton
//
type peec03e62f2_lib_MSEdge struct {
}

func (inst* peec03e62f2_lib_MSEdge) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-eec03e62f29f1187-lib-MSEdge"
	r.Classes = "class-a1bfd25967b1d55d99511da18bc92d78-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* peec03e62f2_lib_MSEdge) new() any {
    return &peec03e62f.MSEdge{}
}

func (inst* peec03e62f2_lib_MSEdge) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*peec03e62f.MSEdge)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)
    com.ExePath = inst.getExePath(ie)
    com.Enabled = inst.getEnabled(ie)
    com.Priority = inst.getPriority(ie)


    return nil
}


func (inst*peec03e62f2_lib_MSEdge) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*peec03e62f2_lib_MSEdge) getExePath(ie application.InjectionExt)string{
    return ie.GetString("${browser.msedge.executable}")
}


func (inst*peec03e62f2_lib_MSEdge) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${browser.msedge.enabled}")
}


func (inst*peec03e62f2_lib_MSEdge) getPriority(ie application.InjectionExt)int{
    return ie.GetInt("${browser.msedge.priority}")
}



// type peec03e62f.BrowserServiceImpl in package:github.com/starter-go/browsers/src/main/golang/lib
//
// id:com-eec03e62f29f1187-lib-BrowserServiceImpl
// class:
// alias:alias-a1bfd25967b1d55d99511da18bc92d78-Service
// scope:singleton
//
type peec03e62f2_lib_BrowserServiceImpl struct {
}

func (inst* peec03e62f2_lib_BrowserServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-eec03e62f29f1187-lib-BrowserServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-a1bfd25967b1d55d99511da18bc92d78-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* peec03e62f2_lib_BrowserServiceImpl) new() any {
    return &peec03e62f.BrowserServiceImpl{}
}

func (inst* peec03e62f2_lib_BrowserServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*peec03e62f.BrowserServiceImpl)
	nop(ie, com)

	
    com.Regs = inst.getRegs(ie)


    return nil
}


func (inst*peec03e62f2_lib_BrowserServiceImpl) getRegs(ie application.InjectionExt)[]pa1bfd2596.Registry{
    dst := make([]pa1bfd2596.Registry, 0)
    src := ie.ListComponents(".class-a1bfd25967b1d55d99511da18bc92d78-Registry")
    for _, item1 := range src {
        item2 := item1.(pa1bfd2596.Registry)
        dst = append(dst, item2)
    }
    return dst
}


