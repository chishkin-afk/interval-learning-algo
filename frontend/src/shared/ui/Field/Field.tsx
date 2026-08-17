import type React from 'react'
import styles from './Field.module.scss'

interface FieldProps {
	inputType?: string
	className?: string
	placeholder?: string
	onInput?(event: React.InputEvent<HTMLInputElement> | undefined): void
	onChange?(event: React.ChangeEvent<HTMLInputElement> | undefined): void
	id: string
	autoComplete: React.HTMLInputAutoCompleteAttribute
	value?: string
}

function Field(props: FieldProps) {
	const {
		inputType = 'text',
		className = '',
		placeholder = '',
		onInput,
		onChange,
		id = '',
		autoComplete,
		value,
	} = props

	return (
		<input
			id={id}
			value={value}
			type={inputType}
			placeholder={placeholder}
			className={`${styles.field} ${className}`}
			onInput={onInput}
			onChange={onChange}
			autoComplete={autoComplete}
		/>
	)
}

export default Field
