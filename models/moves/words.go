package moves

import (
	"fmt"
	"github.com/wongyinlong/walletBackend/models"
	"gopkg.in/mgo.v2/bson"
)

type Words struct {
	WordsList []string `json:"words_list"`
}

func (w *Words) GetHelpWords(prefix string) {
	var words Words
	q := bson.M{"title": bson.M{"$regex": bson.RegEx{Pattern: "/a/", Options: "i"}}}
	db := "TEST"
	c := "WORDS"
	err := models.FindAll(db, c, q, &words)
	fmt.Println(err)
}
