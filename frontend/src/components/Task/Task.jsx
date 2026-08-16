import Button from '../Button/Button'
import IconDelete from '../IconDelete/IconDelete'
import IconLink from '../IconLink/IconLink'
import styles from './Task.module.scss'

function Task(props) {
	const { task } = props

	return (
		<div className={styles.task}>
			<div className={styles.title}>
				<h3>{task.title}</h3>
			</div>
			<div className={styles.actions}>
				<time dateTime={task.created_at.toISOString()} className={styles.createdAt}>
					{task.created_at.toISOString().slice(0, 10)}
				</time>
				<a className={`${styles.button} ${styles.link}`} href={task.leetcode_url}>
					<IconLink />
				</a>
				<Button className={styles.button}>
					<IconDelete />
				</Button>
			</div>
		</div>
	)
}

export default Task
