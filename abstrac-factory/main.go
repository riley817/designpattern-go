package main

type ShoeInterface interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type ShirtInterface interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type SportsFactoryInterface interface {
	makeShoe() ShoeInterface
	makeShirt() ShirtInterface
}

func GetSportsFactory(brand string) (SportsFactoryInterface, error) {
	if brand == "adidas" {
	}
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}

type Adidas struct {
}

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

func (a *Adidas) makeShoe() ShoeInterface {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) makeShirt() ShirtInterface {
	return &AdidasShirt{Shirt{
		logo: "adidas",
		size: 14,
	}}
}

type Nike struct {
}

type NikeShoe struct {
	Nike
}

type NikeShirt struct {
	Nike
}

func main() {

}
