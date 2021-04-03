package text

// AnimationMixer is animation mixer for TextMesh.
type AnimationMixer struct {
	textMesh *Mesh

	animateSpeed float64
	animateFrame float64

	animations []Animation
}

// NewTextAnimationMixer creates TextAnimationMixer.
func NewTextAnimationMixer(tm *Mesh) *AnimationMixer {
	return &AnimationMixer{
		textMesh:     tm,
		animateSpeed: 10,
		animateFrame: 0,
		animations:   []Animation{FadeIn()},
	}
}

// SetAnimation sets animation functions.
func (c *AnimationMixer) SetAnimation(animatinos ...Animation) {
	c.animations = animatinos
}

// SetAnimationSpeed sets animatino speed.
func (c *AnimationMixer) SetAnimationSpeed(speed float64) {
	c.animateSpeed = speed
}

// Restart is ...
func (c *AnimationMixer) Restart() {
	c.animateFrame = 0
}

// Update is ...
func (c *AnimationMixer) Update(delta float64) {

	maxFrame := float64(len(c.textMesh.renderText))
	c.animateFrame += (delta * c.animateSpeed)
	if c.animateFrame > maxFrame {
		c.animateFrame = maxFrame
	}

	// Spriteにアニメーションをかける
	for n, sp := range c.textMesh.characterMeshList {

		offset := float64(n)
		frame := c.animateFrame - offset
		if frame < 0.0 {
			frame = 0.0
		} else if frame > 1.0 {
			frame = 1.0
		}

		visible := (frame > 0)
		sp.SetVisible(visible)

		for _, animate := range c.animations {
			animate(sp, frame)
		}
	}
}
