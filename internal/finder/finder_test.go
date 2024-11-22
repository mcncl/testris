package finder

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindTests(t *testing.T) {
	t.Parallel()
	tmpDir := t.TempDir()

	testFiles := map[string]string{
		"package_test.go": `package example
			func TestOne(t *testing.T) {}
			func TestTwo(t *testing.T) {}
			func helper(t *testing.T) {}`,
		"other/nested_test.go": `package other
			func TestNested(t *testing.T) {}`,
	}

	for path, content := range testFiles {
		fullPath := filepath.Join(tmpDir, path)
		err := os.MkdirAll(filepath.Dir(fullPath), 0755)
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile(fullPath, []byte(content), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	tests, err := FindTests(tmpDir)
	if err != nil {
		t.Fatalf("FindTests failed: %v", err)
	}

	expectedTests := map[string]bool{
		"TestOne":    false,
		"TestTwo":    false,
		"TestNested": false,
	}

	for _, test := range tests {
		if _, ok := expectedTests[test.Name]; !ok {
			t.Errorf("Unexpected test found: %s", test.Name)
			continue
		}
		expectedTests[test.Name] = true
	}

	for name, found := range expectedTests {
		if !found {
			t.Errorf("Expected test not found: %s", name)
		}
	}
}
