import { useState } from 'react'

function useNewTaskForm() {
	const [newTaskTitle, setNewTaskTitle] = useState('')
	const [newLeetcodeURL, setNewLeetcodeURL] = useState('')

	return {
		newTaskTitle,
		setNewTaskTitle,
		newLeetcodeURL,
		setNewLeetcodeURL,
	}
}

export default useNewTaskForm
