package service

import (
	"Users/witch/IdeaProjects/frameworktest/data/request"
	"Users/witch/IdeaProjects/frameworktest/data/response"
	"Users/witch/IdeaProjects/frameworktest/helper"
	"Users/witch/IdeaProjects/frameworktest/model"
	"Users/witch/IdeaProjects/frameworktest/repository"
	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	validate       *validator.Validate
}

func NewTagsServiceImpl(tagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{TagsRepository: tagsRepository, validate: validate}
}

func (t TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func (t TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

func (t TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

func (t TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tags = append(tags, response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		})
	}
	return tags
}
