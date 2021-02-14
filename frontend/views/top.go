package views

import (
	"app/frontend/actions"
	"app/lib/threejs"
	"app/lib/threejs/cameras"
	"app/lib/threejs/controls"
	"app/lib/threejs/effects"
	"app/lib/threejs/lights"
	"app/lib/threejs/loaders"
	"app/lib/threejs/loaders/mmdloaders"
	"app/lib/threejs/objects/sky"
	"app/lib/threejs/objects/water"
	"math"
	"syscall/js"

	"github.com/nobonobo/spago"
	"github.com/nobonobo/spago/dispatcher"
)

//go:generate spago generate -c Top -p views top.html

// Top  ...
type Top struct {
	spago.Core

	canvasWidth  int
	canvasHeight int

	init bool

	renderer threejs.Renderer
	camera   threejs.Camera
	scene    threejs.Scene
	control  controls.OrbitControls
	clock    threejs.Clock

	effector  *effects.OutlineEffect
	character *mmdloaders.MMDAnimationHelper

	renderFunction js.Func
}

// NewTop is ...
func NewTop() *Top {

	top := &Top{
		init:         false,
		canvasWidth:  0,
		canvasHeight: 0,
	}

	return top
}

// Mount is ...
func (c *Top) Mount() {
	if !c.init {
		c.initSceneAndRenderer()

		c.init = true
		c.renderFunction = js.FuncOf(c.render)

		js.Global().Call("addEventListener", "resize", js.FuncOf(c.onResize))

	}

	c.renderer.SetAnimationLoop(c.renderFunction)

}

// Unmount ...
func (c *Top) Unmount() {

	c.renderer.CancelAnimationLoop()

}

func (c *Top) initSceneAndRenderer() {
	// Renderer
	{
		canvas := js.Global().Get("document").Call("querySelector", "#cv")
		renderer := threejs.NewWebGLRenderer(map[string]interface{}{
			"canvas":    canvas,
			"antialias": true,
			// "aplha":     true,
		})
		renderer.SetPhysicallyCorrectLights(true)

		c.renderer = renderer
	}

	// Camera
	{
		const (
			fov    = 40
			aspect = 4 / 3
			near   = 0.1
			far    = 2000
		)
		camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
		camera.Position().SetY(10)
		camera.Position().SetZ(40)
		camera.Up().SetZ(1)
		camera.LookAtXYZ(0, 0, 0)
		c.camera = camera
	}

	// Scene
	{
		scene := threejs.NewScene()
		scene.SetBackgroundColor(threejs.NewColorFromColorValue(0xeeeeee))
		c.scene = scene
	}

	// Clock
	{
		c.clock = threejs.NewClock(true)
	}

	// Control
	{
		control := controls.NewOrbitControls(c.camera, c.renderer.DomElement())
		control.Target().Set2(0, 10, 0)
		control.Update()
		control.SaveState()

		c.control = control
	}

	// Sky Dome
	{
		sky := sky.NewSky()
		sky.Scale().SetScalar(1000)

		sky.SetTurbidity(12.7)
		sky.SetRayleigh(0.1)
		sky.SetMieCoefficient(0)
		sky.SetMieDirectionalG(0.3)
		sky.SetInclination(0)
		sky.SetAzimuth(0.2)

		c.scene.Add(sky)
	}

	// Ocean
	{
		// Normal Texture
		loader := loaders.NewTextureLoader()
		tx := loader.LoadSimply("./assets/threejs/ex/textures/water/Water_1_M_Normal.jpg")
		tx.SetWrapS(threejs.RepeatWrapping)
		tx.SetWrapT(threejs.RepeatWrapping)

		ocean := water.NewOcean(
			1000,
			1000,
			water.TextureSize(512, 512),
			water.NormalizeTexture(tx),
			water.Alpha(1.0),
			water.DistortionScale(3.7),
			water.SunColor(threejs.NewColorFromColorValue(0xffffff)),
			water.OceanColor(threejs.NewColorFromColorValue(0x001e0f)),
			water.Fog(false),
		)
		ocean.SetSize(1.0)

		ocean.Rotation().SetX(-math.Pi / 2)
		ocean.Position().SetY(1)

		c.scene.Add(ocean)
	}

	// Water Ball
	{
		loader := loaders.NewCubeTextureLoader()
		loader.SetPath("./assets/threejs/ex/textures/cube/Park3Med/")
		urls := []string{
			"px.jpg", "nx.jpg",
			"py.jpg", "ny.jpg",
			"pz.jpg", "nz.jpg",
		}
		tx := loader.LoadSimply(urls)
		ball := water.NewBall(tx)

		// ball.Scale().Set2(0.5, 0.5, 0.5)
		ball.Position().Set2(20, 15, -20)

		c.scene.Add(ball)
	}

	// Character
	{
		c.effector = effects.NewOutlineEffect(c.renderer)

		mmdHelper := mmdloaders.NewMMDAnimationHelper(map[string]interface{}{
			"afterglow": 2.0,
		})
		c.character = mmdHelper

		var fn js.Func
		fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			defer fn.Release()

			// model
			// modelFile := "./assets/models/mmd/miku/miku_v2.pmd"
			modelFile := "./assets/models/mmd/lisa/lisa.pmx"
			vmdFiles := []string{"./assets/models/mmd/vmds/wavefile_v2.vmd"}

			mmdLoader := mmdloaders.NewMMDLoader()
			mmdLoader.LoadWithAnimation(modelFile, vmdFiles, func(mesh mmdloaders.MMDMesh, animation mmdloaders.MMDAnimation) {

				mesh.SetCastShadow(true)
				// mesh.SetReceiveShadow(true)

				c.scene.AddMesh(mesh)
				mmdHelper.Add(mesh, map[string]interface{}{
					"animation": animation.JSValue(),
					"physics":   true,
				})

			}, nil, nil)

			return nil

		})

		// Ammo functionが呼ばれていない状態 = Function, コール後、Object
		if js.Global().Get("Ammo").Type() == js.TypeFunction {
			js.Global().Call("Ammo").Call("then", fn)
		} else {
			fn.Invoke()
		}

	}

	// Light

	// HemisphereLight
	{
		const (
			skyColor       = threejs.ColorValue(0xffffff)
			groundColor    = threejs.ColorValue(0xb97a20)
			lightIntensity = threejs.LightIntensity(2)
		)
		light := lights.NewHemisphereLight(skyColor, groundColor, lightIntensity)

		c.scene.AddLight(light)
	}

	// DirectionalLight
	{
		const (
			lightColor     = threejs.ColorValue(0xffffff)
			lightIntensity = threejs.LightIntensity(2)
			width          = 12.0
			height         = 4.0
		)
		light := lights.NewDirectionalLight(lightColor, lightIntensity)
		light.SetCastShadow(true)
		light.Position().Set2(-15, 40, 15)
		light.Target().Position().Set2(-4, 0, -4)

		light.Shadow().Camera().SetLeft(-20)
		light.Shadow().Camera().SetRight(20)
		light.Shadow().Camera().SetTop(20)
		light.Shadow().Camera().SetBottom(-20)

		c.scene.AddLight(light)
		c.scene.Add(light.Target())

		// cameraHelper := cameras.NewCameraHelper(light.Shadow().Camera())
		// scene.Add(cameraHelper)

		helper := lights.NewDirectionalLightHelper(light)
		c.scene.Add(helper)
	}

}

// resizeRendererToDisplaySize resizes render display size.
func (c *Top) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
	canvas := renderer.DomElement()
	height := canvas.Get("height").Int()

	// キャンバスのheightが0だった場合、キャンバスサイズを再設定する
	if height == 0 {
		// 親ノードのサイズをキャンバスサイズとして参照する
		parent := canvas.Get("parentNode")
		rect := parent.Call("getBoundingClientRect")
		w := rect.Get("right").Float() - rect.Get("left").Float()
		h := rect.Get("bottom").Float() - rect.Get("top").Float()

		// キャンバスサイズ(Not DOMのサイズ)を計算
		pixelRatio := js.Global().Get("devicePixelRatio").Float()
		clientWidth := (w * pixelRatio)
		clientHeight := (h * pixelRatio)

		renderer.SetSize(clientWidth, clientHeight, false)
		c.canvasWidth = int(clientWidth)
		c.canvasHeight = int(clientHeight)

		// log.Printf("w: %v, h: %v, cw: %v, ch: %v, pr: %v", w, h, clientWidth, clientHeight, pixelRatio)

		return true
	}

	return false
}

func (c *Top) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	// Update Matrix
	c.control.Update()

	// Update time and animation
	delta := c.clock.Delta()
	c.character.Update(delta)

	// Render
	// c.effector.Render(c.scene, c.camera)
	c.renderer.Render(c.scene, c.camera)

	return nil
}

//
// Event Handling
//

func (c *Top) onResize(this js.Value, args []js.Value) interface{} {

	// Canvas Size Reset
	c.canvasHeight = 0

	dispatcher.Dispatch(actions.Refresh)

	return nil
}

func (c *Top) refresh(ev js.Value) {

	dispatcher.Dispatch(actions.Refresh)

}

func (c *Top) resetCameraPosition(ev js.Value) {

	c.control.Reset()

}
