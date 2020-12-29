package api

type PageParams struct {
	Index int32 `form:"index"`
	Size  int32 `form:"size"`
}
