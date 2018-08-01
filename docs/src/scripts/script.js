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
  // this function limits the networks request to the chocolatery api to 1 at the time
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
          let data = output.data 

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

function reRender(pointer) {
  let data = queries.data[pointer]
  
  let results = document.createElement('div')
  results.classList = 'results'

  for (let i = 0; i < data.length; i++) {
    const el = data[i]
    let result = document.createElement('div')
    result.innerText = el
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

setup()
let searchInputBox = document.querySelector('.searchInputBox')