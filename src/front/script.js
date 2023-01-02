const prefix = "todo-"
const SERVER_ADDR = "http://localhost:1323"

async function statusServer() {
  await fetch(`${SERVER_ADDR}/ping`).then((res) => {
    text = res.status === 200 ? 'OK !' : 'ERROR'
    document.getElementById("status-server").innerText = text;
  });
}

async function statusDB() {
  await fetch(`${SERVER_ADDR}/pingdb`).then((res) => {
    text = res.status === 200 ? 'OK !' : 'ERROR'
    document.getElementById("status-db").innerText = text;
  });
}

async function deleteItem(id) {
  const el = document.getElementById(id)
  el.parentNode.removeChild(el)
  await deleteItemReq({id: id, value: "-"})
}

async function createTodo() {
  value = document.querySelector('#adder input').value
  if (!value) {
    return
  }
  item = { id: `${prefix}${Date.now()}`, value: value }
  addTodo(item)
  await postItemReq(item)
}

async function postItemReq(item) {
  console.log(item)
  await fetch(`${SERVER_ADDR}/create`, {body: JSON.stringify(item), headers: {'Content-Type': 'application/json'}, method: 'POST'}).then((res) => {
    console.log(res)
  });
}

async function deleteItemReq(item) {
  await fetch(`${SERVER_ADDR}/delete`, {body: JSON.stringify(item), headers: {'Content-Type': 'application/json'}, method: 'POST'}).then((res) => {
    console.log(res)
  });
}

function addTodo(item) {
  div = document.createElement('DIV')
  div.setAttribute("id", item.id)
  div.innerHTML = `<span>${item.value}</span><button onclick="deleteItem('${item.id}')">x</button>`
  document.getElementById("items").appendChild(div)
}

async function getList(){
  await fetch(`${SERVER_ADDR}/list`).then((res) => {
      res.json().then(items=> items.forEach(item => { addTodo(item) }))
  });
}

getList()
statusServer()
statusDB()
