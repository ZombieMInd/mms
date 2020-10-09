package sqlstore_test

import (
	"testing"

	"github.com/ZombieMInd/mms/internal/app/model"
	"github.com/ZombieMInd/mms/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestingDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	u := model.TestingUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)

}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestingDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	email := "user@example.com"

	assert.Error(t, s.User().FindByEmail(email))

	u := model.TestingUser(t)
	u.Email = email
	s.User().Create(u)

	assert.NoError(t, s.User().FindByEmail(email))
	assert.NotNil(t, u)
}
