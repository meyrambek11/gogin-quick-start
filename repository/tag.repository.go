package repository

import "goland-crud-gin/model"

type TagRepository interface {
	Save(tag model.TagModel)
	Update(tag model.TagModel)
	Delete(id int)
	FindOneById(id int) (tag model.TagModel, err error)
	FindMany() []model.TagModel
}
