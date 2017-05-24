package sql

import (
	"testing"

	"github.com/gowroc/meetups/sql/user"
	"github.com/gowroc/meetups/sql/vmodl"
)

func TestPosts(t *testing.T) {
	u, err := vmodl.NewUsers()
	if err != nil {
		t.Fatal(err)
	}
	testWithPost(t, u)
}

func testWithPost(t *testing.T, s *vmodl.Users) {
	const usersToInsert = 10
	const postsToInsert = 10

	mustDeleteAllUsers(t, s)
	for i := 0; i < usersToInsert; i++ {
		u := user.RandomUser()
		if err := s.InsertUser(u); err != nil {
			t.Fatal(err)
		}
		for i := 0; i < postsToInsert; i++ {
			p := user.RandomPost(u.ID)
			if err := s.InsertPost(p); err != nil {
				t.Fatal(err)
			}
		}
		uwp, err := s.GetUserWithPosts(u.ID)
		if err != nil {
			t.Fatal(err)
		}
		if len(uwp.UserPosts) != postsToInsert {
			t.Fatalf("Expected %d, got %d", len(uwp.UserPosts), postsToInsert)
		}
	}
	us, err := s.GetAllUsers()
	if err != nil {
		t.Fatal(err)
	}
	if len(us) != usersToInsert {
		t.Fatalf("Expected %d, got %d", len(us), usersToInsert)
	}
	//mustDeleteAllUsers(t, s)
}
