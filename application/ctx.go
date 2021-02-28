package application

import (
	"database/sql"
	"errors"

	// The Application handles our DB connection
	_ "github.com/lib/pq"
)

// Ctx is the Application Context
type Ctx struct {
	db *sql.DB
}

// Application holds the Application behavior
type Application interface {
	GetDBConn() *sql.DB
	TearDown()
}

// Run the application
func Run(run func(app Application)) error {
	app, err := new()
	if err != nil {
		return err
	}

	run(app)

	app.TearDown()

	return nil
}

// New up an instance of the Application
func new() (Application, error) {
	dbConn, err := open()
	if err != nil {
		return nil, errors.New("Could not connect to database: " + err.Error())
	}

	return &Ctx{
		db: dbConn,
	}, nil
}

// GetDBConn retrieves the Application's DB connection
func (a *Ctx) GetDBConn() *sql.DB {
	return a.db
}

// TearDown the Application gracefully - always do this!
func (a *Ctx) TearDown() {
	a.db.Close()
}
