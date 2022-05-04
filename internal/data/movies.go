package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`                // Unique ID integer for the movie
	CreatedAt time.Time `json:"-"`                 // Timestamp for when the movie is added to our database
	Title     string    `json:"title"`             // Movie title
	Year      int32     `json:"year,omitempty"`    // Movie release year
	Runtime   int32     `json:"runtime,omitempty"` // Movie runtime (in minutes)
	Genres    []string  `json:"genres,omitempty"`  // Slice of genres for the movie
	Version   int32     `json:"version"`           // The version number starts at 1 and will be incremented each time the movie information is updated
}
