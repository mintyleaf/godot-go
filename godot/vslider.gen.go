package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVSliderFromPointer(ptr gdnative.Pointer) VSlider {
func newVSliderFromPointer(ptr gdnative.Pointer) VSlider {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VSlider{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Vertical slider. See [Slider]. This one goes from bottom (min) to top (max).
*/
type VSlider struct {
	Slider
	owner gdnative.Object
}

func (o *VSlider) BaseClass() string {
	return "VSlider"
}

// VSliderImplementer is an interface that implements the methods
// of the VSlider class.
type VSliderImplementer interface {
	SliderImplementer
}
