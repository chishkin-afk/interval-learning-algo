import type React from 'react'
import styles from './LinkButton.module.scss'

interface LinkButtonProps {
	href: string
	className?: string
	children: React.ReactNode
}

function LinkButton(props: LinkButtonProps) {
	const { href = '#', className = '', children } = props

	return (
		<a href={href} className={`${styles.link} ${className}`}>
			{children}
		</a>
	)
}

export default LinkButton
