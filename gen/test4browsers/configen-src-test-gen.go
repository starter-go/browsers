package test4browsers
import (
    pa1bfd2596 "github.com/starter-go/browsers"
    pa58a1f205 "github.com/starter-go/browsers/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type pa58a1f205.Test1 in package:github.com/starter-go/browsers/src/test/golang/unit
//
// id:com-a58a1f20578de0b7-unit-Test1
// class:
// alias:
// scope:singleton
//
type pa58a1f2057_unit_Test1 struct {
}

func (inst* pa58a1f2057_unit_Test1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a58a1f20578de0b7-unit-Test1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa58a1f2057_unit_Test1) new() any {
    return &pa58a1f205.Test1{}
}

func (inst* pa58a1f2057_unit_Test1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa58a1f205.Test1)
	nop(ie, com)

	
    com.Browsers = inst.getBrowsers(ie)


    return nil
}


func (inst*pa58a1f2057_unit_Test1) getBrowsers(ie application.InjectionExt)pa1bfd2596.Service{
    return ie.GetComponent("#alias-a1bfd25967b1d55d99511da18bc92d78-Service").(pa1bfd2596.Service)
}


