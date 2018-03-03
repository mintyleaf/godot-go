package particles2d

import (
	"log"
	"reflect"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
2D particle node used to create a variety of particle systems and effects. [code]Particles2D[/code] features an emitter that generates some number of particles at a given rate. Use the [code]process_material[/code] property to add a [ParticlesMaterial] to configure particle appearance and behavior. Alternatively, you can add a [ShaderMaterial] which will be applied to all particles.
*/
type Particles2D struct {
	Node2D
}

func (o *Particles2D) BaseClass() string {
	return "Particles2D"
}

/*

 */
func (o *Particles2D) CaptureRect() *Rect2 {
	log.Println("Calling Particles2D.CaptureRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "capture_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetAmount() gdnative.Int {
	log.Println("Calling Particles2D.GetAmount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_amount", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetDrawOrder() gdnative.Int {
	log.Println("Calling Particles2D.GetDrawOrder()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_draw_order", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetExplosivenessRatio() gdnative.Float {
	log.Println("Calling Particles2D.GetExplosivenessRatio()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_explosiveness_ratio", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetFixedFps() gdnative.Int {
	log.Println("Calling Particles2D.GetFixedFps()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_fixed_fps", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetFractionalDelta() gdnative.Bool {
	log.Println("Calling Particles2D.GetFractionalDelta()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_fractional_delta", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetHFrames() gdnative.Int {
	log.Println("Calling Particles2D.GetHFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_h_frames", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetLifetime() gdnative.Float {
	log.Println("Calling Particles2D.GetLifetime()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_lifetime", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetNormalMap() *Texture {
	log.Println("Calling Particles2D.GetNormalMap()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_normal_map", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetOneShot() gdnative.Bool {
	log.Println("Calling Particles2D.GetOneShot()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_one_shot", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetPreProcessTime() gdnative.Float {
	log.Println("Calling Particles2D.GetPreProcessTime()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_pre_process_time", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetProcessMaterial() *Material {
	log.Println("Calling Particles2D.GetProcessMaterial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_process_material", goArguments, "*Material")

	returnValue := goRet.Interface().(*Material)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetRandomnessRatio() gdnative.Float {
	log.Println("Calling Particles2D.GetRandomnessRatio()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_randomness_ratio", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetSpeedScale() gdnative.Float {
	log.Println("Calling Particles2D.GetSpeedScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_speed_scale", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetTexture() *Texture {
	log.Println("Calling Particles2D.GetTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetUseLocalCoordinates() gdnative.Bool {
	log.Println("Calling Particles2D.GetUseLocalCoordinates()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_use_local_coordinates", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetVFrames() gdnative.Int {
	log.Println("Calling Particles2D.GetVFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_v_frames", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) GetVisibilityRect() *Rect2 {
	log.Println("Calling Particles2D.GetVisibilityRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_visibility_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Particles2D) IsEmitting() gdnative.Bool {
	log.Println("Calling Particles2D.IsEmitting()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_emitting", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Particles2D) Restart() {
	log.Println("Calling Particles2D.Restart()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "restart", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetAmount(amount gdnative.Int) {
	log.Println("Calling Particles2D.SetAmount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_amount", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetDrawOrder(order gdnative.Int) {
	log.Println("Calling Particles2D.SetDrawOrder()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(order)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_draw_order", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetEmitting(emitting gdnative.Bool) {
	log.Println("Calling Particles2D.SetEmitting()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(emitting)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_emitting", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetExplosivenessRatio(ratio gdnative.Float) {
	log.Println("Calling Particles2D.SetExplosivenessRatio()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ratio)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_explosiveness_ratio", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetFixedFps(fps gdnative.Int) {
	log.Println("Calling Particles2D.SetFixedFps()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(fps)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_fixed_fps", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetFractionalDelta(enable gdnative.Bool) {
	log.Println("Calling Particles2D.SetFractionalDelta()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_fractional_delta", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetHFrames(frames gdnative.Int) {
	log.Println("Calling Particles2D.SetHFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(frames)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_h_frames", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetLifetime(secs gdnative.Float) {
	log.Println("Calling Particles2D.SetLifetime()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(secs)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_lifetime", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetNormalMap(texture *Texture) {
	log.Println("Calling Particles2D.SetNormalMap()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_normal_map", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetOneShot(secs gdnative.Bool) {
	log.Println("Calling Particles2D.SetOneShot()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(secs)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_one_shot", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetPreProcessTime(secs gdnative.Float) {
	log.Println("Calling Particles2D.SetPreProcessTime()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(secs)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_pre_process_time", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetProcessMaterial(material *Material) {
	log.Println("Calling Particles2D.SetProcessMaterial()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(material)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_process_material", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetRandomnessRatio(ratio gdnative.Float) {
	log.Println("Calling Particles2D.SetRandomnessRatio()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ratio)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_randomness_ratio", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetSpeedScale(scale gdnative.Float) {
	log.Println("Calling Particles2D.SetSpeedScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(scale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_speed_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetTexture(texture *Texture) {
	log.Println("Calling Particles2D.SetTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetUseLocalCoordinates(enable gdnative.Bool) {
	log.Println("Calling Particles2D.SetUseLocalCoordinates()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_local_coordinates", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetVFrames(frames gdnative.Int) {
	log.Println("Calling Particles2D.SetVFrames()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(frames)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_v_frames", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Particles2D) SetVisibilityRect(aabb *Rect2) {
	log.Println("Calling Particles2D.SetVisibilityRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(aabb)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_visibility_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Particles2DImplementer is an interface for Particles2D objects.
*/
type Particles2DImplementer interface {
	Class
}