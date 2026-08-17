import { useTasksState } from '@/entities/task/model'
import TaskRow from '@/entities/task/ui'
import AddTaskForm from '@/features/AddTaskForm'
import Pagination from '../pagination'
import styles from './TasksList.module.scss'

function TasksList() {
	const { tasks = [] } = useTasksState()

	return (
		<section className={styles.tasks}>
			<ul className={styles.list}>
				<li>
					<AddTaskForm />
				</li>

				{tasks.map((task, index) => (
					<li
						key={task.id}
						className={styles.item}
						style={{ '--delay': `${index * 50}ms` } as React.CSSProperties}
					>
						<TaskRow task={task} />
					</li>
				))}
			</ul>
			<Pagination />
		</section>
	)
}

export default TasksList
