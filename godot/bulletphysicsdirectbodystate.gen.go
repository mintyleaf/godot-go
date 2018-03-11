package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewBulletPhysicsDirectBodyStateFromPointer(ptr gdnative.Pointer) BulletPhysicsDirectBodyState {
func newBulletPhysicsDirectBodyStateFromPointer(ptr gdnative.Pointer) BulletPhysicsDirectBodyState {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BulletPhysicsDirectBodyState{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type BulletPhysicsDirectBodyState struct {
	PhysicsDirectBodyState
	owner gdnative.Object
}

func (o *BulletPhysicsDirectBodyState) BaseClass() string {
	return "BulletPhysicsDirectBodyState"
}

// BulletPhysicsDirectBodyStateImplementer is an interface that implements the methods
// of the BulletPhysicsDirectBodyState class.
type BulletPhysicsDirectBodyStateImplementer interface {
	PhysicsDirectBodyStateImplementer
}