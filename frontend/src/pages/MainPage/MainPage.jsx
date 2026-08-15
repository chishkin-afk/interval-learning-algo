import { TasksProvider } from "../../components/context/TasksContext"
import TaskInfo from "../../components/TaskInfo/TaskInfo"
import TasksList from "../../components/TasksList/TasksList"
import styles from './MainPage.module.scss'

function MainPage() {
    return (
        <TasksProvider>
            <main className={styles.main}>
                <TasksList />
                <TaskInfo />
            </main>
        </TasksProvider>
    )
}

export default MainPage