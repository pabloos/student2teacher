const uploadEndpoint = "/upload";

let input = document.querySelector('form input');

input.onchange = function() {
    document.querySelector('form p').innerHTML = input.files.length + ' files added';
}

document.querySelector('button').onclick = function() {
    let data = new FormData();
    data.append('file', input.files[0]);

    upload(data);
}

function upload(file) {

    console.log("subiendo");
    const request = {
        method: 'POST',
        body: file
    }

    fetch(uploadEndpoint, request)
} 