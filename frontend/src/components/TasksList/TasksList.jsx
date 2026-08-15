import { useContext } from 'react'
import styles from './TasksList.module.scss'
import { TasksContext } from '../context/TasksContext'
import Task from '../Task/Task'
import AddTaskForm from '../AddTaskForm/AddTaskForm'

function TasksList() {
    const {
        tasks = [],
        newTaskTitle
    } = useContext(TasksContext)

    return (
        <section className={styles.tasks}>
            <ul className={styles.list}>
                <li>
                    <AddTaskForm />
                </li>

                {tasks.map((task) => (
                    <li key={task.id}>
                        <Task task={task}/>
                    </li>
                ))}
            </ul>
        </section>
    )
}

export default TasksList