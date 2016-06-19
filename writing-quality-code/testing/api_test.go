package main
// START OMIT
type mocks struct {
	projects []Project
}

type fakeDatabase struct {
	Database
	m mocks
}

func newFakeDB(m mocks) *fakeDatabase {
	db := new(fakeDatabase)
	db.m = m
	return db
}

func (d *fakeDatabase) AllProjects() ([]Project, error) {
	return d.m.projects, nil
}

func TestProjects(t *testing.T) {
    db := newFakeDB(mocks{projects: []Project{...}})}
    // Do some testing
}
// END OMIT