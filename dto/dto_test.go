package dto

import "testing"

func TestNormalize(t *testing.T) {
	t.Run("should replace nil ID and Name with '(uncategorized)' and normalize money value", func(t *testing.T) {
		category := &CategoryResponse{
			ID:          nil,
			Name:        nil,
			TotalAmount: 10050,
		}
		expectedUncategorizedToken := UncategorizedCategoryToken
		expectedTotalAmount := float32(100.50)

		category.Normalize()

		if category.ID == nil || *category.ID != expectedUncategorizedToken {
			t.Errorf("expected ID to be %q, but got %v", expectedUncategorizedToken, category.ID)
		}
		if category.Name == nil || *category.Name != expectedUncategorizedToken {
			t.Errorf("expected ID to be %q, but got %v", expectedUncategorizedToken, category.ID)
		}
		if category.TotalAmount != expectedTotalAmount {
			t.Errorf("expected totalAmount to be %f, but got %v", expectedTotalAmount, category.ID)
		}
	})
}
