package services

import (
	"context"
	"time"
)

type Album struct {
	ID     		string `json:"id"`
	Title  		string `json:"title"`
	Description string `json:"description"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}
type AlbumsList struct {
	Albums []Album `json:"albums"`
}


func (a *Album) GetAllAlbums() ([]*Album, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, title, description, created_at, updated_at FROM albums`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var albums []*Album
	for rows.Next() {
		var album Album
		err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Description,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}
	return albums, nil
}