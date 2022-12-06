package pgsql

import "fmt"

// List of valid keys for accountQueries map.
const (
	QueryExists        = "exists"
	QueryCreateAccount = "createAccount"
	QueryCreatePerson  = "createPerson"
	QueryFindByID      = "findById"
	QueryFindByEmail   = "findByEmail"
	QueryUpdateByID    = "updateByID"
	QueryDeleteByID    = "deleteByID"
)

// AccountQueries is a map holds all queries for account table.
var accountQueries = map[string]string{ //nolint:gochecknoglobals //intended
	QueryExists:        accountExistsQuery,
	QueryCreateAccount: createAccountQuery,
	QueryFindByID:      findByIDQuery,
	QueryFindByEmail:   findByEmailQuery,
	QueryUpdateByID:    updateByIDQuery,
	QueryDeleteByID:    deleteByIDQuery,
}

const accountExistsQuery = `
	SELECT COUNT(1)
	FROM account
	WHERE id = $1;
	`

const createAccountQuery = `
	INSERT INTO account (id, email, password, active, last_login_at)
	VALUES ($1, $2, $3, $4, $5);
	`

const findByIDQuery = `
	SELECT id, email, password, active, last_login_at
	FROM account
	WHERE id = $1;
	`

const findByEmailQuery = `
	SELECT id, email, password, active, last_login_at
	FROM account
	WHERE email = $1;
	`

const updateByIDQuery = `
	UPDATE account
	SET email = $2, password = $3, active = $4, last_login_at = $5, updated_at = NOW()
	FROM account
	WHERE id = $1
	RETURNING *;
	`

const deleteByIDQuery = `
	DELETE FROM account
	WHERE id = $1
	RETURNING *;
	`

// MustBeValidAccountQuery accepts a string parameter that will be used
// as a key to accountQueries map. If the key doesn't exist it will
// throw a panic, otherwise it will return the query.
func MustBeValidAccountQuery(s string) string {
	query, ok := accountQueries[s]
	if !ok {
		panic(fmt.Errorf("%w: `%s` doesn't exists in account queries", ErrInvalidQuery, s))
	}

	return query
}
