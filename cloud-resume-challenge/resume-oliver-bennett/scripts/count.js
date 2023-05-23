fetch('https://q0qfjomyb0.execute-api.us-east-1.amazonaws.com/prod/put')
    .then(() => fetch('https://q0qfjomyb0.execute-api.us-east-1.amazonaws.com/prod/get'))
    .then(response => response.json())
    .then((data) => {
        document.getElementById('replaceme').innerText = data.count;
});