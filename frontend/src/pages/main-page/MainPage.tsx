import TasksProvider from '@/entities/task/model'
import TaskInfo from '@/widgets/task-info/TaskInfo'
import TasksList from '@/widgets/tasks-list'

function MainPage() {
	return (
		<TasksProvider>
			<div>
				<TasksList />
				<TaskInfo />
			</div>
		</TasksProvider>
	)
}

export default MainPage
