package session

import "testing"

func TestMiddleware(t *testing.T) {
	jwtValue, _ := New("076dba72-c24e-11ec-b9fb-ab1edf931d31", 2)
	t.Log(jwtValue)
}
