package repository

import "Users/witch/IdeaProjects/frameworktest/model"

type TagsRepository interface {
	Save(tags model.Tags)
	Update(tags model.Tags)
	Delete(tagsId int)
	FindById(tagsId int) (tags model.Tags, err error)
	FindAll() (tags []model.Tags)
}
