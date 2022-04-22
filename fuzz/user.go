package fuzz

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"

	"github.com/dgkg/api/model"
	"github.com/google/uuid"
)

// fuzzUser_old not used anymore.
func fuzzUser_old(n int) ([]*model.User, error) {
	if n > 255000 {
		return nil, errors.New("fuzz: out of boundaries max 255 000 users")
	}

	content, err := ioutil.ReadFile("testdata/names.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(content))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	us := make([]*model.User, n)
	for i := 0; i < n; i++ {
		us[i] = &model.User{
			ID:        uuid.New().String(),
			FirstName: records[i][1],
			LastName:  strings.ToUpper(records[rand.Intn(len(records))][1]),
			Email:     strings.ToLower(records[i][1] + listDomainName[rand.Intn(len(listDomainName))]),
		}
	}
	return us, nil
}

func FuzzUser(n int) (map[string]*model.User, error) {
	if n > 255000 {
		return nil, errors.New("fuzz: out of boundaries max 255 000 users")
	}

	content, err := ioutil.ReadFile("testdata/names.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(content))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	us := make(map[string]*model.User, n)
	for i := 0; i < n; i++ {
		id := uuid.New().String()
		us[id] = &model.User{
			ID:        id,
			FirstName: records[i][1],
			LastName:  strings.ToUpper(records[rand.Intn(len(records))][1]),
			Email:     records[i][1] + listDomainName[rand.Intn(len(listDomainName))],
		}
	}
	return us, nil
}

var listDomainName = []string{
	"@google.com",
	"@gmail.com",
	"@yopmail.com",
	"@altavista.com",
}
