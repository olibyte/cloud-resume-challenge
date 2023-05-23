fetch('https://d-jouiuecq10.execute-api.us-east-1.amazonaws.com/prod/put')
    .then(() => fetch('https://d-jouiuecq10.execute-api.us-east-1.amazonaws.com/prod/get'))
    .then(response => response.json())
    .then((data) => {
        document.getElementById('replaceme').innerText = data.count;
});