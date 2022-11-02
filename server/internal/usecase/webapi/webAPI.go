package webapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"thumbs/server/internal/entity"
	"thumbs/server/internal/usecase"
	"thumbs/server/pkg/webclient"
)

// YtAPI implements usecase.ThumbWebAPI
type YtAPI struct {
	c *webclient.Conn
}

var _ usecase.ThumbWebAPI = (*YtAPI)(nil)

// New is a constructor for YtAPI
func New(c *webclient.Conn) *YtAPI {
	return &YtAPI{
		c: c,
	}
}

// GetThumbFromAPI return entity.Pic from YouTube API, entity.ErrNotFound if there is no such id
func (api *YtAPI) GetThumbFromAPI(ctx context.Context, id string) (entity.Pic, error) {
	url := "https://img.youtube.com/vi/" + id + "/0.jpg"
	res, err := api.c.W.Get(url)
	if err != nil {
		return entity.Pic{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return entity.Pic{}, entity.ErrNotFound
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return entity.Pic{}, err
	}
	return entity.Pic{ID: id, Data: b}, nil
}
