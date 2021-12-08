fetch('https://k07f479ka3.execute-api.us-east-1.amazonaws.com/Prod/put')
    .then(() => fetch('https://k07f479ka3.execute-api.us-east-1.amazonaws.com/Prod/get'))
    .then(response => response.json())
    .then((data) => {
        document.getElementById('replaceme').innerText = data.count;
});