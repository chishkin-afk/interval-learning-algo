import type { MouseEvent } from 'react'

interface Props {
	className?: string
	children: React.ReactNode
	to: string
}

function RouterLink(props: Props) {
	const { children, to, ...rest } = props

	function handleClick(event: MouseEvent<HTMLAnchorElement>) {
		event.preventDefault()
		window.history.pushState({}, '', to)
		window.dispatchEvent(new PopStateEvent('popstate'))
	}

	return (
		<a href={to} onClick={handleClick} {...rest}>
			{children}
		</a>
	)
}

export default RouterLink
