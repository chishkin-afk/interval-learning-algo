import TasksProvider from '@/entities/task/model'
import Navbar from '@/widgets/navbar'
import TaskInfo from '@/widgets/task-info/TaskInfo'
import TasksList from '@/widgets/tasks-list'
import styles from './MainPage.module.scss'

function MainPage() {
	return (
		<TasksProvider>
			<Navbar />
			<main className={styles.main}>
				<TasksList />
				<TaskInfo />
			</main>
		</TasksProvider>
	)
}

export default MainPage
