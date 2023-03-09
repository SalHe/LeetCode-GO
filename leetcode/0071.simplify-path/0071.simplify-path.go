package leetcode

import "strings"

func simplifyPath(path string) string {
	dirs := strings.SplitN(path, "/", -1)
	for i := 0; i < len(dirs); {
		dir := dirs[i]
		if dir == "." || dir == "" {
			dirs = remove(dirs, i, i+1)
		} else if dir == ".." {
			dirs = remove(dirs, i-1, i+1)
			i--
		} else {
			i++
		}
		if i < 0 {
			i = 0
		}
	}
	return "/" + strings.Join(dirs, "/")
}

func remove(dirs []string, left, right int) []string {
	if left < 0 {
		left = 0
	}
	if right >= len(dirs) {
		right = len(dirs)
	}
	return append(dirs[:left], dirs[right:]...)
}

func simplifyPathCharByChar(path string) string {
	if path == "/" {
		return "/"
	}
	// 预分配空间
	bytes := make([]byte, len(path))
	dirNameLens := make([]int, (len(path)+1)/2)
	dir := 0
	p := 0

	back := func() {
		if dir > 0 {
			dir--
			p -= dirNameLens[dir]
			p--
			if p < 0 {
				p = 0
			}
		}
	}

	for i := 0; i < len(path); {
		for i < len(path) && path[i] == '/' {
			i++
		}
		start := i
		for i < len(path) && path[i] != '/' {
			i++
		}
		end := i

		dirNameLen := end - start
		if dirNameLen == 1 && path[start] == '.' {
			// do nothing
		} else if dirNameLen == 2 && path[start] == '.' && path[start+1] == '.' {
			back()
		} else if dirNameLen != 0 {
			dirNameLens[dir] = dirNameLen
			dir++
			bytes[p] = '/'
			p++
			copy(bytes[p:p+dirNameLen], path[start:end])
			p += dirNameLen
		}
	}

	if p < 1 {
		p = 1
	}
	bytes[0] = '/'
	return string(bytes[:p])
}
