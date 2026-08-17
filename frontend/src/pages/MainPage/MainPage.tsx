import TasksProvider from '@/entities/Task/model'
import Tasks from '@/widgets/tasks'

function MainPage() {
	return (
		<TasksProvider>
			<Tasks />
		</TasksProvider>
	)
}

export default MainPage
