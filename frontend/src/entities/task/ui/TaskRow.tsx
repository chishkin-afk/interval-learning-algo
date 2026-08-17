import Button from '@/shared/ui/Button'
import { IconDelete, IconLink } from '@/shared/ui/Icon'
import { useTasksActions } from '../model'
import type { Task } from '../model/types'
import styles from './TaskRow.module.scss'

interface TaskProps {
	task: Task
}

function TaskRow(props: TaskProps) {
	const { task } = props

	const { setCurrentTask } = useTasksActions()

	return (
		<div className={styles.task} onClick={() => setCurrentTask(task)}>
			<div className={styles.title}>
				<h3>{task.title}</h3>
			</div>
			<div className={styles.actions}>
				<time dateTime={task.created_at.toISOString()} className={styles.createdAt}>
					{task.created_at.toISOString().slice(0, 10)}
				</time>
				<a className={`${styles.button} ${styles.link}`} href={task.leetcode_url}>
					<IconLink />
				</a>
				<Button className={styles.button}>
					<IconDelete />
				</Button>
			</div>
		</div>
	)
}

export default TaskRow
