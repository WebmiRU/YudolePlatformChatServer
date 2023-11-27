let main = document.querySelector('main')
let message = document.querySelector('#message').innerHTML

function connect() {
    var ws = new WebSocket('ws://127.0.0.1:5800/chat');
    ws.onopen = function () {
        ws.send(JSON.stringify({"type": "subscribe", "events": ["chat/message"]}));
        console.log('OPEN')
    };

    ws.onmessage = function (e) {
        let msg = JSON.parse(e.data);
        console.log(e.data)

        messageSend({
            'class': 'terran',
            'name': msg.user?.name ?? '[NONE]',
            'html': msg.html ?? '[NONE]',
        })
    };

    ws.onclose = function (e) {
        console.log('Socket is closed. Reconnect will be attempted in 1 second.', e.reason);
        setTimeout(function () {
            connect();
        }, 1000);
    };

    ws.onerror = function (err) {
        console.error('Socket encountered error: ', err.message, 'Closing socket');
        ws.close();
    };
}

connect();

function messageSend(v) {
    main.insertAdjacentHTML("afterbegin", messageFormat(v))
}

function testMsg() {
    let msg = JSON.parse('{"id":"","type":"chat/message","service":"goodgame","html":"432432432","text":"432432432","user":{"id":"","name":"E.Wolf","login":"E.Wolf","meta":{"badges":null}}}');

    messageSend({
        'class': 'terran',
        'name': msg.user.name,
        'html': msg.html,
    })
}

testMsg();
testMsg();

setInterval(() => {

}, 1000)

function messageFormat(v) {
    let msg = message

    for (let key in v) {
        if (v.hasOwnProperty(key)) {
            msg = msg.replace('{{' + key + '}}', v[key])
        }
    }

    return msg
}