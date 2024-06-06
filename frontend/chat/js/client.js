let sse = new EventSource("http://127.0.0.1/events?subscribe[]=event1&subscribe[]=stream/chat/message");
let msgTemplate = document.getElementById("template-message");
let messages = document.getElementById("messages");

let macros = [];
[...msgTemplate.innerHTML.matchAll(/{{\s*(([a-z]+\.?)+)\s*}}/gm)].forEach(v => {
    macros.push({
        macro: v[0],
        path: v[1],
    });
})

function value(obj, path) {
    return path.split('.').reduce(function(prev, curr) {
        return prev ? prev[curr] : null
    }, obj || self)
}

// setTimeout(() => {
//     sse.close()
//
// }, 5000)

// Server Sent Events

sse.onmessage = function (event) {
    let msg = JSON.parse(event.data);

    let message = msgTemplate.innerHTML
    macros.forEach(m => {
        message = message.replaceAll(m.macro, value(msg.payload, m.path))
    })

    messages.insertAdjacentHTML('beforeend', message)
    window.scrollTo(0, document.body.scrollHeight);
}

sse.onopen = function (event) {
    console.log(event);
}

sse.onerror = function (event) {
    console.log(event);
}

sse.addEventListener('join', event => {
    console.log(event.data);
});

// sse.addEventListener('message', event => {
//     let msg = JSON.parse(event.data);
//
//     let message = msgTemplate.innerHTML
//     macros.forEach(m => {
//         message = message.replaceAll(m.macro, value(msg.payload, m.path))
//     })
//
//     messages.insertAdjacentHTML('beforeend', message)
//     console.log(message)
// });
//
// sse.addEventListener()
//
// sse.addEventListener('leave', event => {
//     console.log(event.data);
// });
