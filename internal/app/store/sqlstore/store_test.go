package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// TestMain ...
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=mms_dev password=denis895 sslmode=disable"
	}

	os.Exit(m.Run())
}
