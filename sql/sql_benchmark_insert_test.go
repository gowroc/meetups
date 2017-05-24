package sql

import (
	"testing"

	"fmt"

	"github.com/gowroc/meetups/sql/user"
)

func BenchmarkInsert(t *testing.B) {
	mustSetup(t)
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s insert", tc.name), func(t *testing.B) {
			benchmarkInsert(t, tc.service)
		})
	}
}

func benchmarkInsert(b *testing.B, s user.Service) {
	b.StopTimer()
	mustDeleteAllUsers(b, s)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if err := s.InsertUser(user.RandomUser()); err != nil {
			b.Fatal("Failed to insert user", err)
		}
	}
}
