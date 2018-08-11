import React from 'react'
import fetch from 'unfetch'

class Step extends React.Component {
  constructor(props) {
    super()
    this.state = {
      searchPkgs: [],
      selectedPkgs: [],
      search: ''
    }
    this.lastRequest = ''
    this.lastInput = ''
    this.searchBuzzy = false
    this.queries = {
      names: [''],
      data: [[]]
    }
    this.updateSearchRes = this.updateSearchRes.bind(this)
    this.fileInput = React.createRef()
  }
  updateSearchRes(event) {
    let search = event.target.value
    this.lastInput = search
    this.newSearch()
    this.setState({search})
  }
  newSearch() {
    // this function asks for packages that are available and limits the networks request to the chocolatery api to 1 at the time
    // it also caches the request so you have less network traffic
    let vm = this
    const dune = err => {
      if (err) console.error(err)
      vm.searchBuzzy = false
      vm.newSearch()
    }
    if (vm.lastInput != '' && vm.lastInput != vm.lastRequest && !this.searchBuzzy) {
      vm.lastRequest = vm.lastInput
      vm.searchBuzzy = true
      
      // check if the search is cached
      let check = vm.queries.names.indexOf(vm.lastRequest)
      
      if (check == -1) {
        fetch('/search/' + encodeURI(vm.lastRequest))
        .then(data => data.json())
        .then(output => {
  
          if (output.status) {
            let data = []
            for (let i = 0; i < output.data.length; i++) {
              data.push({
                name: output.data[i].replace(/https?:\/\/.{0,}\(Id='|',.{0,}'\)/gmi, ''),
                url: output.data[i]
              })
            }
  
            // cache the output
            vm.queries.data.push(data)
            vm.queries.names.push(vm.lastRequest)
            
            vm.updateOutput(vm.queries.data.length - 1)
          }
          dune()
        })
        .catch(dune)
      } else {
        vm.updateOutput(check)
        dune()
      }
    } else if (vm.lastInput == '') {
      vm.updateOutput(0)
    }
  }
  updateOutput(pointer) {
    this.setState({searchPkgs: this.queries.data[pointer]})
  }
  render() {
    return <div className="step">
      <h3>{this.props.item.name.replace(/([A-Z])/, ' $1').toLowerCase()}</h3>
      <p className="info">{this.props.item.dis}</p>
      <div className="contents">
        {this.props.item.type == 'bool' ? 
          <div className="bool">
            <div 
              className={(!this.props.item.data ? "selected " : "") + "first"}
              onClick={() => this.props.changeData(false)}
            >No</div>
            <div 
              className={(this.props.item.data ? "selected " : "") + "last"}
              onClick={() => this.props.changeData(true)}
            >Yes</div>
          </div>
        : this.props.item.type == 'options' ?
          <div className="options">
            {this.props.item.options.map((el, i, arr) => 
              <div 
                className={(i==0?'first ':'') + (arr.length-1==i?'last ':'') + (el == this.props.item.data ? 'selected' : '')}
                key={i}
                onClick={() => this.props.changeData(el)}
              >{el}</div>
            )}
          </div>
        : this.props.item.type == 'color' ?
          <div className="color">
            <input 
              type="color" 
              id="ColorPicker" 
              name="ColorPicker"
              value={'#' + this.props.item.data} 
              onChange={event => this.props.changeData(event.target.value.replace('#','').toUpperCase())}/>
            <label htmlFor="ColorPicker">Theme color: #{this.props.item.data.toLowerCase()}</label>
          </div>
        : this.props.item.type == 'chocolatey-search' ? 
          <div className="search">
            <div className="search-input">
              <input autoFocus placeholder="search" type="text" value={this.state.search} onChange={this.updateSearchRes} />
            </div>
            <div className="output">
              <div className="row row-1">
                <h3>Search ouput</h3>
                {this.state.searchPkgs.map((el,id) => 
                  <div 
                    key={id} 
                    onClick={() => {
                      let selectedPkgs = Object.assign([], this.state.selectedPkgs)
                      selectedPkgs.push(el)
                      this.setState({
                        searchPkgs: [],
                        search: '',
                        selectedPkgs
                      })
                      this.props.changeData(selectedPkgs.map(el => el.name))
                    }}
                    className="item"
                  >{ el.name }</div>
                )}
              </div>
              <div className="row row-2">
                <h3>To install</h3>
                {this.state.selectedPkgs.map((el,id) => 
                  <div 
                    key={id} 
                    onClick={() => {
                      let selectedPkgs = Object.assign([], this.state.selectedPkgs)
                      selectedPkgs.splice(id,1)
                      this.setState({
                        selectedPkgs
                      })
                    }}
                    className="item"
                  >{ el.name }</div>
                )}
              </div>
            </div>
          </div>
        : this.props.item.type == 'fileSelect' ? 
          <div className="fileSelect">
          <form>
            <button
              onClick={() => {
                this.props.changeData('')
                this.fileInput.current.value = null
              }}
            >Remove selected file</button>
            <input 
              id="upload" 
              type="file"
              accept="image/*"
              ref={this.fileInput}
              onInput={() => {
                this.props.changeData(this.fileInput.current.files[0])
              }}
              onClick={event => { 
                event.target.value = null
              }}
            />
          </form>
          </div>
        : console.log(this.props.item.type)}
      </div>
    </div>
  }
}

export default Step