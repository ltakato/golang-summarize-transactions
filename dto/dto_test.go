package dto

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	t.Run("should replace nil ID and Name with '(uncategorized)' and normalize money value", func(t *testing.T) {
		category := &CategoryResponse{
			ID:          "",
			Name:        "",
			TotalAmount: 10050,
		}
		expectedUncategorizedToken := CategoryResponseText(UncategorizedCategoryToken)

		category.Normalize()

		if category.ID != expectedUncategorizedToken {
			t.Errorf("expected ID to be %q, but got %v", expectedUncategorizedToken, category.ID)
		}
		if category.Name != expectedUncategorizedToken {
			t.Errorf("expected ID to be %q, but got %v", expectedUncategorizedToken, category.ID)
		}
	})
}
