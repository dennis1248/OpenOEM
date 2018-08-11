import React from 'react'
import JSZip from 'jszip'
import { saveAs } from 'file-saver/FileSaver'
import fetch from 'unfetch'

import CB from './svg-icons/check-box.js'
import CBUC from './svg-icons/check-box-outline-blank.js'

class MKpackage extends React.Component {
  constructor() {
    super()
    this.state = {
      config: {},
      part: 0,
      running: false,
      DownloadLatestsReleaseInfo: false,
      DownloadLatestsReleaseZip: false,
      OpenZip: false,
      EditZipFile: false
    }
    this.running = false
  }
  makeConfig(configFromFile) {
    let config = this.state.config
    for (const key in config) {
      if (config.hasOwnProperty(key)) {
        if (!/(\/\/)|(INFO)/.test(key) && typeof configFromFile[key] != undefined) {
          configFromFile[key] = config[key]
        }
      }
    }
    return configFromFile
  }
  runningInstaller() {
    if (!this.running) {
      this.running = true
      this.props.run()
      this.setState({running: true}, () => {
        fetch('https://api.github.com/repos/dennis1248/OpenOEM/releases')
        .then(data => data.json())
        .then(output => {
          this.setState({DownloadLatestsReleaseInfo: true})
          return fetch(
            '/download/' + 
            encodeURIComponent(
              output[0]
              .assets[0]
              .browser_download_url
              .replace(/http?s:\/\/.{0,}\/.{0,}\/releases\/download\//ig,'')
              .replace('/', '|||')
            )
          )
        })
        .then(data => data.blob())
        .then(data => {
          this.setState({DownloadLatestsReleaseZip: true})
          return new JSZip().loadAsync(data)
        })
        .then(zip => {
          return Promise.all([zip.file("config.json").async("string"), zip.file("setup.exe").async("uint8array")])
        })
        .then(data => {
          this.setState({OpenZip: true})
          let config = JSON.stringify(this.makeConfig(JSON.parse(data[0])))
          let setup = data[1]
          let zip = new JSZip()
          zip.file('config.json', config)
          zip.file('setup.exe', setup)
          return zip.generateAsync({type:"blob"})
        })
        .then(blob => {
          this.setState({EditZipFile: true, })
          saveAs(blob, "OpenOEM.zip")
        })
        .catch(console.error)
      })
    }
  }
  render() {
    return (this.state.running 
      ? <div className="checklist">
        <div className="checkListHeader">Creating zip...</div>
        <div>{this.state.DownloadLatestsReleaseInfo ? <CB /> : <CBUC />} Download latests release info</div>
        <div>{this.state.DownloadLatestsReleaseZip ? <CB /> : <CBUC />} Download latests release zip</div>
        <div>{this.state.OpenZip ? <CB /> : <CBUC />} Open zip</div>
        <div>{this.state.EditZipFile ? <CB /> : <CBUC />} Edit zip file</div>
        <div>{this.state.EditZipFile ? <CB /> : <CBUC />} Download zip</div>
      </div>:<button
        onClick={() => {
          console.log("click")
          this.runningInstaller()
        }}
      >Create</button>
    )
  }
}

export default MKpackage