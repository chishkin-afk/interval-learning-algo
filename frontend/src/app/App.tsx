import NotFoundPage from '../pages/NotFoundPage'
import Router from './Router'
import './styles/index.scss'

function App() {
	const routes = {
		'*': NotFoundPage,
	}

	return <Router routes={routes} />
}

export default App
