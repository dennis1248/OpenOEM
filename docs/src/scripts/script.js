const log = console.log

function setup() {
  showTab()
  window.onhashchange = showTab
  setTimeout(function() {
    // auto load the markdown if the user doesn't click on the readme part
    loadMarkDown()
  }, 4000);
}

function showTab(type) {
  type = location.hash || '#readme'
  document.querySelector('#generator').hidden = type != '#generator'
  document.querySelector('#readme').hidden = type != '#readme'
  if (type == '#readme') {
    loadMarkDown()
  }
}

let markdownIsLoaded = false
function loadMarkDown() {
  if (!markdownIsLoaded) {
    markdownIsLoaded = true
    
    // fetch the readme
    unfetch('./markdown/README.md')
    .then(function(data) { return data.text() })
    .then(function(README) {
      
      // remove the logo and header from the readme
      README = README.replace(/^.{0,}[#|\s|a-z"0-9|]{0,}\n!\[.{0,}]\(.{0,}\s{0,}".{0,}"\)/gmi, '')

      // replace the github emojis with real emojis
      README = README.replace(/:heavy_check_mark:/g, '✅').replace(/:x:/g, '❌')

      // place the javascript
      document.querySelector('#readme').innerHTML = snarkdown(README)
    })
  }
}

let queries = {
  names: [''],
  data: [[]]
}

let lastRequest = ''
let lastInput = ''
let searchBuzzy = false
function newSearch () {
  // this function asks for packages that are available and limits the networks request to the chocolatery api to 1 at the time
  // it also caches the request so you have less network traffic
  function dune(err) {
    if (err) console.error(err)
    searchBuzzy = false
    newSearch()
  }
  if (lastInput != '' && lastInput != lastRequest && !searchBuzzy) {
    lastRequest = lastInput
    searchBuzzy = true
    
    // check if the search is cached
    let check = queries.names.indexOf(lastRequest)
    
    if (check == -1) {
      unfetch('/search/' + encodeURI(lastRequest))
      .then(function(data) { return data.json() })
      .then(function(output) {

        if (output.status) {
          let data = []
          for (let i = 0; i < output.data.length; i++) {
            data.push({
              name: output.data[i].replace(/https?:\/\/.{0,}\(Id='|',.{0,}'\)/gmi, ''),
              url: output.data[i]
            })
          }

          // cache the output
          queries.data.push(data)
          queries.names.push(lastRequest)
          
          reRender(queries.data.length - 1)
        }
        dune()
      })
      .catch(dune)
    } else {
      reRender(check)
      dune()
    }
  } else if (lastInput == '') {
    reRender(0)
  }
}

function renderPkgList() {
  let results = document.createElement('div')
  results.classList = 'choosen'
  
  let nextBtn = document.querySelector("#generator .options .nextBtn")
  if (PKGs.length > 0) {
    let info = document.createElement('p')
    info.classList = 'info'
    info.innerText = 'Programs that will be included in the config file:'
    results.appendChild(info)
    nextBtn.disabled = false
  } else {
    nextBtn.disabled = true
  }
  
  for (let i = 0; i < PKGs.length; i++) {
    const el = PKGs[i]
    let result = document.createElement('div')
    result.classList = 'result'
    let text = document.createElement('p')
    text.innerText = el
    result.appendChild(text)
    let add = document.createElement('button')
    add.innerText = 'Remove'
    add.onclick = function() {
      PKGs.splice(i, 1)
      renderPkgList()
    }
    result.appendChild(add)
    results.appendChild(result)
  }
  let oldRes = document.querySelector('.choosen')
  oldRes.parentNode.replaceChild(results, oldRes) 
}

let PKGs = []
function addPkgToList(name) {
  PKGs.push(name)
  renderPkgList()
  searchInputBox.value = ''
  lastInput = ''
  newSearch()
}

function reRender(pointer) {
  // re-render data to the screen as fast as possible
  let data = queries.data[pointer]
  
  let results = document.createElement('div')
  results.classList = 'results'

  if (data.length > 0) {
    let info = document.createElement('p')
    info.classList = 'info'
    info.innerText = 'Found packages:'
    results.appendChild(info)
  }

  for (let i = 0; i < data.length; i++) {
    const el = data[i]
    let result = document.createElement('div')
    result.classList = 'result'
    let text = document.createElement('p')
    text.innerText = el.name
    result.appendChild(text)
    let add = document.createElement('button')
    add.innerText = 'Add'
    add.onclick = function() {
      addPkgToList(el.name)
    }
    result.appendChild(add)
    results.appendChild(result)
  }
  
  let oldRes = document.querySelector('.results')
  oldRes.parentNode.replaceChild(results, oldRes) 
}

let searchPKG = function() {
  if (!searchInputBox) {
    searchInputBox = document.querySelector('.searchInputBox')
  }
  lastInput = searchInputBox.value
  newSearch()
}

function createZip() {
  // create custom zip file
  unfetch('https://api.github.com/repos/dennis1248/Automated-Windows-10-configuration/releases')
    .then(function(data) { return data.json() })
    .then(function(output) {
      return unfetch(
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
    .then(function(data) {return data.blob()})
    .then(function(data) {return new JSZip().loadAsync(data)})
    .then(function(zip) {
      return Promise.all([zip.file("config.json").async("string"), zip.file("setup.exe").async("uint8array")])
    })
    .then(function(data) {
      let config = JSON.parse(data[0])
      config.programs = PKGs
      config = JSON.stringify(config)
      let setup = data[1]
      let zip = new JSZip()
      zip.file('config.json', config)
      zip.file('setup.exe', setup)
      return zip.generateAsync({type:"blob"})
    })
    .then(function(blob) {
      log(blob)
      saveAs(blob, "winconfig.zip")
    })
}

let currentStep = 1
let totalSteps = 2
let nextStep = function() {
  currentStep ++
  for (let index = 0; index < totalSteps; index++) {
    let step = index + 1
    document.querySelector('.step' + step).hidden = currentStep != step
  }
  if (currentStep == totalSteps) createZip()
}

setup()
let searchInputBox = undefined