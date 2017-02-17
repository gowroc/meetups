package main

import (
	"database/sql"
)

func main() {
    db := sql.Open()
    defer db.Close()

    if something {
        return
    }

    some
    more
    code

    if something_else {
        return
    }

    even
    more
    code
    
    return
}
