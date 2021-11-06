
fetch('https://k07f479ka3.execute-api.us-east-1.amazonaws.com/Stage/put')
    .then(() => fetch('https://k07f479ka3.execute-api.us-east-1.amazonaws.com/Stage/get'))
    .then(response => response.json())
    .then((data) => {
        document.getElementById('replaceme').innerText = data.count;
});