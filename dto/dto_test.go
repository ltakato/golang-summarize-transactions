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

		result := category.Normalize()

		if result.ID == nil || *result.ID != expectedUncategorizedToken {
			t.Errorf("expected ID to be %q, but got %v", expectedUncategorizedToken, result.ID)
		}
		if result.Name == nil || *result.Name != expectedUncategorizedToken {
			t.Errorf("expected ID to be %q, but got %v", expectedUncategorizedToken, result.ID)
		}
		if result.TotalAmount != expectedTotalAmount {
			t.Errorf("expected totalAmount to be %f, but got %v", expectedTotalAmount, result.ID)
		}
	})
}
