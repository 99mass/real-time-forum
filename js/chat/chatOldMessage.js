

const OldChatMessage=()=>{
    const socket = new WebSocket("ws://localhost:8080/communication");

    let Messdata={
        User1: "breukh",
        User2: "samba",
    }
    // Send user connected message when connection is open
    socket.onopen = () => {
        socket.send(JSON.stringify(Messdata));
        console.log("WebSocket communication on.");

    };

    socket.onmessage = (message) => {

        var _data = JSON.parse(message.data);
        console.log(_data);

    };
}

export{OldChatMessage}