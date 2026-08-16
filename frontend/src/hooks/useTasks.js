import { useCallback, useEffect, useState } from 'react'

const MOCK_INIT_STATE = [
	{
		id: '536a5460-230a-41d2-bec6-a99e5cc66576',
		user_id: '3cdba3ad-ecf1-49e1-9313-50aaf028fbe2',
		title: 'Polish notation',
		is_active: true,
		description: 'Some description...',
		leetcode_url: 'https://leetcode.com/problems/evaluate-reverse-polish-notation',
		notification_count: 2,
		next_notification: new Date('2026-08-16T15:04:05.999999999Z'),
		created_at: new Date('2026-08-12T15:04:05.999999999Z'),
	},
	{
		id: 'b4c89be9-22a7-409a-a030-d9651c54a577',
		user_id: '3cdba3ad-ecf1-49e1-9313-50aaf028fbe2',
		title: 'Two sum',
		is_active: true,
		description: 'Some description...',
		leetcode_url: 'https://leetcode.com/problems/two-sum',
		notification_count: 2,
		next_notification: new Date('2026-08-16T15:04:05.999999999Z'),
		created_at: new Date('2026-08-12T15:04:05.999999999Z'),
	},
	{
		id: '0a01a8e9-426a-4568-b76a-0b499feb43eb',
		user_id: '3cdba3ad-ecf1-49e1-9313-50aaf028fbe2',
		title: 'Basic calculator IV',
		is_active: true,
		description: 'Some description...',
		leetcode_url: 'https://leetcode.com/problems/basic-calculator-iv',
		notification_count: 2,
		next_notification: new Date('2026-08-16T15:04:05.999999999Z'),
		created_at: new Date('2026-08-12T15:04:05.999999999Z'),
	},
	{
		id: '536a5460-230a-41d2-bec6-a99e5cc66577',
		user_id: '3cdba3ad-ecf1-49e1-9313-50aaf028fbe2',
		title: 'Polish notation',
		is_active: true,
		description: 'Some description...',
		leetcode_url: 'https://leetcode.com/problems/evaluate-reverse-polish-notation',
		notification_count: 2,
		next_notification: new Date('2026-08-16T15:04:05.999999999Z'),
		created_at: new Date('2026-08-12T15:04:05.999999999Z'),
	},
	{
		id: 'b4c89be9-22a7-409a-a030-d9651c54a578',
		user_id: '3cdba3ad-ecf1-49e1-9313-50aaf028fbe2',
		title: 'Two sum',
		is_active: true,
		description: 'Some description...',
		leetcode_url: 'https://leetcode.com/problems/two-sum',
		notification_count: 2,
		next_notification: new Date('2026-08-16T15:04:05.999999999Z'),
		created_at: new Date('2026-08-12T15:04:05.999999999Z'),
	},
	{
		id: '0a01a8e9-426a-4568-b76a-0b499feb43e9',
		user_id: '3cdba3ad-ecf1-49e1-9313-50aaf028fbe2',
		title: 'Basic calculator IV',
		is_active: true,
		description: 'Some description...',
		leetcode_url: 'https://leetcode.com/problems/basic-calculator-iv',
		notification_count: 2,
		next_notification: new Date('2026-08-16T15:04:05.999999999Z'),
		created_at: new Date('2026-08-12T15:04:05.999999999Z'),
	},
	{
		id: '536a5460-230a-41d2-bec6-a99e5cc66570e',
		user_id: '3cdba3ad-ecf1-49e1-9313-50aaf028fbe2',
		title: 'Polish notation',
		is_active: true,
		description: 'Some description...',
		leetcode_url: 'https://leetcode.com/problems/evaluate-reverse-polish-notation',
		notification_count: 2,
		next_notification: new Date('2026-08-16T15:04:05.999999999Z'),
		created_at: new Date('2026-08-12T15:04:05.999999999Z'),
	},
]

const PAGE_SIZE = 5

function useTasks() {
	const [tasks, setTasks] = useState(MOCK_INIT_STATE)
	const [currentTask, setCurrentTask] = useState(tasks[0] ?? null)

	const [page, setPage] = useState(1)

	const totalPages = 10

	const addTask = useCallback((task) => {
		console.log('adding task')
	}, [])

	const deleteTask = useCallback((taskId) => {
		console.log('deleting task:', taskId)
	}, [])

	const updateTask = useCallback((taskId, updTask) => {
		console.log('editing task:', taskId)
	}, [])

	useEffect(() => {
		console.log('changing page...')
	}, [page])

	const nextPage = useCallback(() => {
		if (page + 1 > totalPages) {
			console.log('Nothing to change')
			return
		}

		setPage((prev) => prev + 1)
	}, [page])

	const prevPage = useCallback(() => {
		if (page - 1 < 1) {
			console.log('Nothing to change')
			return
		}

		setPage((prev) => prev - 1)
	}, [page])

	return {
		tasks,
		currentTask,
		page,
		totalPages,

		addTask,
		deleteTask,
		updateTask,

		setCurrentTask,

		nextPage,
		prevPage,
	}
}

export default useTasks
