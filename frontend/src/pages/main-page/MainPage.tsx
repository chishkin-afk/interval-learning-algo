import TasksProvider from '@/entities/task/model'
import Navbar from '@/widgets/navbar'
import TaskInfo from '@/widgets/task-info/TaskInfo'
import TasksList from '@/widgets/tasks-list'

function MainPage() {
	return (
		<TasksProvider>
			<Navbar />
			<div>
				<TasksList />
				<TaskInfo />
			</div>
		</TasksProvider>
	)
}

export default MainPage
