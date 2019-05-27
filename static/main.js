var socket = new WebSocket("ws://localhost:9000/echo");

socket.onopen = function () {
    console.log("Status: Connected\n");
};

socket.onmessage = function (e) {
    console.log(e);
};