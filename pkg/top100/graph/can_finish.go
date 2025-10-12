package graph

// NO.207
// 拓扑排序 + 广度优先搜索，
func canFinish(numCourses int, prerequisites [][]int) bool {
	//定义一个数组，保存每一个课程的入度
	var courses []int
	for i := 0; i < numCourses; i++ {
		courses = append(courses, 0)
	}

	//根据课程先后顺序获取课程的入度
	for _, v := range prerequisites {
		courses[v[1]]++
	}

	//待学习的课程， 入度为0的课程入队
	var queue []int

	for i, c := range courses {
		if c == 0 {
			queue = append(queue, i)
		}
	}

	//开始学习，
	var ans []int
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		ans = append(ans, c)

		for _, v := range prerequisites {
			if v[0] == c {
				courses[v[1]]--
				if courses[v[1]] == 0 {
					queue = append(queue, v[1])
				}
			}
		}
	}

	if len(ans) != numCourses {
		return false
	}

	return true
}
