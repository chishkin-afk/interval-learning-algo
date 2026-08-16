import { useContext, useEffect, useState } from 'react'
import { TasksContext } from '../../context/TasksContext'
import Button from '../Button/Button'
import Field from '../Field/Field'
import styles from './TaskInfo.module.scss'

function TaskInfo() {
	const { currentTask } = useContext(TasksContext)

	if (!currentTask) {
		return (
			<section className={styles.info}>
				<h3>No tasks selected</h3>
			</section>
		)
	}

	const [description, setDescription] = useState(currentTask.description)
	useEffect(() => {
		setDescription(currentTask.description)
	}, [currentTask])

	const [leetcodeURL, setLeetcodeURL] = useState(currentTask.leetcode_url)
	useEffect(() => {
		setLeetcodeURL(currentTask.leetcode_url)
	}, [currentTask])

	const handleDescription = ({ target }) => {
		setDescription(target.value)
	}

	const onInput = () => {}

	return (
		<section className={styles.info}>
			<h3>{currentTask.title}</h3>
			<Field placeholder="change url..." value={leetcodeURL} onInput={onInput} />
			<textarea
				placeholder="change description..."
				name="description"
				id="new-description"
				className={styles.description}
				value={description}
				onChange={handleDescription}
			/>
			<div className={styles.stats}>
				<p>
					Next notify:{' '}
					<span className={styles.stats__value}>
						{currentTask.created_at.toISOString().slice(0, 10)}
					</span>
				</p>
				<p>
					Notify count:{' '}
					<span className={styles.stats__value}>{currentTask.notification_count}</span>
				</p>
				<p>
					Acitve:{' '}
					<span className={styles.stats__value}>
						{currentTask.is_active ? 'true' : 'false'}
					</span>
				</p>
			</div>
			<div className={styles.actions}>
				<Button className={`${styles.button} ${styles.actions__disable}`}>disable</Button>
				<div className={styles.actions__ud}>
					<Button className={styles.button}>save</Button>
					<Button className={`${styles.button} ${styles.delete}`}>del</Button>
				</div>
			</div>
		</section>
	)
}

export default TaskInfo
