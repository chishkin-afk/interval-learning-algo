import { IconProfile } from '@/shared/ui/Icon'
import RouterLink from '@/shared/ui/RouterLink'
import { useEffect, useState } from 'react'
import styles from './Navbar.module.scss'

function Navbar() {
	const [isScrolled, setIsScrolled] = useState(false)
	useEffect(() => {
		function handleScroll() {
			setIsScrolled(window.scrollY > 20)
		}

		window.addEventListener('scroll', handleScroll)

		return () => window.removeEventListener('scroll', handleScroll)
	}, [])

	return (
		<nav className={!isScrolled ? styles.navbar : styles.scrolled}>
			<h1>intask</h1>
			<RouterLink to="/profile" className={`${styles.link} ${styles.button}`}>
				<IconProfile />
			</RouterLink>
		</nav>
	)
}

export default Navbar
