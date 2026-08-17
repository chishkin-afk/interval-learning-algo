import TasksProvider from '@/entities/Task/model'
import Tasks from '@/widgets/Tasks'

function MainPage() {
	return (
		<TasksProvider>
			<Tasks />
		</TasksProvider>
	)
}

export default MainPage
