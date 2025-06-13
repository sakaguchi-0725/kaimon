package usecase

import "context"

type (
	GetShoppingItems interface {
		Execute(ctx context.Context, in GetShoppingItemsInput) (GetShoppingItemsOutput, error)
	}

	GetShoppingItemsInput struct {
		UserID     string
		GroupID    string
		Status     string
		CategoryID *int64
		Limit      int64
		Offset     int64
	}

	GetShoppingItemsOutput struct {
		Items      []ShoppingItem
		TotalCount int64
		Offset     int64
	}

	ShoppingItem struct {
		ID       int64
		Name     string
		MemberID string
		Status   string
	}

	getShoppingItemsInteractor struct{}
)

func (g *getShoppingItemsInteractor) Execute(ctx context.Context, in GetShoppingItemsInput) (GetShoppingItemsOutput, error) {
	// TODO: implement get shopping items logic
	items := []ShoppingItem{
		{
			ID:       1,
			Name:     "Item 1",
			MemberID: "1",
			Status:   "UNPURCHASED",
		},
		{
			ID:       2,
			Name:     "Item 2",
			MemberID: "2",
			Status:   "PURCHASED",
		},
		{
			ID:       3,
			Name:     "Item 3",
			MemberID: "2",
			Status:   "UNPURCHASED",
		},
		{
			ID:       4,
			Name:     "Item 4",
			MemberID: "1",
			Status:   "IN_CART",
		},
	}

	return GetShoppingItemsOutput{
		Items:      items,
		TotalCount: int64(len(items)),
		Offset:     in.Offset,
	}, nil
}

func NewGetShoppingItems() GetShoppingItems {
	return &getShoppingItemsInteractor{}
}
