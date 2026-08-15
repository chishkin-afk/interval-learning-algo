import linkIcon from '@/assets/icons/link-icon.svg'
import deleteIcon from '@/assets/icons/delete-icon.svg'
import Button from '../Button/Button'
import styles from './Task.module.scss'
import IconLink from '../IconLink/IconLink'
import IconDelete from '../IconDelete/IconDelete'

function Task(props) {
    const {
        task,
    } = props

    return (
        <div className={styles.task}>
            <div className={styles.title}>
                <h3>{task.title}</h3>
            </div>
            <div className={styles.actions}>
                <h3 className={styles.createdAt}>
                    {task.created_at.toISOString().slice(0, 10)}
                </h3>
                <Button
                    className={styles.button}
                >
                    <IconLink />
                </Button>
                <Button
                    className={styles.button}
                >
                    <IconDelete />
                </Button>
            </div>
        </div>
    )
}

export default Task