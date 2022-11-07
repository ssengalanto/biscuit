package account

type Email string

// IsValid checks the validity of the email address.
func (e Email) IsValid() (bool, error) {
	// TODO: add validation logic
	return true, nil
}

// String converts type Email to type string.
func (e Email) String() string {
	return string(e)
}

type Password string

// IsValid checks the validity of the password.
func (p Password) IsValid() (bool, error) {
	// TODO: add validation logic
	return true, nil
}

// String converts type Password to type string.
func (p Password) String() string {
	return string(p)
}
