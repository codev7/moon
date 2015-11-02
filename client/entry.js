import React from 'react'
import ReactDOM from 'react-dom'
import App from './components/app.js'
import About from './components/about.js'
import Contact from './components/contact.js'
import { Router, Route } from 'react-router'
import createBrowserHistory from 'history/lib/createBrowserHistory'

import './style.scss'

const routes = {
  path: '/',
  component: App,
  childRoutes: [
    { path: 'about', component: About },
    { path: 'contact', component: Contact },
  ]
};

ReactDOM.render(
	(<Router
		routes={routes}
		history={createBrowserHistory()} />),
		document.getElementById('app'))

if(module.hot) {
  module.hot.accept();
}
