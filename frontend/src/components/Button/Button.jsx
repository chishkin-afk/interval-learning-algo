import styles from './Button.module.scss'

function Button(props) {
	const { className = '', onClick, children, type } = props

	return (
		<button onClick={onClick} className={`${styles.button} ${className}`} type={type}>
			{children}
		</button>
	)
}

export default Button
