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
            A Windows configuration tool which requires no user interaction
            It installs programs, configures Windows
            and fixes/removes various Windows "features" that no-one likes
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
