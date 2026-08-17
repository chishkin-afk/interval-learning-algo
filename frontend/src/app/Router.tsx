import { useEffect, useState } from 'react'

// Params is a list of path params
interface Params {
	[key: string]: string
}

type RouteComponent = React.ComponentType<{ params: Params }>

interface Props {
	routes: Record<string, RouteComponent>
}

function matchPath(path: string, route: string): Params | null {
	const splittedPath = path.split('/')
	const splittedRoute = route.split('/')

	if (splittedPath.length !== splittedRoute.length) return null

	const params: Params = {}

	for (const [index, part] of splittedRoute.entries()) {
		if (part.startsWith(':')) {
			const paramName = part.slice(1)
			params[paramName] = splittedPath[index] ?? ''
		} else if (part !== splittedPath[index]) {
			return null
		}
	}

	return params
}

function usePath(): string {
	const [path, setPath] = useState(() => window.location.pathname)

	useEffect(() => {
		const handlePopstate = () => {
			setPath(window.location.pathname)
		}

		window.addEventListener('popstate', handlePopstate)

		return () => window.removeEventListener('popstate', handlePopstate)
	}, [])

	return path
}

function Router({ routes }: Props) {
	const path = usePath()

	for (const route of Object.keys(routes)) {
		const params = matchPath(path, route)
		if (!params) {
			continue
		}

		const Page = routes[route]
		return <Page params={params} />
	}

	const PageNotFound = routes['*']
	if (!PageNotFound) return null

	return <PageNotFound params={{}} />
}

export default Router
