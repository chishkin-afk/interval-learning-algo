import { useEffect, useState } from "react";

function matchPath(path, route) {
    const [splittedRoute, splittedPath] = [route.split('/'), path.split('/')] 
    if (splittedPath.length !== splittedPath.length) return null
    
    const params = {}

    for (let i = 0; i < splittedRoute.length; i++) {
        const part = splittedRoute[i]

        if (part.startsWith(':')) {
            const paramName = part.slice(1)
            params[paramName] = splittedPath[i]
        } else if (part !== splittedPath[i]) {
            return null
        }
    }

    return params
}

const usePath = () => {
    const [path, setPath] = useState(() => window.location.pathname)

    useEffect(() => {
        const onLocation = () => {
            setPath(window.location.pathname)
        }

        window.addEventListener('popstate', onLocation)

        return () => {
            window.removeEventListener('popstate', onLocation)
        }
    }, [])

    return path
}

const Router = ({routes}) => {
    const path = usePath()
    console.log(routes)

    for (const route of Object.keys(routes)) {
        const params = matchPath(path, route)

        if (params) {
            const Page = routes[path]
            return <Page params={params} />
        }
    }

    const NotFound = routes['*']
    return NotFound ? <NotFound /> : <div>Page not found</div>
}

export default Router