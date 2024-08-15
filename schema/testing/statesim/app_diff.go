package statesim

import (
	"fmt"

	"cosmossdk.io/schema/view"
)

// DiffAppStates compares the app state of two objects that implement AppState and returns a string with a diff if they
// are different or the empty string if they are the same.
func DiffAppStates(expected, actual view.AppState) string {
	res := ""

	expectNumModules, err := expected.NumModules()
	if err != nil {
		res += fmt.Sprintf("ERROR getting expected num modules: %s\n", err)
		return res
	}

	actualNumModules, err := actual.NumModules()
	if err != nil {
		res += fmt.Sprintf("ERROR getting actual num modules: %s\n", err)
		return res
	}

	if expectNumModules != actualNumModules {
		res += fmt.Sprintf("MODULE COUNT ERROR: expected %d, got %d\n", expectNumModules, actualNumModules)
	}

	expected.Modules(func(expectedMod view.ModuleState, err error) bool {
		if err != nil {
			res += fmt.Sprintf("ERROR getting expected module: %s\n", err)
			return true
		}

		moduleName := expectedMod.ModuleName()
		actualMod, err := actual.GetModule(moduleName)
		if err != nil {
			res += fmt.Sprintf("ERROR getting actual module: %s\n", err)
			return true
		}
		if actualMod == nil {
			res += fmt.Sprintf("Module %s: actual module NOT FOUND\n", moduleName)
			return true
		}

		diff := DiffModuleStates(expectedMod, actualMod)
		if diff != "" {
			res += "Module " + moduleName + "\n"
			res += indentAllLines(diff)
		}

		return true
	})

	return res
}
