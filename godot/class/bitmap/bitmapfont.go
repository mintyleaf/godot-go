package bitmap

import (
	"log"
	"reflect"

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

/*
Renders text using [code]*.fnt[/code] fonts containing texture atlases. Supports distance fields. For using vector font files like TTF directly, see [DynamicFont].
*/
type BitmapFont struct {
	Font
}

func (o *BitmapFont) BaseClass() string {
	return "BitmapFont"
}

/*
   Undocumented
*/
func (o *BitmapFont) X_GetChars() *PoolIntArray {
	log.Println("Calling BitmapFont.X_GetChars()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_get_chars", goArguments, "*PoolIntArray")

	returnValue := goRet.Interface().(*PoolIntArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *BitmapFont) X_GetKernings() *PoolIntArray {
	log.Println("Calling BitmapFont.X_GetKernings()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_get_kernings", goArguments, "*PoolIntArray")

	returnValue := goRet.Interface().(*PoolIntArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *BitmapFont) X_GetTextures() *Array {
	log.Println("Calling BitmapFont.X_GetTextures()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_get_textures", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *BitmapFont) X_SetChars(arg0 *PoolIntArray) {
	log.Println("Calling BitmapFont.X_SetChars()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_chars", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *BitmapFont) X_SetKernings(arg0 *PoolIntArray) {
	log.Println("Calling BitmapFont.X_SetKernings()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_kernings", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *BitmapFont) X_SetTextures(arg0 *Array) {
	log.Println("Calling BitmapFont.X_SetTextures()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_textures", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a character to the font, where [code]character[/code] is the unicode value, [code]texture[/code] is the texture index, [code]rect[/code] is the region in the texture (in pixels!), [code]align[/code] is the (optional) alignment for the character and [code]advance[/code] is the (optional) advance.
*/
func (o *BitmapFont) AddChar(character gdnative.Int, texture gdnative.Int, rect *Rect2, align *Vector2, advance gdnative.Float) {
	log.Println("Calling BitmapFont.AddChar()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(character)
	goArguments[1] = reflect.ValueOf(texture)
	goArguments[2] = reflect.ValueOf(rect)
	goArguments[3] = reflect.ValueOf(align)
	goArguments[4] = reflect.ValueOf(advance)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_char", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a kerning pair to the [code]BitmapFont[/code] as a difference. Kerning pairs are special cases where a typeface advance is determined by the next character.
*/
func (o *BitmapFont) AddKerningPair(charA gdnative.Int, charB gdnative.Int, kerning gdnative.Int) {
	log.Println("Calling BitmapFont.AddKerningPair()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(charA)
	goArguments[1] = reflect.ValueOf(charB)
	goArguments[2] = reflect.ValueOf(kerning)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_kerning_pair", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a texture to the [code]BitmapFont[/code].
*/
func (o *BitmapFont) AddTexture(texture *Texture) {
	log.Println("Calling BitmapFont.AddTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(texture)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_texture", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Clears all the font data and settings.
*/
func (o *BitmapFont) Clear() {
	log.Println("Calling BitmapFont.Clear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Creates a BitmapFont from the [code]*.fnt[/code] file at [code]path[/code].
*/
func (o *BitmapFont) CreateFromFnt(path gdnative.String) gdnative.Int {
	log.Println("Calling BitmapFont.CreateFromFnt()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "create_from_fnt", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the size of a character, optionally taking kerning into account if the next character is provided.
*/
func (o *BitmapFont) GetCharSize(char gdnative.Int, next gdnative.Int) *Vector2 {
	log.Println("Calling BitmapFont.GetCharSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(char)
	goArguments[1] = reflect.ValueOf(next)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_char_size", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *BitmapFont) GetFallback() *BitmapFont {
	log.Println("Calling BitmapFont.GetFallback()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_fallback", goArguments, "*BitmapFont")

	returnValue := goRet.Interface().(*BitmapFont)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a kerning pair as a difference.
*/
func (o *BitmapFont) GetKerningPair(charA gdnative.Int, charB gdnative.Int) gdnative.Int {
	log.Println("Calling BitmapFont.GetKerningPair()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(charA)
	goArguments[1] = reflect.ValueOf(charB)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_kerning_pair", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the font atlas texture at index [code]idx[/code].
*/
func (o *BitmapFont) GetTexture(idx gdnative.Int) *Texture {
	log.Println("Calling BitmapFont.GetTexture()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_texture", goArguments, "*Texture")

	returnValue := goRet.Interface().(*Texture)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the number of textures in the BitmapFont atlas.
*/
func (o *BitmapFont) GetTextureCount() gdnative.Int {
	log.Println("Calling BitmapFont.GetTextureCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_texture_count", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *BitmapFont) SetAscent(px gdnative.Float) {
	log.Println("Calling BitmapFont.SetAscent()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(px)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_ascent", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *BitmapFont) SetDistanceFieldHint(enable gdnative.Bool) {
	log.Println("Calling BitmapFont.SetDistanceFieldHint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_distance_field_hint", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *BitmapFont) SetFallback(fallback *BitmapFont) {
	log.Println("Calling BitmapFont.SetFallback()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(fallback)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_fallback", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *BitmapFont) SetHeight(px gdnative.Float) {
	log.Println("Calling BitmapFont.SetHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(px)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_height", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   BitmapFontImplementer is an interface for BitmapFont objects.
*/
type BitmapFontImplementer interface {
	Class
}