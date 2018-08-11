import React from 'react'

import BackIcon from './svg-icons/arrow-back.js'
import config from './../../../examples/config.basic.json'
import Step from './step.js'
import Loadable from 'react-loadable'

const MKpackage = Loadable({
  loader: () => import('./makepackage.js'),
  loading: () => <div className="loading"></div>,
});

config.programs = []

class Configure extends React.Component {
  constructor() {
    super()
    this.state = {
      config,
      currentItem: 0,
      runningInstaller: false,
      items: [
        {
          name: 'programs',
          dis: 'Select programs to install using chocolatey',
          type: 'chocolatey-search'
        },{
          name: 'removeEdgeIcon',
          dis: 'Remove the Microsoft Edge icon from the start screen',
          type: 'bool'
        },{
          name: 'removeJunkApps',
          dis: 'Remove apps from start menu (this feature has issues, use with caution)',
          type: 'bool'
        },{
          name: 'removePeople',
          dis: 'Remove the people button from the taskbar',
          type: 'bool'
        },{
          name: 'search',
          dis: 'The type of searchbar on the taskbar',
          type: 'options',
          options: ['full','icon','hidden']
        },{
          name: 'taskView',
          dis: 'Show the task view button on the taskbar',
          type: 'bool'
        },{
          name: 'themeColor',
          dis: 'Set the Windows theme color',
          type: 'color',
        },{
          name: 'wallpaper',
          dis: 'Select wallpaper, skip this if you don\'t want a wallpaper',
          type: 'fileSelect'
        }
      ]
    }
  }
  render() {
    return (
      <div className="part configure">
        <a href="#home" className="head">
          <BackIcon />
          <h2>OpenOEM</h2>
        </a>
        <div className="steps">
          {this.state.currentItem == this.state.items.length
            ? <div className="step">
                <h3>Install</h3>
                {!this.state.runningInstaller ?
                  <p className="inf">Press the button below to create an installer + config</p>
                :''}
                <div className="InstallBtn">
                <MKpackage
                  config={config}
                  run={() => {
                    this.setState({runningInstaller: true})
                  }}
                />
                </div>
              </div>
            : <Step
              item={
                Object.assign(
                  {},
                  this.state.items[this.state.currentItem],
                  {data: config[this.state.items[this.state.currentItem].name]}
                )
              }
              changeData={toSet => {
                let newConfig = Object.assign(this.state.config)
                newConfig[this.state.items[this.state.currentItem].name] = toSet
                this.setState({
                  config: newConfig
                })
              }}
            />
          }
          {!this.state.runningInstaller
            ?<div className="statusBar">
              <button
                disabled={this.state.currentItem == 0}
                className="previous"
                onClick={() => this.setState({
                  currentItem: this.state.currentItem - 1
                })}
              >previous</button>
              <div className="state">
                Step <b>{this.state.currentItem + 1}</b> of <b>{this.state.items.length + 1}</b>
              </div>
              <button
                disabled={this.state.currentItem == this.state.items.length}
                className="next"
                onClick={() => this.setState({
                  currentItem: this.state.currentItem + 1
                })}
              >next</button>
            </div>
          :''}
        </div>
      </div>
    )
  }
}

export default Configure
