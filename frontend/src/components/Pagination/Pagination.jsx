import { useContext } from 'react'
import { TasksContext } from '../../context/TasksContext'
import Button from '../Button/Button'
import IconArrowLeft from '../IconArrowLeft/IconArrowLeft'
import IconArrowRight from '../IconArrowRight/IconArrowRight'
import styles from './Pagination.module.scss'

function Pagination() {
	const { page, totalPages, prevPage, nextPage } = useContext(TasksContext)

	return (
		<div className={styles.pagination}>
			<div className={styles.container}>
				<Button className={styles.button} onClick={prevPage}>
					<IconArrowLeft />
				</Button>
				<h3 className={styles.page}>
					{page}
					<span>/{totalPages}</span>
				</h3>
				<Button className={styles.button} onClick={nextPage}>
					<IconArrowRight />
				</Button>
			</div>
		</div>
	)
}

export default Pagination
