import { createContext, useContext } from 'react'
import type { Task } from './types'

interface TasksState {
	tasks: Task[]
	currentTask: Task | null
	pages: number
	currentPage: number
}

interface TasksActions {
	addTask(task: Task): void
	deleteTask(taskId: string): void
	updateTask(taskId: string, task: Task): Task | undefined
	setCurrentTask: React.Dispatch<React.SetStateAction<Task | null>>
	nextPage(): void
	prevPage(): void
}

export const TasksStateContext = createContext<TasksState | null>(null)
export const TasksActionsContext = createContext<TasksActions | null>(null)

export function useTasksState(): TasksState {
	const ctx = useContext(TasksStateContext)
	if (!ctx) throw new Error('useTasksState must be used in TasksProvider')

	return ctx
}

export function useTasksActions(): TasksActions {
	const ctx = useContext(TasksActionsContext)
	if (!ctx) throw new Error('useTasksActions must be used in TasksProvider')

	return ctx
}
