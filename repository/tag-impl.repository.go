package repository

import (
	"errors"
	"goland-crud-gin/data/requests"
	"goland-crud-gin/helper"
	"goland-crud-gin/model"

	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagRepositoryImpl(Db *gorm.DB) TagRepository {
	return &TagRepositoryImpl{Db: Db}
}

// Delete implements TagRepository.
func (t *TagRepositoryImpl) Delete(id int) {
	var tag model.TagModel
	result := t.Db.Where("id = ?", id).Delete(&tag)
	helper.ErrorPanic(result.Error)
}

// FindMany implements TagRepository.
func (t *TagRepositoryImpl) FindMany() []model.TagModel {
	var tags []model.TagModel
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindOneById implements TagRepository.
func (t *TagRepositoryImpl) FindOneById(id int) (tag model.TagModel, err error) {
	var tagItem model.TagModel
	result := t.Db.Find(&tagItem, id)

	if result == nil {
		return tagItem, errors.New("tag not found")
	}

	return tagItem, nil
}

// Save implements TagRepository.
func (t *TagRepositoryImpl) Save(tag model.TagModel) {
	result := t.Db.Create(&tag)
	helper.ErrorPanic(result.Error)
}

// Update implements TagRepository.
func (t *TagRepositoryImpl) Update(tag model.TagModel) {
	var updateTag = requests.UpdateTagRequest{
		Id:   tag.Id,
		Name: tag.Name,
	}
	result := t.Db.Model(&tag).Updates(updateTag)

	helper.ErrorPanic(result.Error)
}
