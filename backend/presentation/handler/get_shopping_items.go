package handler

import (
	"backend/core"
	"backend/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	GetShoppingItemsResponse struct {
		Items      []ShoppingItem `json:"items"`
		TotalCount int64          `json:"totalCount"`
		HasNext    bool           `json:"hasNext"`
	}

	ShoppingItem struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		MemberID string `json:"memberId"`
		Status   string `json:"status"`
	}
)

func NewGetShoppingItems(usecase usecase.GetShoppingItems) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, err := makeGetShoppingItemsInput(c)
		if err != nil {
			return err
		}

		output, err := usecase.Execute(c.Request().Context(), input)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, makeGetShoppingItemsResponse(output))
	}
}

func makeGetShoppingItemsInput(c echo.Context) (usecase.GetShoppingItemsInput, error) {
	var (
		categoryID *int64
		limit      int64 = 20
		offset     int64 = 0
	)

	categoryIdParam := c.QueryParam("categoryId")
	if categoryIdParam != "" {
		id, err := strconv.ParseInt(categoryIdParam, 10, 64)
		if err != nil {
			return usecase.GetShoppingItemsInput{}, core.NewInvalidError(err).
				WithMessage("カテゴリーIDが不正です")
		}
		categoryID = &id
	}

	limitParam := c.QueryParam("limit")
	if limitParam != "" {
		limit, err := strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			return usecase.GetShoppingItemsInput{}, core.NewInvalidError(err).
				WithMessage("limitパラメータが不正です")
		}
		if limit <= 0 {
			limit = 20 // デフォルト値に戻す
		}
	}

	offsetParam := c.QueryParam("offset")
	if offsetParam != "" {
		offset, err := strconv.ParseInt(offsetParam, 10, 64)
		if err != nil {
			return usecase.GetShoppingItemsInput{}, core.NewInvalidError(err).
				WithMessage("offsetパラメータが不正です")
		}
		if offset < 0 {
			offset = 0 // 負の値は0にする
		}
	}

	return usecase.GetShoppingItemsInput{
		UserID:     core.GetUserID(c.Request().Context()),
		GroupID:    c.Param("id"),
		Status:     c.QueryParam("status"),
		CategoryID: categoryID,
		Limit:      limit,
		Offset:     offset,
	}, nil
}

func makeGetShoppingItemsResponse(output usecase.GetShoppingItemsOutput) GetShoppingItemsResponse {
	items := make([]ShoppingItem, len(output.Items))

	for i, item := range output.Items {
		items[i] = ShoppingItem{
			ID:       item.ID,
			Name:     item.Name,
			MemberID: item.MemberID,
			Status:   item.Status,
		}
	}

	return GetShoppingItemsResponse{
		Items:      items,
		TotalCount: output.TotalCount,
		HasNext:    len(output.Items) > 0 && int64(len(output.Items))+output.Offset < output.TotalCount,
	}
}
