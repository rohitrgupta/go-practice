package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Vval struct {
	Name  string
	value float64
}

type Var struct {
	v []Vval
}

func (a *Var) UnmarshalJSON(b []byte) error {
	fmt.Println("unmarshaling", b)
	var s interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	fmt.Println("um data", s)

	return nil
}

func (a Var) MarshalJSON() ([]byte, error) {
	var s string
	for k, v := range a.v {
		s += fmt.Sprintf("\"%s\": %f", v.Name, v.value)
		if k < len(a.v)-1 {
			s += ","
		}
	}
	rep := []byte("{" + s + "}")
	return rep, nil
}

func main() {
	var v = Var{
		v: []Vval{
			Vval{Name: "a", value: 1.0},
			Vval{Name: "b", value: 2.0},
		},
	}
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

}
