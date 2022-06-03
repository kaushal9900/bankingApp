package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Hello", City: "Bhuj", ZipCode: "370001", DateofBirth: "10-10-2021", Status: "1"},
		{Id: "1002", Name: "Hello1", City: "Bhuj1", ZipCode: "370002", DateofBirth: "10-10-2021", Status: "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
