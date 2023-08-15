package mob

import (
	"encoding/json"
	"testing"

	"go.olapie.com/utils"
)

func TestMarshalStringList(t *testing.T) {
	l := new(StringList)
	l.Add("a")
	l.Add("b")
	data, err := json.Marshal(l)
	utils.MustNotErrorT(t, err)
	t.Log(string(data))

	var l2 StringList
	err = json.Unmarshal(data, &l2)
	utils.MustNotErrorT(t, err)
	t.Log(l2.Len())
}
