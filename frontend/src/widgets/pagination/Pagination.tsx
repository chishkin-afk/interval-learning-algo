import { useTasksActions, useTasksState } from '@/entities/task/model'
import Button from '@/shared/ui/Button'
import { IconArrowLeft, IconArrowRight } from '@/shared/ui/Icon'
import styles from './Pagination.module.scss'

function Pagination() {
	const { currentPage, pages } = useTasksState()
	const { prevPage, nextPage } = useTasksActions()

	return (
		<div className={styles.pagination}>
			<div className={styles.container}>
				<Button className={styles.button} onClick={prevPage}>
					<IconArrowLeft />
				</Button>
				<h3 className={styles.page}>
					{currentPage}
					<span>/{pages}</span>
				</h3>
				<Button className={styles.button} onClick={nextPage}>
					<IconArrowRight />
				</Button>
			</div>
		</div>
	)
}

export default Pagination
