function formatData (data) {
  let arr = []
  for (let name in data) {
    arr.push(encodeURIComponent(name) + '=' + encodeURIComponent(data[name]))
  }
  arr.push(('v=' + Math.random()).replace('0.', ''))
  return arr.join('&')
}

function get (url, data) {
  return new Promise((resolve, reject) => {
    let xhr = new XMLHttpRequest()
    xhr.open('GET', url + '?' + formatData(data), true)
    xhr.send(null)
    xhr.onreadystatechange = () => {
      if (xhr.readyState !== 4) return
      if (xhr.status === 200) {
        resolve(xhr.responseText, xhr.responseXML)
      } else reject(xhr.status)
    }
  })
}

function post (url, data) {
  return new Promise((resolve, reject) => {
    let xhr = new XMLHttpRequest()
    xhr.open('POST', url, true)
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded')
    xhr.send(formatData(data))
    xhr.onreadystatechange = () => {
      if (xhr.readyState !== 4) return
      if (xhr.status === 200) {
        resolve(xhr.responseText, xhr.responseXML)
      } else reject(xhr.status)
    }
  })
}

function loadCSS (url) {
  return new Promise((resolve, reject) => {
    let head = document.getElementsByTagName('head')[0]
    let link = document.createElement('link')
    link.type = 'text/css'
    link.rel = 'stylesheet'
    link.href = url
    link.onload = resolve
    link.onerror = reject
    head.appendChild(link)
  })
}

function loadJS (url, onload) {
  return new Promise((resolve, reject) => {
    let body = document.getElementsByTagName('body')[0]
    let script = document.createElement('script')
    script.type = 'text/javascript'
    script.src = url
    script.onload = resolve
    script.onerror = reject
    body.appendChild(script)
  })
}

function appendCSS (css) {
  let head = document.getElementsByTagName('head')[0]
  let style = document.createElement('style')
  style.type = 'text/css'
  style.appendChild(document.createTextNode(css))
  head.appendChild(style)
}

// for ES6
export default {
  get, post, loadJS, loadCSS, appendCSS
}
