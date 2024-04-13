// Copyright 2023 W192547975
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This package provides functions for handling environment variables.
package env

// Set inserts an environment variable into a set of environment variables.
//
// If there is no key in the environment variable, the function will return nil.
//
// Incorrect elements in the set will be removed.
//
// If the value of the environment variable is empty, all elements in the set
// that have the same key as the environment variable will be removed.
//
// Otherwise, all elements in the set that have the same key as the environment
// variable will be updated to an element that is the same as the environment variable.
func Set(env []string, new string) []string {
	newkey, newvalue := Split(new)
	if newkey == "" {
		return nil
	}
	var old, real int
	for ; old < len(env); old++ {
		oldkey, oldvalue := Split(env[old])
		if oldkey == "" || oldvalue == "" {
			continue // remove
		}
		if oldkey == newkey {
			if newvalue == "" {
				continue // remove
			}
			env[old] = new // replace
			newvalue = ""  // clear others
		}
		env[real] = env[old] // move
		real++
	}
	env = env[:real]
	if newvalue != "" {
		env = append(env, new) // new
	}
	return env
}

// Split divides an environment variable into a key and a value along the first "=" in it.
//
// If the parameter does not contain "=",
// all return values of this function are empty strings.
func Split(env string) (key string, value string) {
	var r bool // for handling unusual strings
	for i, j := range env {
		if r {
			value = env[i:]
			break
		}
		if j == '=' {
			key = env[:i]
			r = true
		}
	}
	return
}
