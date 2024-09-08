package services

import (
	"goland-crud-gin/data/requests"
	"goland-crud-gin/data/responses"
	"goland-crud-gin/helper"
	"goland-crud-gin/model"
	"goland-crud-gin/repository"

	"github.com/go-playground/validator/v10"
)

type TagServiceImpl struct {
	TagRepository repository.TagRepository
	Validate      *validator.Validate
}

func NewTagServiceImpl(tagReository repository.TagRepository, validate *validator.Validate) TagService {
	return &TagServiceImpl{
		TagRepository: tagReository,
		Validate:      validate,
	}
}

// Create implements TagService.
func (t *TagServiceImpl) Create(tag requests.CreateTagRequest) {
	err := t.Validate.Struct(tag)
	helper.ErrorPanic(err)
	tagModel := model.TagModel{
		Name: tag.Name,
	}
	t.TagRepository.Save(tagModel)
}

// Delete implements TagService.
func (t *TagServiceImpl) Delete(tagId int) {
	t.TagRepository.Delete(tagId)
}

// FindMany implements TagService.
func (t *TagServiceImpl) FindMany() []responses.TagResponse {
	results := t.TagRepository.FindMany()
	var tags []responses.TagResponse
	for _, value := range results {
		tag := responses.TagResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindOneById implements TagService.
func (t *TagServiceImpl) FindOneById(tagId int) responses.TagResponse {
	result, err := t.TagRepository.FindOneById(tagId)
	helper.ErrorPanic(err)
	return responses.TagResponse{
		Id:   result.Id,
		Name: result.Name,
	}
}

// Update implements TagService.
func (t *TagServiceImpl) Update(tag requests.UpdateTagRequest) {
	result, err := t.TagRepository.FindOneById(tag.Id)
	helper.ErrorPanic(err)
	result.Name = tag.Name
	t.TagRepository.Update(result)
}
