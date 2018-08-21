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
      filesToInclude: [],
      items: [
        {
          name: 'programs',
          dis: 'Select programs to install using chocolatey',
          type: 'chocolatey-search',
          screenshot: false
        },{
          name: 'removeEdgeIcon',
          dis: 'Remove the Microsoft Edge icon from the desktop',
          type: 'bool',
          screenshot: 'edge.png'
        },{
          name: 'removeJunkApps',
          dis: 'Remove apps from start menu (this feature has issues, use with caution)',
          type: 'bool',
          screenshot: 'start-menu.png'
        },{
          name: 'removePeople',
          dis: 'Remove the people button from the taskbar',
          type: 'bool',
          screenshot: 'people.png'
        },{
          name: 'search',
          dis: 'The type of searchbar on the taskbar',
          type: 'options',
          options: ['full','icon','hidden'],
          screenshot: 'search.png'
        },{
          name: 'taskView',
          dis: 'Show the task view button on the taskbar',
          type: 'bool',
          screenshot: 'tasks-view.png'
        },{
          name: 'themeColor',
          dis: 'Set the Windows theme color',
          type: 'color',
          screenshot: 'color.png'
        },{
          name: 'wallpaper',
          dis: 'Select wallpaper, skip this if you don\'t want a wallpaper',
          type: 'fileSelect',
          screenshot: false
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
                  config={this.state.config}
                  filesToInclude={this.state.filesToInclude}
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
                  {data: this.state.config[this.state.items[this.state.currentItem].name]}
                )
              }
              changeData={toSet => {
                let newConfig = Object.assign(this.state.config)
                let filesToInclude = Object.assign(this.state.filesToInclude)
                if (this.state.items[this.state.currentItem].type == 'fileSelect') {
                  for (let i = 0; i < this.state.filesToInclude.length; i++) {
                    const el = this.state.filesToInclude[i];
                    if (el.name == newConfig[this.state.items[this.state.currentItem].name]) {
                      this.state.filesToInclude.splice(i,1)
                    }
                  }
                }
                let newName = ''
                if (toSet instanceof File) {
                  filesToInclude.push(toSet)
                  newName = toSet.name
                } else  {
                   newName = toSet
                }
                newConfig[this.state.items[this.state.currentItem].name] = newName
                this.setState({
                  config: newConfig,
                  filesToInclude: filesToInclude
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
          {this.state.items[this.state.currentItem] && this.state.items[this.state.currentItem].screenshot ? 
            <img className="previewImage" src={`/imgs/${this.state.items[this.state.currentItem].screenshot}`}/>
          :''}
        </div>
      </div>
    )
  }
}

export default Configure
