package abstractfactory

import "testing"

func TestAbstactFactory(t *testing.T) {
	var factory IFactory
	factory = NewIFactory("haier")
	PrintProducts(factory)
	factory = NewIFactory("hisense")
	PrintProducts(factory)
	factory = NewIFactory("tcl")
	PrintProducts(factory)

}
func PrintProducts(factory IFactory) {
	var tv ITV
	var freezer IFreezer
	var airController IAirController
	tv = factory.produceTV()
	tv.TVSay()
	freezer = factory.produceFreezer()
	freezer.FreezerSay()
	airController = factory.produceAirController()
	airController.AirControllerSay()
}
