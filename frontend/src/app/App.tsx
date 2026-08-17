import MainPage from '@/pages/MainPage'
import NotFoundPage from '../pages/NotFoundPage'
import Router from './Router'
import './styles/index.scss'

function App() {
	const routes = {
		'/': MainPage,
		'*': NotFoundPage,
	}

	return <Router routes={routes} />
}

export default App
