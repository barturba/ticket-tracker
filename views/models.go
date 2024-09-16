package views

var Id = 0

type Contact struct {
	Name  string
	Email string
	Id    int
}

func NewContact(name, email string) Contact {
	Id++
	return Contact{
		Name:  name,
		Email: email,
		Id:    Id,
	}
}

type Contacts []Contact

type Data struct {
	Contacts Contacts
}

func NewData() Data {
	return Data{[]Contact{
		NewContact("Alice", "aoeu"),
		NewContact("Bob", "bob@gmail.com"),
		NewContact("Charlie", "charlie@gmail.com"),
	}}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Form FormData
	Data Data
}

func NewPage() Page {
	return Page{
		Data: NewData(),
		Form: NewFormData(),
	}
}
