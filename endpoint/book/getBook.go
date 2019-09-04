package book

import (
	"context"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetBookLogic struct {
	Log        logging.Logger
	FileReader FileReaderWriter
}

type FileReaderWriter interface {
	Read() ([][]string, error)
}

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func (gl *GetBookLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	data, err := gl.FileReader.Read()
	if err != nil {
		gl.Log.LogErrorf("Could not read data file: %v", err)
		res.HTTPStatus = 400
	}
	books := make([]Book, len(data))
	for i, d := range data {
		books[i].Name = d[0]
		books[i].Author = d[1]
	}
	res.Body = books
	res.HTTPStatus = 200
}
