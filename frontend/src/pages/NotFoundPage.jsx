import RouterLink from "../components/RouterLink/RouterLink"

const NotFoundPage = () => {
    return (
        <div>
            Page not found
            <RouterLink 
                to="/"
            >
                <h1>На главную</h1>
            </RouterLink>
        </div>
    )
}

export default NotFoundPage