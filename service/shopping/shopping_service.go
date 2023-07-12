package shopping

import (
	context "context"
	"fmt"
)

type service struct{}

func NewShoppingService() service {
	return service{}
}

func (service) mustEmbedUnimplementedShoppingServiceServer() {}

func (service) AddItem(context.Context, *AddItemRequest) (*AddItemResponse, error) {
	return nil, nil
}

func (service) UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, nil
}

func (service) RemoveItem(context.Context, *RemoveItemRequest) (*RemoveItemResponse, error) {
	return nil, nil
}

func (service) GetItems(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	fmt.Println("test")
	return nil, nil
}
