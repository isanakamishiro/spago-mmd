package views

import (
	"app/frontend/components"
	"github.com/nobonobo/spago"
)

// Render ...
func (c *Top) Render() spago.HTML {
	return spago.Tag("body", 
		spago.Tag("section", 			
			spago.A("class", spago.S(`hero is-info is-fullheight`)),
			spago.Tag("div", 				
				spago.A("class", spago.S(`hero-head`)),
				spago.C(&components.Header{}),
			),
			spago.Tag("div", 				
				spago.A("class", spago.S(`hero-body p-0`)),
				spago.Tag("canvas", 					
					spago.A("id", spago.S(`cv`)),
					spago.A("width", spago.S(``, spago.S(c.canvasWidth), ``)),
					spago.A("height", spago.S(``, spago.S(c.canvasHeight), ``)),
					spago.A("style", spago.S(`width: 100%;`)),
				),
			),
			spago.Tag("div", 				
				spago.A("class", spago.S(`hero-foot`)),
				spago.Tag("nav", 					
					spago.A("class", spago.S(`tabs is-boxed is-fullwidth`)),
					spago.Tag("div", 						
						spago.A("class", spago.S(`container`)),
						spago.Tag("ul", 
							spago.Tag("li", 
								spago.Tag("a", 									
									spago.Event("click", c.refresh),
									spago.T(`DOM Refresh`),
								),
							),
							spago.Tag("li", 
								spago.Tag("a", 									
									spago.Event("click", c.resetCameraPosition),
									spago.T(`Reset Camera Pos`),
								),
							),
						),
					),
				),
			),
		),
	)
}
