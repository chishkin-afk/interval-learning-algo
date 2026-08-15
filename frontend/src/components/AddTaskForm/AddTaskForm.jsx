import { useContext } from 'react'
import Field from '../Field/Field'
import styles from './AddTaskForm.module.scss'
import { TasksContext } from '../context/TasksContext'
import Button from '../Button/Button'

function AddTaskForm(props) {
    const {
        className = '',
    } = props

    const {
        newTaskTitle,
        setNewTaskTitle,
        newLeetcodeURL,
        setNewLeetcodeURL,
    } = useContext(TasksContext)

    const handleSubmit = (event) => {
        event.preventDefault()
        
        event.clear()
    }

    const onInputTitle = (event) => {
        setNewTaskTitle(event.target.value)
    }

    const onInputLeetcodeURL = (event) => {
        setNewLeetcodeURL(event.target.value)
    }

    return (
        <form 
            className={`${styles.form} ${className}`}
            onSubmit={handleSubmit}
        >
            <div className={styles.main}>
                <Field 
                    placeholder="new task title..."
                    id='new-task-title'
                    onInput={onInputTitle}
                    className={styles.input__title}
                    autoComplete="off"
                />
                <Button
                    type="submit"
                    className={styles.button}
                >
                    Add
                </Button>
            </div>
            {(newTaskTitle !== "" || newLeetcodeURL !== "") && (
                <Field 
                    placeholder="leetcode url..."
                    id='new-task-leetcode'
                    onInput={onInputLeetcodeURL}
                    className={styles.input__leetcode}
                    autoComplete="off"
                />
            )}
        </form>
    )
}

export default AddTaskForm