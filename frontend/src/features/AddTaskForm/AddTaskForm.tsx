import { useTasksActions, useTasksState } from '@/entities/Task/model/TasksContext'
import Button from '@/shared/ui/Button'
import Field from '@/shared/ui/Field'
import type React from 'react'
import styles from './AddTaskForm.module.scss'

interface AddTaskFormProps {
	className?: string
}

function AddTaskForm({ className }: AddTaskFormProps) {
	const { newTaskTitle, setNewTaskTitle, newLeetcodeURL, setNewLeetcodeURL } = useTasksState()
	const { addTask } = useTasksActions()

	const handleSubmit = (event: React.SubmitEvent<HTMLFormElement>) => {
		event.preventDefault()
		addTask(newTaskTitle, newLeetcodeURL)
		setNewTaskTitle('')
		setNewLeetcodeURL('')
	}

	const onInputTitle = ({ currentTarget }: React.InputEvent<HTMLInputElement>) => {
		setNewTaskTitle(currentTarget.value)
	}

	const onInputLeetcodeURL = ({ currentTarget }: React.InputEvent<HTMLInputElement>) => {
		setNewLeetcodeURL(currentTarget.value)
	}

	return (
		<form className={`${styles.form} ${className}`} onSubmit={handleSubmit}>
			<div className={styles.main}>
				<Field
					placeholder="new task title..."
					id="new-task-title"
					onInput={onInputTitle}
					className={styles.input__title}
					autoComplete="off"
					value={newTaskTitle}
				/>
				<Button type="submit" className={styles.button}>
					Add
				</Button>
			</div>
			{(newTaskTitle !== '' || newLeetcodeURL !== '') && (
				<Field
					placeholder="leetcode url..."
					id="new-task-leetcode-url"
					onInput={onInputLeetcodeURL}
					className={styles.input__leetcode}
					autoComplete="off"
					value={newLeetcodeURL}
				/>
			)}
		</form>
	)
}

export default AddTaskForm
