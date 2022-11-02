package webAPI

import (
	"context"
	"io/ioutil"
	"net/http"
	"thumbs/server/internal/entity"
	"thumbs/server/pkg/webclient"
)

// YtApi implements usecase.ThumbWebAPI
type YtApi struct {
	c *webclient.Conn
}

// New is a constructor for YtApi
func New(c *webclient.Conn) *YtApi {
	return &YtApi{
		c: c,
	}
}

// GetThumbFromAPI return entity.Pic from YouTube API, entity.ErrNotFound if there is no such id
func (api *YtApi) GetThumbFromAPI(ctx context.Context, id string) (entity.Pic, error) {
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
