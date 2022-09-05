package data

import (
	"time"

	"github.com/qbitty/greenlight/internal/validator"
)

// Note: You can also prevent a struct field from appearing in the JSON output by simply
// making it unexported. But using the json:"-" struct tag is generally a better choice:
// it’s an explicit indication to both Go and any future readers of your code that you don’t
// want the field included in the JSON, and it helps prevents problems if someone
// changes the field to be exported in the future without realizing the consequences.
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Use the - directive
	Title     string    `json:"title"`
	// 	Hint: If you want to use omitempty and not change the key name then you can leave it
	// blank in the struct tag — like this: json:",omitempty". Notice that the leading comma
	// is still required.
	Year int32 `json:"year,omitempty"` // Add the omitempty directive
	// 	Note that the string directive will only work on struct fields which have int*, uint*,
	// float* or bool types. For any other type of struct field it will have no effect.
	Runtime Runtime  `json:"runtime,omitempty,string"` // Add the string directive
	Genres  []string `json:"genres,omitempty"`         // Add the omitempty directive
	Version int32    `json:"version"`
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
