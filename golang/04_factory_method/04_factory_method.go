package factory_method

type producer interface {
	createProduct(owner string) Product
	registerProduct(Product)
}

type Factory struct{}

func (self *Factory) Create(prod producer, owner string) Product {
	p := prod.createProduct(owner)
	prod.registerProduct(p)
	return p
}

type Product interface {
	Use() string
}

type IDCard struct {
	owner string
}

func (self *IDCard) Use() string {
	return self.owner
}

type IDCardFactory struct {
	*Factory
	owners []*string
}

func (self *IDCardFactory) createProduct(owner string) Product {
	return &IDCard{owner}
}

func (self *IDCardFactory) registerProduct(prod Product) {
	owner := prod.Use()
	self.owners = append(self.owners, &owner)
}
