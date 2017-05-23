package sql

import (
	"sync"
	"testing"

	"github.com/gowroc/meetups/sql/user"
	"github.com/gowroc/meetups/sql/vmodl"
	"github.com/gowroc/meetups/sql/vpgx"
	"github.com/gowroc/meetups/sql/vsql"
	"github.com/gowroc/meetups/sql/vsqlprep"
	"github.com/gowroc/meetups/sql/vsqlx"
)

var setupOnce = sync.Once{}

var testCases = []struct {
	name    string
	service user.Service
}{}

func mustSetup(t testing.TB) {
	setupOnce.Do(
		func() {
			add := func(s user.Service, err error) func(string) {
				if err != nil {
					t.Fatal(err)
				}
				return func(name string) {
					testCases = append(testCases, struct {
						name    string
						service user.Service
					}{name: name, service: s})
				}

			}
			add(vsql.NewUsers())("sql")
			add(vsqlprep.NewUsers())("sqlprep")
			add(vpgx.NewUsers())("pgx")
			add(vsqlx.NewUsers())("sqlx")
			add(vmodl.NewUsers())("modl")
		})
}
