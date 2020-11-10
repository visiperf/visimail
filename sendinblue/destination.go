package sendinblue

type destination struct {
	Email string `json:"email"`
}

func newDestinations(emails []string) []*destination {
	var destinations []*destination
	for _, email := range emails {
		destinations = append(destinations, &destination{Email: email})
	}

	return destinations
}
