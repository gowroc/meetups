package sql

import (
	"testing"

	"fmt"

	"github.com/google/uuid"
	"github.com/gowroc/meetups/sql/user"
)

func BenchmarkByID(t *testing.B) {
	mustSetup(t)
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s row by ID", tc.name), func(t *testing.B) {
			benchmarkSelectIDManyRows(t, 1000, tc.service)
		})
	}
}

func benchmarkSelectIDManyRows(b *testing.B, rows int, s user.Service) {
	b.StopTimer()
	mustDeleteAllUsers(b, s)
	ids := []uuid.UUID{}
	for i := 0; i < rows; i++ {
		u := user.RandomUser()
		if err := s.InsertUser(u); err != nil {
			b.Fatal("Failed to insert user", err)
		}
		ids = append(ids, u.ID)
	}
	for _, id := range ids {
		if _, err := s.GetUser(id); err != nil {
			b.Fatal("Failed to get user", err)
		}
	}
	b.StartTimer()
	j := 0
	for i := 0; i < b.N; i++ {
		id := ids[j]
		j += 1
		if j >= len(ids) {
			j = 0
		}
		if _, err := s.GetUser(id); err != nil {
			b.Fatal("Failed to get user", err)
		}
	}

}
