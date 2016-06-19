package main
// START OMIT
type Database interface {
    Projects() ([]Project, error)
}

func NewPostgres(db *sql.DB) Database {
	return &postgres{db: db}
}

type postgres struct {
	db *sql.DB
}

func (d *postgres) AllProjects() ([]Project, error) {
	// Query DB, return some projects.
	return projs, nil
}

func main() {
    dbConn, _ := sql.Open("postgres", config.PSQLConn)
    db := NewPostgres(dbConn)
    projs, _ := db.Projects()
    fmt.Println(projs)
}
// END OMIT