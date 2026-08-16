import { TasksContext } from '@/context/TasksContext'
import { useContext } from 'react'
import AddTaskForm from '../AddTaskForm/AddTaskForm'
import Pagination from '../Pagination/Pagination'
import Task from '../Task/Task'
import styles from './TasksList.module.scss'

function TasksList() {
	const { tasks = [] } = useContext(TasksContext)

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
						style={{ '--delay': `${index * 50}ms` }}
					>
						<Task task={task} />
					</li>
				))}
			</ul>
			<Pagination />
		</section>
	)
}

export default TasksList
