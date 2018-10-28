const { app, BrowserWindow } = require('electron')
const childProcess = require('child_process')
const path = require('path')

let mainWindow
let backend

function init () {
  backend = childProcess.spawn(
    path.resolve(__dirname, 'backend.exe'), [],
    { cwd: path.resolve(__dirname) }
  )

  mainWindow = new BrowserWindow({ width: 4000, height: 3000 })
  // win.webContents.openDevTools()
  mainWindow.loadFile('index.html')
  mainWindow.on('closed', quit)
}

function quit () {
  backend.kill()
  app.quit()
}

app.on('ready', init)
app.on('window-all-closed', quit)
