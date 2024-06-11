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

// let css = 'h1 { background: red; }'
// Добавляем тег <style type="text/css"/> для управления CSS-переменными
const style = document.createElement('style')
style.type = 'text/css';
document.head.appendChild(style)

// let sheet = window.document.styleSheets[0];
// sheet.insertRule('.service-icon { background-color: yellow; }', sheet.cssRules.length);

// sheet.addRule('strong', 'color: red;', -1);


function value(obj, path) {
    return path.split('.').reduce(function (prev, curr) {
        return prev ? prev[curr] : null
    }, obj || self)
}

// function setCssVar(variable, value) {
//     document.documentElement.style.setProperty(variable, value);
// }

// Server Sent Events

sse.onmessage = function (event) {
    let msg = JSON.parse(event.data);
    const date = new Date();

    console.log(msg)

    switch (msg.type) {
        case "system/channel/config":
            config = msg
            const theme = msg.payload.theme.name

            // Перезаписываем динамические стили темы для изменения иконки сообщений и прочих параметров
            style.innerHTML = '';
            Object.keys(msg.payload.service_icons).forEach(key => {
                const val = msg.payload.service_icons[key]
                console.log(".service-icon." + key + " { background-image: url(" + val + ") }")
                style.innerHTML += ".service-icon." + key + " { background-image: url(" + val + ") }"
            })

            // Показывать/скрывать время
            if (msg.payload.theme.config[theme].tabs.main.fields.show_time.value) {
                style.innerHTML += ".time { display: inline }"
            } else {
                style.innerHTML += ".time { display: none }"
            }

            break;

        case "stream/chat/message":
            // @TODO Заменить картинку на DIV + URL + переменную в CSS
            const template = document.querySelector("#template-message")
            // template.content.querySelector('.service-icon').setAttribute('src', config.payload.service_icons[msg.service])
            template.content.querySelector('.service-icon').classList.add(msg.service)
            template.content.querySelector('.time').innerHTML = date.toTimeString().split(' ')[0]
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
