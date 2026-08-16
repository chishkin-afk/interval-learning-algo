import styles from './Field.module.scss'

function Field(props) {
	const {
		inputType = 'text',
		className = '',
		placeholder = '',
		onInput,
		id = '',
		autoComplete,
		value = '',
	} = props

	return (
		<input
			id={id}
			value={value}
			type={inputType}
			placeholder={placeholder}
			className={`${styles.field} ${className}`}
			onInput={onInput}
			autoComplete={autoComplete}
		/>
	)
}

export default Field
