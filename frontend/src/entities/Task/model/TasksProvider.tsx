import React, { useCallback, useState } from 'react'
import { TasksActionsContext, TasksStateContext } from './TasksContext'
import type { Task } from './types'

interface TaskProviderProps {
	children: React.ReactNode
}

function TasksProvider({ children }: TaskProviderProps) {
	const [tasks, setTasks] = useState<Task[]>([])
	const [currentTask, setCurrentTask] = useState<Task | null>(null)

	const pages = 1000
	const [currentPage, setCurrentPage] = useState<number>(1)

	const addTask = useCallback((task: Task) => {
		setTasks((prevTasks) => [...prevTasks, task])
	}, [])

	const deleteTask = useCallback((taskId: string) => {
		setTasks((prevTasks) => prevTasks.filter(({ id }) => id !== taskId))
	}, [])

	const updateTask = useCallback(
		(taskId: string, upd: Task): Task | undefined => {
			let task: Task | undefined = undefined
			setTasks((prevTasks) => {
				task = prevTasks.find(({ id }) => id === taskId)
				if (!task) throw new Error(`Unknown task: ${taskId}`)
				task.description = upd.description
				task.title = upd.title
				task.is_active = upd.is_active
				task.leetcode_url = upd.leetcode_url

				task

				return prevTasks
			})

			return task
		},
		[tasks]
	)

	const nextPage = useCallback(() => {
		setCurrentPage((page) => page + 1)
	}, [])

	const prevPage = useCallback(() => {
		setCurrentPage((page) => page - 1)
	}, [])

	return (
		<TasksStateContext.Provider
			value={{
				tasks,
				currentTask,
				currentPage,
				pages,
			}}
		>
			<TasksActionsContext.Provider
				value={{
					addTask,
					deleteTask,
					updateTask,
					setCurrentTask,
					nextPage,
					prevPage,
				}}
			>
				{children}
			</TasksActionsContext.Provider>
		</TasksStateContext.Provider>
	)
}

export default TasksProvider
