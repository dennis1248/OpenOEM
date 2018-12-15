import React from 'react'

class Header extends React.Component {
  constructor() {
    super()
  }
  render() {
    return (
      <div className="part header">
        <div className="side side1">
          <h1>OpenOEM</h1>
          <p className="intro">
            OpenOEM will Install programs from the Chocolatey repositories and can make some basic
            tweaks such as changing the UI color, setting up a wallpaper and removing the people button among other things.
          </p>
          <div className="buttons">
            <a className="button" href="https://github.com/dennis1248/OpenOEM">About</a>
            <a className="button" href="#configure">Configurator</a>
          </div>
        </div>
        <div className="side side2">
          <div className="icon"></div>
        </div>
      </div>
    )
  }
}

export default Header
