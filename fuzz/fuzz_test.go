package fuzz

import "testing"

func TestNew(t *testing.T) {
	d, err := New(100000, 100000)
	if err != nil {
		t.Log(err)
	}
	for k := range d.UserList {
		t.Logf("%#v", d.UserList[k])
		break
	}

	for k := range d.TaxiList {
		t.Logf("%#v", d.TaxiList[k])
		break
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, _ := New(100000, 100000)
		_ = data
	}
}
