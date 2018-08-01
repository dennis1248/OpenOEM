var log = console.log

function setup() {
  showTab()
  window.onhashchange = function () {showTab()}
}

function showTab(type) {
  var type = location.hash || '#readme'
  document.querySelector('#generator').hidden = type != '#generator'
  document.querySelector('#readme').hidden = type != '#readme'
  if (type == '#readme') {
    loadMarkDown()
  }
}

var markdownIsLoaded = false
function loadMarkDown() {
  if (!markdownIsLoaded) {
    markdownIsLoaded = true
    
    // fetch the readme
    unfetch('./markdown/README.md')
    .then(function(data) { return data.text() })
    .then(function(README) {
      // remove the logo and header from the readme
      README = README.replace(/^.{0,}[#|\s|a-z"0-9|]{0,}\n!\[.{0,}]\(.{0,}\s{0,}".{0,}"\)/gmi, '')
      var outHTML = snarkdown(README)
      document.querySelector('#readme').innerHTML = outHTML
    })
  }
}

setup()

setTimeout(function() {
  // auto load the markdown if the user doesn't click on the readme part
  loadMarkDown()
}, 6000);