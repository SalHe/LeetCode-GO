package leetcode

type parenthesesTask struct {
	parentheses string
	stack       []int32
	pStack      int
}

// 1 <= n <= 8
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	}

	hashMap := make(map[string]interface{}, n) // 这个预置大小是按估计给的
	tasks, pTask := make([]*parenthesesTask, n), 0
	results, pResult := make([]string, n), 0

	// 容量不足，扩容
	checkTaskCap := func() {
		if pTask >= len(tasks) {
			newTasks := make([]*parenthesesTask, len(tasks)+n)
			copy(newTasks, tasks)
			tasks = newTasks
		}
	}
	insertResult := func(target string) {
		if _, ok := hashMap[target]; !ok {
			if pResult >= len(results) {
				nr := make([]string, len(results)+n)
				copy(nr, results)
				results = nr
			}
			results[pResult] = target
			hashMap[target] = nil
			pResult++
		}
	}

	// 初始化任务队列
	tasks[pTask] = &parenthesesTask{
		parentheses: "(",
		stack:       make([]int32, n),
		pStack:      0,
	}
	tasks[pTask].stack[0] = '('

	for pTask >= 0 {
		task := tasks[pTask]
		//tasks[pTask] = nil // 移除引用，GC
		pTask--

		// 加左括号
		if task.pStack < n-1 && len(task.parentheses) < 2*n-1 && task.pStack+1+len(task.parentheses) <= 2*n { // 左括号最多显然只能有n个，并且保证不会在末尾加入左括号
			pTask++

			checkTaskCap()

			if pTask >= len(tasks) {
				newTasks := make([]*parenthesesTask, len(tasks)+n)
				copy(newTasks, tasks)
				tasks = newTasks
			}

			newTask := &parenthesesTask{
				parentheses: task.parentheses + "(",
				stack:       make([]int32, n),
				pStack:      task.pStack + 1,
			}
			copy(newTask.stack, task.stack)
			newTask.stack[newTask.pStack] = '('
			tasks[pTask] = newTask
		}

		// 加右括号
		if task.pStack >= 0 { // 加右括号得保证栈内至少有一个左括号
			if len(task.parentheses) == 2*n-1 && task.pStack == 0 {
				insertResult(task.parentheses + ")")
			} else {
				pTask++

				checkTaskCap()

				newTask := &parenthesesTask{
					parentheses: task.parentheses + ")",
					stack:       make([]int32, n),
					pStack:      task.pStack - 1,
				}
				copy(newTask.stack, task.stack)
				tasks[pTask] = newTask
			}
		}
	}

	return results[0:pResult]
}

// 1 <= n <= 8
func generateParenthesis_version1(n int) []string {
	if n <= 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	}

	hashMap := make(map[string]interface{}, n) // 这个预置大小是按估计给的
	tasks, pTask := make([]*parenthesesTask, n), 0
	results, pResult := make([]string, n), 0

	// TODO 扩容方式可改进
	// 容量不足，扩容
	checkTaskCap := func() {
		if pTask >= len(tasks) {
			newTasks := make([]*parenthesesTask, len(tasks)+n)
			copy(newTasks, tasks)
			tasks = newTasks
		}
	}
	insertResult := func(target string) {
		if _, ok := hashMap[target]; !ok {
			if pResult >= len(results) {
				nr := make([]string, len(results)+n)
				copy(nr, results)
				results = nr
			}
			results[pResult] = target
			hashMap[target] = nil
			pResult++
		}
	}

	// 初始化任务队列
	tasks[pTask] = &parenthesesTask{
		parentheses: "(",
		stack:       make([]int32, n),
		pStack:      0,
	}
	tasks[pTask].stack[0] = '('

	for pTask >= 0 {
		task := tasks[pTask]
		//tasks[pTask] = nil // 移除引用，GC
		pTask--

		// 加左括号
		if task.pStack < n-1 && len(task.parentheses) < 2*n-1 && task.pStack+1+len(task.parentheses) <= 2*n { // 左括号最多显然只能有n个，并且保证不会在末尾加入左括号
			pTask++

			checkTaskCap()

			if pTask >= len(tasks) {
				newTasks := make([]*parenthesesTask, len(tasks)+n)
				copy(newTasks, tasks)
				tasks = newTasks
			}

			newTask := &parenthesesTask{
				parentheses: task.parentheses + "(",
				stack:       make([]int32, n),
				pStack:      task.pStack + 1,
			}
			copy(newTask.stack, task.stack)
			newTask.stack[newTask.pStack] = '('
			tasks[pTask] = newTask
		}

		// 加右括号
		if task.pStack >= 0 { // 加右括号得保证栈内至少有一个左括号
			if len(task.parentheses) == 2*n-1 && task.pStack == 0 {
				insertResult(task.parentheses + ")")
			} else {
				pTask++

				checkTaskCap()

				newTask := &parenthesesTask{
					parentheses: task.parentheses + ")",
					stack:       make([]int32, n),
					pStack:      task.pStack - 1,
				}
				copy(newTask.stack, task.stack)
				tasks[pTask] = newTask
			}
		}
	}

	return results[0:pResult]
}
