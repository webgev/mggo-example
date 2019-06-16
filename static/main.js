var socket = new WebSocket("ws://localhost:9000/echo");

socket.onopen = function () {
    console.log("Status: Connected\n");
};
var socketID;
socket.onmessage = function (e) {
    var params = JSON.parse(e.data)
    if (params.event_name == "Socket.Connect") {
        socketID = params.msg.addr;
    } if (events[params.event_name]) {
        events[params.event_name](params.msg)
        console.log(e.data);
    }

};

var sendSocket = function(name, msg) {
    socket.send(JSON.stringify({
        event_name: name,
        msg: msg
    }))
}

var events = {};
var subscribe = function(event, callback) {
    events[event] = callback;
}