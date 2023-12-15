import { sendMessages,recipientMessages } from "../layout/corps.js";



const sendMessage = (userName_) => {
    // Create WebSocket connection
    const socket = new WebSocket("ws://localhost:8080/message");

    // Send user connected message when connection is open
    socket.onopen = () => {
        socket.send(JSON.stringify({ Username: userName_ }));
        console.log("WebSocket message on.");
    };

    // Handle form submission
    const formChat = document.querySelector('.form-chat');
    formChat.addEventListener('submit', async function (event) {
        event.preventDefault();

        let _Sender = document.querySelector('input[name="Sender"]').value.trim();
        let _Recipient = document.querySelector('input[name="Recipient"]').value.trim();
        let _Message = document.querySelector('textarea[name="Message"]').value.trim();

        let messageData = {
            Sender: _Sender,
            Recipient: _Recipient,
            Message: _Message
        };

        // Send message when connection is open
        if (socket.readyState === WebSocket.OPEN) {
            // socket.send(JSON.stringify({ Username: _Sender}));
            socket.send(JSON.stringify(messageData));
            document.querySelector('textarea[name="Message"]').value = "";
            console.log("message send.");
        }
    });

    // Handle incoming messages
    socket.onmessage = (message) => {
        const chatBody = document.querySelector('.chat-body');
        var _data = JSON.parse(message.data);
        console.log(_data);
        if (_data["sender"] == userName_) {
            let send = sendMessages(_data["sender"], _data["message"], null);
            console.log("send: " + send);
            chatBody.appendChild(send);
        }
        if (_data["recipient"] == userName_) {
            let recip = recipientMessages(_data["recipient"], _data["message"], null);
            console.log("recip: " + recip);
            chatBody.appendChild(recip);

        }
    };

    // Handle errors
    socket.onerror = (error) => {
        console.error(`WebSocket error: ${error}`);
    };
}

export{sendMessage}