import RouterLink from "./components/RouterLink/RouterLink"
import MainPage from "./pages/MainPage/MainPage"
import NotFoundPage from "./pages/NotFoundPage"
import Router from "./Router"

const App = () => {
    const routes = {
        '/': MainPage,
        '*': NotFoundPage,
    }

    return (
        <Router routes={routes}/>
    )
}

export default App