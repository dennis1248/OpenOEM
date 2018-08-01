const express = require('express')
const path = require('path')
const fetch = require('node-fetch')
const XmlDom = require('xmldom')
const DOMParser = XmlDom.DOMParser;
const app = express()
const log = console.log

app.use(express.static(path.join(__dirname, 'src')))
app.use('/markdown', express.static(path.join(__dirname, 'markdown')))

app.get('/search/:query', (req, res) => {
  fetch('https://chocolatey.org/api/v2/Search()?$filter=IsLatestVersion&$skip=0&$top=30&searchTerm=%27' + encodeURI(req.params.query) + '%27&targetFramework=%27%27&includePrerelease=false')
  .then(res => res.text())
  .then(data => {
    let parser = new DOMParser()
    let xmlDom = parser.parseFromString(data, "text/xml")
    let nodes = xmlDom.getElementsByTagName('entry')
    let toReturn = []
    for (let i = 0; i < nodes.length; i++) {
      const element = nodes[i].getElementsByTagName('id')[0];
      toReturn.push(element.firstChild.data)
    }
    res.json({status: true, data: toReturn})
  })
  .catch(err => {
    console.log(err)
    res.json({status: false})
  })
})

app.listen(3123, () => log('Example app listening on port 3123!'))