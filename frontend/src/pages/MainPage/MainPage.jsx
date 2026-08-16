import { TasksProvider } from '@/context/TasksContext'
import Navbar from '../../components/Navbar/Navbar'
import TaskInfo from '../../components/TaskInfo/TaskInfo'
import TasksList from '../../components/TasksList/TasksList'
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
