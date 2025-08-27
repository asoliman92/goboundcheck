package testslices

func mapAccessShouldNotTrigger() string {
	m := map[string]int{"hello": 42}
	
	// These map accesses should NOT trigger the linter (maps don't have bounds)
	_ = m["nonexistent"] // This returns 0, doesn't panic - should be fine
	_ = m["another"]     // This should also be fine
	
	// This should also NOT trigger the linter
	if _, ok := m["key"]; ok {
		return "found"
	}
	
	// Nested map access should also be fine
	nested := map[string]map[string]int{
		"outer": {"inner": 123},
	}
	_ = nested["outer"]["inner"] // Should not trigger
	
	return "done"
}

func mixedMapAndSliceAccess() {
	m := map[string]int{"test": 1}
	slice := []int{1, 2, 3}
	
	// Map access - should NOT trigger linter
	_ = m["key"]
	
	// Slice access without bounds check - SHOULD trigger linter
	_ = slice[10] // want "Slice or array access is not enclosed in an if-statement that validates capacity!"
}

func properSliceBoundsCheck() {
	m := map[string]int{"test": 1}
	slice := []int{1, 2, 3}
	
	// Map access - should NOT trigger linter
	_ = m["key"]
	
	// Slice access WITH bounds check - should NOT trigger linter
	if 10 < len(slice) {
		_ = slice[10]
	}
}
