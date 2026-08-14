import RouterLink from "./components/RouterLink/RouterLink"
import NotFoundPage from "./pages/NotFoundPage"
import Router from "./Router"

const App = () => {
    const routes = {
        '*': NotFoundPage
    }

    return (
        <Router routes={routes}/>
    )
}

export default App