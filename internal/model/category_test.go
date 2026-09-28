package model

import "testing"

func TestCategoryValid(t *testing.T) {
	tests := []struct {
		name     string
		category Category
		valid    bool
	}{
		{
			name:     "generic",
			category: CategoryGeneric,
			valid:    true,
		},
		{
			name:     "cache",
			category: CategoryCache,
			valid:    true,
		},
		{
			name:     "logs",
			category: CategoryLogs,
			valid:    true,
		},
		{
			name:     "downloads",
			category: CategoryDownloads,
			valid:    true,
		},
		{
			name:     "trash",
			category: CategoryTrash,
			valid:    true,
		},
		{
			name:     "developer",
			category: CategoryDeveloper,
			valid:    true,
		},
		{
			name:     "unknown",
			category: Category("unknown"),
			valid:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.category.Valid(); got != tt.valid {
				t.Fatalf(
					"Category(%q).Valid() = %v, want %v",
					tt.category,
					got,
					tt.valid,
				)
			}
		})
	}
}
