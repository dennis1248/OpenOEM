import React from 'react'
import ReactDOM from 'react-dom'

import './style/home.styl'
import Header from './components/header.js'
import Configure from './components/configure.js'

class Base extends React.Component {
  constructor() {
    super()
    this.state = {
      view: location.hash == '#configure' ? 'configure' : 'home' 
    }
    window.onhashchange = () => this.hashChange()
  }
  hashChange() {
    this.setState({
      view: location.hash == '#configure' ? 'configure' : 'home' 
    })
  }
  render() {
    return (
      <div className="root">
        { (this.state.view == 'home')
          ? <Header />
          : <Configure />
        }
      </div>
    )
  }
}

ReactDOM.render(<Base />, document.getElementById('app'))