import { useEffect, useState } from 'react'
import IconProfile from '../IconProfile/IconProfile'
import RouterLink from '../RouterLink/RouterLink'
import styles from './Navbar.module.scss'

function Navbar() {
	const [scrolled, setScrolled] = useState(false)
	useEffect(() => {
		function handleScroll() {
			setScrolled(window.scrollY > 10)
		}

		window.addEventListener('scroll', handleScroll)
		return () => window.removeEventListener('scroll', handleScroll)
	}, [])

	return (
		<nav className={!scrolled ? styles.navbar : styles.scrolled}>
			<h1>intask</h1>
			<RouterLink to="/profile" className={`${styles.link} ${styles.button}`}>
				<IconProfile />
			</RouterLink>
		</nav>
	)
}

export default Navbar
