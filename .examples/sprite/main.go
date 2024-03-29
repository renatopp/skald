package main

import (
	"math/rand"

	sk "github.com/renatopp/skald"
	"github.com/renatopp/skald/core"
)

type DRotater struct{}

var Rotater = sk.NewComponent[DRotater]()

var RotatersQuery = sk.NewQuery[struct {
	Transform *core.DTransform
	Rotater   *DRotater
}]()

var RotateRotators = sk.NewSystem(func(g *sk.Game) error {
	for _, r := range RotatersQuery.Query() {
		r.Transform.RotateBy(180 * g.Timer.DeltaTime)
	}
	return nil
}, RotatersQuery)

func main() {
	game := sk.NewGame()
	handle := game.Assets.Load(".examples/_assets/rabbitv3.png")

	rabbits := sk.NewEntityWithOptions(sk.EntityOptions{
		Components: sk.Components{core.Transform, core.Sprite, Rotater},
		OnSpawned: func(e *sk.EntityInstance) {
			s := core.Sprite.Get(e)
			s.Texture = handle.AsTexture()
		},
	})

	game.AddSystem(core.SpriteRenderer)
	game.AddSystem(RotateRotators)
	game.World.SpawnMulti(15000, rabbits, func(e *sk.EntityInstance) {
		t := core.Transform.Get(e)
		t.MoveTo(rand.Float64()*2-1, rand.Float64()*2-1)
	})
	// game.World.Spawn(rabbits)
	game.Play()
}
