const { app, BrowserWindow } = require('electron')
const childprocess = require('child_process')
const path = require('path')
const url = require('url')

let mainWindow
let backend

function init() {
  backend = childprocess.spawn(
    path.resolve(__dirname, 'backend.exe'), [],
    {cwd: path.resolve(__dirname)}
  )

  mainWindow = new BrowserWindow({width: 4000, height: 3000, windowsHide: true})
  mainWindow.maximize()
  mainWindow.loadURL(url.format({
    pathname: path.join(__dirname, 'index.html'),
    protocol: 'file:',
    slashes: true
  }))
  mainWindow.on('closed', quit)
}

function quit() {
  backend.kill()
  app.quit()
}

app.on('ready', init)
app.on('window-all-closed', quit)
