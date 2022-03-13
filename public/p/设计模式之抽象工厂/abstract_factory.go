package abstractfactory

type ITV interface {
	TVSay()
}
type IFreezer interface {
	FreezerSay()
}
type IAirController interface {
	AirControllerSay()
}

type IFactory interface {
	produceTV() ITV
	produceFreezer() IFreezer
	produceAirController() IAirController
}

type HaierFactory struct {
}
type HisenseFactory struct {
}
type TCLFactory struct {
}
type HaierTV struct{}
type HaierFreezer struct{}
type HaierAirController struct{}
type HisenseTV struct{}
type HisenseFreezer struct{}
type HisenseAirController struct{}
type TCLTV struct{}
type TCLFreezer struct{}
type TCLAirController struct{}

func (h HaierTV) TVSay() {
	print("I'm Haier TV\n")
}
func (h HaierFreezer) FreezerSay() {
	print("I'm Haier Freezer\n")
}
func (h HaierAirController) AirControllerSay() {
	print("I'm Haier AirController\n")
}
func (h HisenseTV) TVSay() {
	print("I'm Hisense TV\n")
}
func (h HisenseFreezer) FreezerSay() {
	print("I'm Hisense Freezer\n")
}
func (h HisenseAirController) AirControllerSay() {
	print("I'm Hisense AirController\n")
}
func (h TCLTV) TVSay() {
	print("I'm TCL TV\n")
}
func (h TCLFreezer) FreezerSay() {
	print("I'm TCL Freezer\n")
}
func (h TCLAirController) AirControllerSay() {
	print("I'm TCL AirController\n")
}

func (TCLFactory) produceTV() ITV {
	return TCLTV{}
}
func (TCLFactory) produceFreezer() IFreezer {
	return TCLFreezer{}
}
func (TCLFactory) produceAirController() IAirController {
	return TCLAirController{}
}
func (HisenseFactory) produceTV() ITV {
	return HisenseTV{}
}
func (HisenseFactory) produceFreezer() IFreezer {
	return HisenseFreezer{}
}
func (HisenseFactory) produceAirController() IAirController {
	return HisenseAirController{}
}
func (HaierFactory) produceTV() ITV {
	return HaierTV{}
}
func (HaierFactory) produceFreezer() IFreezer {
	return HaierFreezer{}
}
func (HaierFactory) produceAirController() IAirController {
	return HaierAirController{}
}
func NewIFactory(t string) IFactory {
	switch t {
	case "haier":
		return HaierFactory{}
	case "hisense":
		return HisenseFactory{}
	case "tcl":
		return TCLFactory{}
	}
	return nil
}
