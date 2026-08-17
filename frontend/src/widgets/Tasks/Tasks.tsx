import type { Task } from '@/entities/Task/model/types'
import TaskRow from '@/entities/Task/ui'

function Tasks() {
	const task: Task = {
		id: 'b4c89be9-22a7-409a-a030-d9651c54a577',
		user_id: '3cdba3ad-ecf1-49e1-9313-50aaf028fbe2',
		title: 'Two sum',
		is_active: true,
		description: 'Some description...',
		leetcode_url: 'https://leetcode.com/problems/two-sum',
		notification_count: 2,
		next_notification: new Date('2026-08-16T15:04:05.999999999Z'),
		created_at: new Date('2026-08-12T15:04:05.999999999Z'),
	}

	return (
		<div>
			<TaskRow task={task} />
		</div>
	)
}

export default Tasks
