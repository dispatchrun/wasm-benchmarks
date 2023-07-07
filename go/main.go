// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "unsafe"

//go:nosplit
func appendCleanPath(buf []byte, path string, lookupParent bool) ([]byte, bool) {
	i := 0
	for i < len(path) {
		for i < len(path) && path[i] == '/' {
			i++
		}

		j := i
		for j < len(path) && path[j] != '/' {
			j++
		}

		s := path[i:j]
		i = j

		switch s {
		case "":
			continue
		case ".":
			continue
		case "..":
			if !lookupParent {
				k := len(buf)
				for k > 0 && buf[k-1] != '/' {
					k--
				}
				for k > 1 && buf[k-1] == '/' {
					k--
				}
				buf = buf[:k]
				if k == 0 {
					lookupParent = true
				} else {
					s = ""
					continue
				}
			}
		default:
			lookupParent = false
		}

		if len(buf) > 0 && buf[len(buf)-1] != '/' {
			buf = append(buf, '/')
		}
		buf = append(buf, s...)
	}
	return buf, lookupParent
}

// joinPath concatenates dir and file paths, producing a cleaned path where
// "." and ".." have been removed, unless dir is relative and the references
// to parent directories in file represented a location relative to a parent
// of dir.
//
// This function is used for path resolution of all wasi functions expecting
// a path argument; the returned string is heap allocated, which we may want
// to optimize in the future. Instead of returning a string, the function
// could append the result to an output buffer that the functions in this
// file can manage to have allocated on the stack (e.g. initializing to a
// fixed capacity). Since it will significantly increase code complexity,
// we prefer to optimize for readability and maintainability at this time.
func joinPath(dir, file string) string {
	buf := make([]byte, 0, len(dir)+len(file)+1)
	if isAbs(dir) {
		buf = append(buf, '/')
	}
	buf, lookupParent := appendCleanPath(buf, dir, false)
	buf, _ = appendCleanPath(buf, file, lookupParent)
	// The appendCleanPath function cleans the path so it does not inject
	// references to the current directory. If both the dir and file args
	// were ".", this results in the output buffer being empty so we handle
	// this condition here.
	if len(buf) == 0 {
		buf = append(buf, '.')
	}
	// If the file ended with a '/' we make sure that the output also ends
	// with a '/'. This is needed to ensure that programs have a mechanism
	// to represent dereferencing symbolic links pointing to directories.
	if buf[len(buf)-1] != '/' && isDir(file) {
		buf = append(buf, '/')
	}
	return unsafe.String(&buf[0], len(buf))
}

func isAbs(path string) bool {
	return hasPrefix(path, "/")
}

func isDir(path string) bool {
	return hasSuffix(path, "/")
}

func hasPrefix(s, p string) bool {
	return len(s) >= len(p) && s[:len(p)] == p
}

func hasSuffix(s, x string) bool {
	return len(s) >= len(x) && s[len(s)-len(x):] == x
}
