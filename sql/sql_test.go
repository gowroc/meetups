package sql

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gowroc/meetups/sql/user"
)

func TestSQL(t *testing.T) {
	mustSetup(t)
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Testing %s", tc.name),
			func(t *testing.T) {
				test(t, tc.service)
			})
	}

}

func test(t *testing.T, s user.Service) {
	const usersToInsert = 10

	mustDeleteAllUsers(t, s)
	for i := 0; i < usersToInsert; i++ {
		u := user.RandomUser()
		if err := s.InsertUser(u); err != nil {
			t.Fatal(err)
		}
		u2, err := s.GetUser(u.ID)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(u, u2) {
			t.Fatalf("Expected %+v, got %+v", u, u2)
		}
	}
	us, err := s.GetAllUsers()
	if err != nil {
		t.Fatal(err)
	}
	if len(us) != usersToInsert {
		t.Fatalf("Expected %d, got %d", len(us), usersToInsert)
	}
	mustDeleteAllUsers(t, s)
}

func mustDeleteAllUsers(t testing.TB, s user.Service) {
	if err := s.DeleteAllUsers(); err != nil {
		t.Fatal(err)
	}
	us, err := s.GetAllUsers()
	if err != nil {
		t.Fatal(err)
	}
	if len(us) != 0 {
		t.Fatal("Failed to delete all users", us)
	}
}
