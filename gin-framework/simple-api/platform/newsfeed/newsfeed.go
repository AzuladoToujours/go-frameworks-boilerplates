package newsfeed

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

// Item structure
type Item struct {
	Title string `json: "title"`
	Post  string `json: "post"`
}

// Repository
type Repo struct {
	Items []Item
}

// Create a new repo of Items
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// Add an item to items structure
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// Get all the items
func (r *Repo) GetAll() []Item {
	return r.Items
}
