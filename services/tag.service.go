package services

import (
	"goland-crud-gin/data/requests"
	"goland-crud-gin/data/responses"
)

type TagService interface {
	Create(tag requests.CreateTagRequest)
	Update(tag requests.UpdateTagRequest)
	FindMany() []responses.TagResponse
	FindOneById(tagId int) responses.TagResponse
	Delete(tagId int)
}
