import React from 'react'
import styles from './Button.module.scss'

interface ButtonProps {
	className?: string
	onClick?(event: React.MouseEvent<HTMLButtonElement> | undefined): void
	children: React.ReactNode
	type?: 'submit' | 'reset' | 'button' | undefined
}

function Button(props: ButtonProps) {
	const { className = '', onClick, children, type } = props

	return (
		<button onClick={onClick} className={`${styles.button} ${className}`} type={type}>
			{children}
		</button>
	)
}

export default Button
