package models

type Movie struct {
	ID       string    `json:"id"`
	IMDbID   string    `json:"imdbid"`
	Title    string    `json:"title"`
	Year     string    `json:"year"`
	Genre    string    `json:"genre"`
	Director *Director `json:"director"`
}

func (m Movie) Init() []Movie {
	movies := []Movie{}

	movies = append(movies,
		Movie{
			ID:     "1",
			IMDbID: "tt0109830",
			Title:  "Forrest Gump",
			Year:   "1994",
			Genre:  "Drama, Romance",
			Director: &Director{
				FirstName: "Robert",
				LastName:  "Zemeckis"},
		})

	movies = append(movies,
		Movie{
			ID:     "2",
			IMDbID: "tt0111161",
			Title:  "The Shawshank Redemption",
			Year:   "1994",
			Genre:  "Drama",
			Director: &Director{
				FirstName: "Frank",
				LastName:  "Darabont"},
		})

	movies = append(movies,
		Movie{
			ID:     "3",
			IMDbID: "tt1375666",
			Title:  "Inception",
			Year:   "2010",
			Genre:  "Action, Adventure, Sci-Fi",
			Director: &Director{
				FirstName: "Christopher",
				LastName:  "Nolan"},
		})

	movies = append(movies,
		Movie{
			ID:     "4",
			IMDbID: "tt0816692",
			Title:  "Interstellar",
			Year:   "2014",
			Genre:  "Adventure, Drama, Sci-Fi",
			Director: &Director{
				FirstName: "Christopher",
				LastName:  "Nolan"},
		})

	movies = append(movies,
		Movie{
			ID:     "5",
			IMDbID: "tt2543472",
			Title:  "Wonder",
			Year:   "2017",
			Genre:  "Drama, Family",
			Director: &Director{
				FirstName: "Stephen",
				LastName:  "Chbosky"},
		})

	return movies
}
