package service

import (
	"example/web-service-gin/data/request"
	"example/web-service-gin/data/response"
	"example/web-service-gin/helper"
	"example/web-service-gin/model"
	"example/web-service-gin/repository"
	"github.com/go-playground/validator/v10"
)

type tagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &tagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

func (t *tagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t *tagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func (t *tagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

func (t *tagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagsData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagsData.Id,
		Name: tagsData.Name,
	}

	return tagResponse
}

func (t *tagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()
	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}
