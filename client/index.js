window.onload = (e) => {
  document.getElementById('camera').src = 'http://' + window.location.hostname + ':8080';
}

function pan(direction) {
  let apiUrl = '/api/robots/CCTVBot/commands/pan'
  fetch(apiUrl, {
    mode: 'no-cors',
    method: 'POST',
    body: JSON.stringify({
      direction: direction
    })
  })
}

document.onkeydown = (e) => {
  let key = e.which ? e.which : e.keyCode
  let LEFT = 37, UP = 38, RIGHT = 39, DOWN = 40
  switch (key) {
    case LEFT:
      pan('left')
      break
    case UP:
      pan('up')
      break
    case DOWN:
      pan('down')
      break
    case RIGHT:
      pan('right')
      break
    default:
      break
  }
}
