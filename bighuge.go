package thesaurus

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type BigHuge struct {
	APIKEy string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get("http://words.bighugelabs.com/api/2/" +
		b.APIKEy +"/" + term + "/json")

	if err != nil {
		return syns, fmt.Errorf("bighuge: Failed searching synonyms %q: %v", term, err)
	}

	var data synonyms
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)

	return syns, nil
}