const statusTable = document.querySelector("#statusTable");
const dateNode = document.querySelector("#date");

async function getData(id) {
    fetch('http://' + window.location.hostname + '/status/' + id, {method: 'GET'}).then((response) => response.json()).then((data) => showEntity(data));
    await sleep(1000);
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

function showEntity(data) {
    let id = data['id'];
    let name = 'unknown';
    let status = data['status'];

    if (status) {
        return
    }

    let newRow = statusTable.insertRow(-1);

    let idCell = newRow.insertCell(0);
    let idTextNode = document.createTextNode(id);
    idCell.appendChild(idTextNode);

    let nameCell = newRow.insertCell(1);
    let nameTextNode = document.createTextNode('Unknown');
    nameCell.appendChild(nameTextNode);

    let statusCell = newRow.insertCell(2);
    if (status === true) {
        status = '✓';
    } else if (status === false) {
        status = '❌';
    } else {
        status = '❌';
    }
    let statusTextNode = document.createTextNode(status);
    let statusNode = document.createTextNode(status);
    statusCell.appendChild(statusTextNode);
}

const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
const days = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];


const d = new Date();
let month = months[d.getMonth()];
let day = days[d.getDay()];
let mday = d.getDate();

let newDateNode = document.createElement("h3");
let dateText = document.createTextNode(day + ', ' + mday + ' ' + month);
newDateNode.appendChild(dateText);
dateNode.appendChild(newDateNode);

for (let id = 2010860201; id < 2010860230; id++) {
    getData(id);
}