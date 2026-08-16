import useNewTaskForm from '@/hooks/useNewTaskForm'
import { createContext } from 'react'
import useTasks from '../hooks/useTasks'

export const TasksContext = createContext({})

export const TasksProvider = ({ children }) => {
	const {
		tasks,
		currentTask,
		page,
		totalPages,
		addTask,
		deleteTask,
		updateTask,
		nextPage,
		prevPage,
		setCurrentTask,
	} = useTasks()
	const { newTaskTitle, setNewTaskTitle, newLeetcodeURL, setNewLeetcodeURL } = useNewTaskForm()

	return (
		<TasksContext.Provider
			value={{
				tasks,
				currentTask,
				page,
				totalPages,
				addTask,
				deleteTask,
				updateTask,

				newTaskTitle,
				setNewTaskTitle,
				newLeetcodeURL,
				setNewLeetcodeURL,

				setCurrentTask,

				nextPage,
				prevPage,
			}}
		>
			{children}
		</TasksContext.Provider>
	)
}
