package factory_method

import "testing"

type testCase struct {
	in  string
	out string
	err error
}

func TestIDCardFactory(t *testing.T) {
	cases := []testCase{
		{"owner1", "owner1", nil},
		{"owner2", "owner2", nil},
		{"A B C D", "A B C D", nil},
	}

	factory := &IDCardFactory{}
	// test
	for i := range cases {
		test := cases[i]
		product := factory.Create(factory, test.in)
		if product.Use() != test.out {
			t.Errorf("IDCardFactory should create %v as Product.Use(), from %v. But, create %v.",
				test.out, test.in, product.Use())
		}
		if len(factory.owners) != i+1 {
			t.Error("Invalid number of register items. Some items may be register wrongly.")
		}
	}
}
