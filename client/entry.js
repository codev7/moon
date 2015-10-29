import React from 'react'
import ReactDOM from 'react-dom'

import App from './components/app.js'

import './style.scss'

ReactDOM.render(<App />, document.getElementById('app'))

if(module.hot) {
	module.hot.accept();
}
