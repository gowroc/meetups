package sql

import (
	"fmt"
	"testing"

	"github.com/gowroc/meetups/sql/user"
)

func BenchmarkSelectMany(t *testing.B) {
	mustSetup(t)
	for _, tc := range testCases {
		for _, n := range []int{2, 5, 10, 20, 50, 100} {
			t.Run(fmt.Sprintf("%s on selecting %d rows", tc.name, n), func(t *testing.B) {
				benchmarkSelectManyRows(t, n, tc.service)
			})
		}

	}
}

func benchmarkSelectManyRows(b *testing.B, rows int, s user.Service) {
	b.StopTimer()
	mustDeleteAllUsers(b, s)
	for i := 0; i < rows; i++ {
		if err := s.InsertUser(user.RandomUser()); err != nil {
			b.Fatal("Failed to insert user", err)
		}
	}
	for i := 0; i < 10; i++ {
		if _, err := s.GetAllUsers(); err != nil {
			b.Fatal("Failed to get users", err)
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if _, err := s.GetAllUsers(); err != nil {
			b.Fatal("Failed to get users", err)
		}
	}

}
