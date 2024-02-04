package wordCount

import "testing"

func TestWorldCount(t *testing.T) {
	tests := []struct {
		text        string
		expectedMap map[string]int
	}{
		{
			text:        "word count test",
			expectedMap: map[string]int{"word": 1, "count": 1, "test": 1},
		},
		{
			text:        "abc def abc qwe asdf def abc",
			expectedMap: map[string]int{"abc": 3, "def": 2, "qwe": 1, "asdf": 1},
		},
	}

	for _, test := range tests {
		t.Run("test", func(t *testing.T) {
			wordMap := WordCount(test.text)

			for key, expectedCount := range test.expectedMap {
				actualCount, ok := wordMap[key]
				if !ok {
					t.Errorf("Expected key %s not found int wordMap", key)
					continue
				}
				if actualCount != expectedCount {
					t.Errorf("Word count mismatch for key %s. Expected: %d, Actual: %d", key, expectedCount, actualCount)
				}
			}

			t.Logf("Test passed")
		})
	}
}
