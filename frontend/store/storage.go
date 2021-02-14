package store

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

func GetRendererInfo() *RendererInfo {
	return rendererInfo
}
