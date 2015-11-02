import React from 'react'
import { Link } from 'react-router'

export default
class App extends React.Component {
	render() {
		return (
			<div>
				<Link to="/">Home</Link>
				<Link to='/about'>About</Link>
				<Link to='/contact'>Contact</Link>
				{this.props.children}
			</div>
		)
	}
}

