const sse = new EventSource("http://127.0.0.1/events?subscribe[]=event1&subscribe[]=stream/chat/message&channel=stream")
const msgTemplate = document.querySelector("#template-message")
const messages = document.querySelector("#messages")
let config = {}


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

// Server Sent Events

sse.onmessage = function (event) {
    let msg = JSON.parse(event.data);
    const date = new Date();

    console.log(msg)

    switch (msg.type) {
        case "system/channel/config":
            config = msg
            const themeName = msg.payload.theme.name

            console.log(themeName)

            break;
        case "stream/chat/message":
            // @TODO Заменить картинку на DIV + URL + переменную в CSS
            const template = document.querySelector("#template-message")
            template.content.querySelector('.service-icon').setAttribute('src', config.payload.service_icons[msg.service])
            template.content.querySelector('.time').innerHTML = date.toTimeString().split(' ')[0]

            // Показывать/скрывать время
            if (config.payload.theme.config[config.payload.theme.name].tabs.main.fields.show_time.value) {
                document.documentElement.style.setProperty('--time_display', 'inline');
            } else {
                document.documentElement.style.setProperty('--time_display', 'none');
            }


            let message = template.innerHTML


            macros.forEach(m => {
                message = message.replaceAll(m.macro, value(msg.payload, m.path))
            })

            messages.insertAdjacentHTML('beforeend', message)
            window.scrollTo(0, document.body.scrollHeight);
            break;

        default:
            console.log("UNKNOWN MESSAGE TYPE")
            console.log(msg)
    }
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
