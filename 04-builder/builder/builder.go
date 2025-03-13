package builder

type Cake struct {
	Dressing string
	flavor   string
}

type Builder interface {
	BuildDressing()
	BuildFlavor()
	GetCake() Cake
}

type BaseCakeBuilder struct {
	cake Cake
}

func (b *BaseCakeBuilder) GetCake() Cake {
	return b.cake
}

type ChocolateCakeBuilder struct {
	BaseCakeBuilder
}

func (b *ChocolateCakeBuilder) BuildDressing() {
	b.cake.Dressing = "butter"
}

func (b *ChocolateCakeBuilder) BuildFlavor() {
	b.cake.flavor = "chocolate"
}

type VanillaCakeBuilder struct {
	BaseCakeBuilder
}

func (b *VanillaCakeBuilder) BuildDressing() {
	b.cake.Dressing = "butter"
}

func (b *VanillaCakeBuilder) BuildFlavor() {
	b.cake.flavor = "vanilla"
}

type CakeMaker struct{}

func (c *CakeMaker) MakeCake(cakeBuilder Builder) Cake {
	cakeBuilder.BuildDressing()
	cakeBuilder.BuildFlavor()

	return cakeBuilder.GetCake()
}
