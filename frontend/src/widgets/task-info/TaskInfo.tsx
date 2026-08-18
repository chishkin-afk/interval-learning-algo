import { useTasksState } from '@/entities/task/model'
import Button from '@/shared/ui/Button'
import Field from '@/shared/ui/Field'
import React, { useEffect, useState } from 'react'
import styles from './TaskInfo.module.scss'

function TaskInfo() {
	const { currentTask } = useTasksState()
	const [description, setDescription] = useState('')

	const [leetcodeURL, setLeetcodeURL] = useState('')

	useEffect(() => {
		setLeetcodeURL(currentTask?.leetcode_url ?? '')
	}, [currentTask])

	useEffect(() => {
		setDescription(currentTask?.description ?? '')
	}, [currentTask])

	if (!currentTask) {
		return (
			<section className={styles.info}>
				<h3>No tasks selected</h3>
			</section>
		)
	}

	const handleDescription = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
		setDescription(event.currentTarget.value)
	}

	const handleLeetcodeURL = (event: React.ChangeEvent<HTMLInputElement>) => {
		setLeetcodeURL(event.currentTarget.value)
	}

	return (
		<section className={styles.info}>
			<h3>{currentTask.title}</h3>
			<Field
				id="new-leetcode-url"
				placeholder="change url..."
				value={leetcodeURL}
				onChange={handleLeetcodeURL}
				autoComplete="off"
			/>
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
						{currentTask.next_notification.toISOString().slice(0, 10)}
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
