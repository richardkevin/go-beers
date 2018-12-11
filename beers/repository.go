package beers

import (
	"errors"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const BeerCollection = "beer"

var ErrDuplicatedBeer = errors.New("Duplicated beer")

type Beer struct {
	Id      string `bson:"_id"`
	Name    string `bson:"name"`
	Sold_out bool   `bson:sold_out`
}

type BeerRepository struct {
	session *mgo.Session
}

func (r *BeerRepository) Create(p *Beer) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(BeerCollection)
	err := collection.Insert(p)
	mongoErr, ok := err.(*mgo.LastError)
	if ok && mongoErr.Code == 11000 {
		return ErrDuplicatedBeer
	}
	return err
}

func (r *BeerRepository) Update(p *Beer) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(BeerCollection)
	return collection.Update(bson.M{"_id": p.Id}, p)
}

func (r *BeerRepository) Remove(id string) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(BeerCollection)
	return collection.Remove(bson.M{"_id": id})
}

func (r *BeerRepository) FindAllActive() ([]*Beer, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(BeerCollection)
	query := bson.M{"sold_out": false}

	documents := make([]*Beer, 0)

	err := collection.Find(query).All(&documents)
	return documents, err
}

func (r *BeerRepository) FindById(id string) (*Beer, error) {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(BeerCollection)
	query := bson.M{"_id": id}

	beer := &Beer{}

	err := collection.Find(query).One(beer)
	return beer, err
}

func NewBeerRepository(session *mgo.Session) *BeerRepository {
	return &BeerRepository{session}
}

// func main() {
// 	session, err := mgo.Dial("localhost:27017/go-beers")
//
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	repository := NewBeerRepository(session)
//
// 	// creating a beer
// 	beer := &Beer{Id: "1", Name: "Heineken"}
// 	err = repository.Create(beer)
//
// 	if err == ErrDuplicatedBeer {
// 		log.Printf("%s is already created\n", beer.Name)
// 	} else if err != nil {
// 		log.Println("Failed to create a beer: ", err)
// 	}
//
// 	// updating a beer
// 	beer.Name = "Heikenen updated"
// 	err = repository.Update(beer)
//
// 	if err != nil {
// 		log.Println("Failed to update a beer: ", err)
// 	}
//
// 	repository.Create(&Beer{Id: "2", Name: "Heineken"})
// 	repository.Create(&Beer{Id: "3", Name: "Stela"})
// 	repository.Create(&Beer{Id: "4", Name: "Amstel"})
// 	repository.Create(&Beer{Id: "5", Name: "Brahma"})
//
// 	// remove
// 	err = repository.Remove("4")
// 	if err != nil {
// 		log.Println("Failed to remove a beer: ", err)
// 	}
//
// 	// findAll
// 	brand, err := repository.FindAllActive()
// 	if err != nil {
// 		log.Println("Failed to fetch brand: ", err)
// 	}
//
// 	for _, beer := range brand {
// 		log.Printf("Have in database: %#v\n", beer)
// 	}
//
// 	// FindById
// 	beer, err = repository.FindById("1")
// 	if err == nil {
// 		log.Printf("Result of findById: %v\n", beer)
// 	} else {
// 		log.Println("Failed to findById ", err)
// 	}
// }
