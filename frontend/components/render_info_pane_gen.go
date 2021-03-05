package components

import (
	"app/frontend/store"
	"github.com/nobonobo/spago"
)

// Render ...
func (c *RendererInfoPane) Render() spago.HTML {
	return spago.Tag("div", 		
		spago.A("class", spago.S(`content`)),
		spago.Tag("ul", 
			spago.Tag("li", 
				spago.T(`MemoryGeometries : `, spago.S(store.RendererInfoStore().MemoryGeometries), ``),
			),
			spago.Tag("li", 
				spago.T(`MemoryTextures : `, spago.S(store.RendererInfoStore().MemoryTextures), ``),
			),
		),
		spago.Tag("ul", 
			spago.Tag("li", 
				spago.T(`RenderCalls : `, spago.S(store.RendererInfoStore().RenderCalls), ``),
			),
			spago.Tag("li", 
				spago.T(`RenderTriangles : `, spago.S(store.RendererInfoStore().RenderTriangles), ``),
			),
			spago.Tag("li", 
				spago.T(`RenderPoints : `, spago.S(store.RendererInfoStore().RenderPoints), ``),
			),
			spago.Tag("li", 
				spago.T(`RenderLines : `, spago.S(store.RendererInfoStore().RenderLines), ``),
			),
			spago.Tag("li", 
				spago.T(`RenderFrame : `, spago.S(store.RendererInfoStore().RenderFrame), ``),
			),
		),
	)
}
