import MainPage from '@/pages/main-page'
import NotFoundPage from '../pages/not-found-page'
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
