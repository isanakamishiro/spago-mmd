package store

// RendererInfo is ...
type RendererInfo struct {
	MemoryGeometries string
	MemoryTextures   string

	RenderCalls     string
	RenderTriangles string
	RenderPoints    string
	RenderLines     string
	RenderFrame     string
}

var rendererInfo *RendererInfo = &RendererInfo{}

// RendererInfoStore gets ...
func RendererInfoStore() *RendererInfo {
	return rendererInfo
}
