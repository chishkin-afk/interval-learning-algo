import Field from '@/shared/ui/Field'
import RouterLink from '@/shared/ui/RouterLink'

function NotFoundPage() {
	return (
		<div>
			<h1>Page not found</h1>
			<Field id="not-found" autoComplete="off" placeholder="sfsdf" />
			<RouterLink to="/">
				<h3>To main</h3>
			</RouterLink>
		</div>
	)
}

export default NotFoundPage
