package types

type InputSkinCreate struct {
	Description string `form:"description"`
	Public      bool   `form:"public"`
}
