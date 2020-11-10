package sendinblue

type destination struct {
	Email string `json:"email"`
}

func newDestination(email string) *destination {
	if email == "" {
		return nil
	}

	return &destination{Email: email}
}

func newDestinations(emails []string) []*destination {
	var destinations []*destination
	for _, email := range emails {
		destinations = append(destinations, newDestination(email))
	}

	return destinations
}
